package copy

import (
	"errors"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// ---------------------------------------------------------------------------
// Public interfaces
// ---------------------------------------------------------------------------

// Handler defines how to copy a value between two types using unsafe pointers.
type Handler interface {
	Copy(destType, srcType reflect.Type, dest, src unsafe.Pointer) error
}

// TypeSampler provides zero-value samples so the copier can infer the
// source and destination types via reflect.TypeOf.
type TypeSampler interface {
	TypeSample() (dest, src any)
}

// TypedHandler is a Handler bound to a specific (dest, src) type pair,
// determined by TypeSample.
type TypedHandler interface {
	Handler
	TypeSampler
}

// ---------------------------------------------------------------------------
// NameFunc — field name resolution strategy
// ---------------------------------------------------------------------------

// NameFunc extracts a matching name from a struct field.
// Returning "" means the field should be skipped.
type NameFunc func(field reflect.StructField) string

// NameByFieldName uses the Go field name as-is.
func NameByFieldName(field reflect.StructField) string {
	return field.Name
}

// NameByJSONTag uses the json struct tag as the field name.
func NameByJSONTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

// NameByCopyTag uses the copy struct tag as the field name.
func NameByCopyTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("copy"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

// ---------------------------------------------------------------------------
// Option — functional options for Copier
// ---------------------------------------------------------------------------

// Option configures a Copier during construction.
type Option func(c *Copier)

// WithNameFunc sets the field name resolution strategy.
func WithNameFunc(fn NameFunc) Option {
	return func(c *Copier) {
		c.nameFunc = fn
	}
}

// ---------------------------------------------------------------------------
// Copier
// ---------------------------------------------------------------------------

// Copier copies values between structs using unsafe memory operations.
type Copier struct {
	cache    sync.Map // typePairKey → Handler
	nameFunc NameFunc
}

// New creates a Copier with the given options.
// Built-in handlers (time.Time ↔ int64) are registered by default.
func New(opts ...Option) *Copier {
	c := &Copier{
		nameFunc: NameByFieldName,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.Register(timeToInt64Handler{}, int64ToTimeHandler{})
	return c
}

// Register adds typed handlers to the copier's cache.
func (c *Copier) Register(handlers ...TypedHandler) {
	for _, h := range handlers {
		dest, src := h.TypeSample()
		key := typePairKey{reflect.TypeOf(dest), reflect.TypeOf(src)}
		c.cache.Store(key, h)
	}
}

// Unregister removes handlers identified by their type samples.
func (c *Copier) Unregister(handlers ...TypeSampler) {
	for _, h := range handlers {
		dest, src := h.TypeSample()
		key := typePairKey{reflect.TypeOf(dest), reflect.TypeOf(src)}
		c.cache.Delete(key)
	}
}

// Copy copies field values from src into dest.
// Both dest and src should be pointers to structs (or assignable values).
func (c *Copier) Copy(dest, src any) error {
	destType := indirectType(reflect.TypeOf(dest))
	srcType := indirectType(reflect.TypeOf(src))
	return c.copyValue(dataPtr(dest), dataPtr(src), &fieldMapping{
		direct:   isAssignable(destType, srcType),
		destType: destType,
		srcType:  srcType,
		destSize: destType.Size(),
		srcSize:  srcType.Size(),
	})
}

// ---------------------------------------------------------------------------
// Global convenience API
// ---------------------------------------------------------------------------

// ErrUnsupportedCopy is returned when no handler or struct mapping exists
// for the given type pair.
var ErrUnsupportedCopy = errors.New("unsupported type pair for copy")

var defaultCopier = New()

// Copy copies field values from src into dest using the default Copier.
func Copy(dest, src any) error {
	return defaultCopier.Copy(dest, src)
}

// ---------------------------------------------------------------------------
// Internal — core copy dispatch
// ---------------------------------------------------------------------------

func (c *Copier) copyValue(dest, src unsafe.Pointer, fm *fieldMapping) error {
	if fm.direct {
		fm.doCopy(dest, src)
		return nil
	}

	key := typePairKey{fm.destType, fm.srcType}

	// fast path: cached handler
	if h, ok := c.cache.Load(key); ok {
		return h.(Handler).Copy(fm.destType, fm.srcType, dest, src)
	}

	// slow path: build struct field mappings and cache
	mappings := buildFieldMappings(fm.destType, fm.srcType, c.nameFunc)
	if mappings != nil {
		h := &structHandler{copier: c, fields: mappings}
		c.cache.Store(key, h)
		return h.Copy(fm.destType, fm.srcType, dest, src)
	}

	return ErrUnsupportedCopy
}

// ---------------------------------------------------------------------------
// Internal — fieldMapping (src/dest field pair)
// ---------------------------------------------------------------------------

// fieldMapping describes how to copy one (possibly merged) region
// from a source struct to a destination struct.
type fieldMapping struct {
	merged     bool // absorbed into another mapping via tryMerge
	direct     bool // can copy via raw memory (no handler needed)
	destSize   uintptr
	srcSize    uintptr
	destOffset uintptr
	srcOffset  uintptr
	destType   reflect.Type
	srcType    reflect.Type
}

// doCopy performs the actual memory copy for this field mapping.
func (fm *fieldMapping) doCopy(dest, src unsafe.Pointer) {
	destPtr := unsafe.Add(dest, fm.destOffset)
	srcPtr := unsafe.Add(src, fm.srcOffset)

	if fm.tryCopyNumber(destPtr, srcPtr) {
		return
	}

	n := min(fm.destSize, fm.srcSize)

	copy(
		unsafe.Slice((*byte)(destPtr), n),
		unsafe.Slice((*byte)(srcPtr), n),
	)

	// sign extension: fill remaining dest bytes with 0x00 or 0xFF
	if fm.destType != nil && fm.srcType != nil && fm.destSize > n {
		var fill byte
		if isSignedKind(fm.srcType.Kind()) && isSignedKind(fm.destType.Kind()) {
			signByte := *(*byte)(unsafe.Add(srcPtr, fm.srcSize-1))
			if signByte&0x80 != 0 {
				fill = 0xFF
			}
		}
		remaining := unsafe.Slice((*byte)(unsafe.Add(destPtr, n)), fm.destSize-n)
		for i := range remaining {
			remaining[i] = fill
		}
	}
}

// tryMerge attempts to merge a contiguous field mapping into this one.
// Returns true if the merge succeeded.
func (fm *fieldMapping) tryMerge(other *fieldMapping) bool {
	if fm.destOffset+fm.destSize != other.destOffset {
		return false
	}
	if fm.srcOffset+fm.srcSize != other.srcOffset {
		return false
	}

	fm.destSize += other.destSize
	fm.srcSize += other.srcSize
	fm.destType = nil
	fm.srcType = nil
	other.merged = true
	return true
}

// tryCopyNumber attempts a numeric type conversion copy.
// Returns false if the types are not numeric.
func (fm *fieldMapping) tryCopyNumber(dest, src unsafe.Pointer) bool {
	if fm.destType == nil || fm.srcType == nil {
		return false
	}
	return copyNumber(dest, src, fm.destType.Kind(), fm.srcType.Kind())
}

// ---------------------------------------------------------------------------
// Internal — struct handler
// ---------------------------------------------------------------------------

type structHandler struct {
	copier *Copier
	fields []*fieldMapping
}

func (h *structHandler) Copy(destType, srcType reflect.Type, dest, src unsafe.Pointer) error {
	for _, fm := range h.fields {
		if err := h.copier.copyValue(dest, src, fm); err != nil {
			return err
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// Internal — field analysis and mapping
// ---------------------------------------------------------------------------

// typePairKey is the cache key for handler lookups.
type typePairKey struct {
	dest, src reflect.Type
}

// fieldInfo holds metadata about a single struct field (possibly nested).
type fieldInfo struct {
	raw      reflect.StructField
	name     string  // resolved name (via NameFunc)
	path     string  // parent path (e.g. "Embedding")
	fullPath string  // complete path (e.g. "Embedding.Field")
	offset   uintptr // absolute offset from struct base
}

// flattenFields recursively extracts all leaf fields from a struct type,
// expanding anonymous (embedded) structs.
func flattenFields(typ reflect.Type, parent *fieldInfo, namer NameFunc) []fieldInfo {
	typ = indirectType(typ)
	var fields []fieldInfo
	for i := range typ.NumField() {
		sf := typ.Field(i)
		name := namer(sf)
		if name == "" {
			continue
		}

		fi := fieldInfo{
			raw:  sf,
			name: name,
		}
		if parent != nil {
			fi.path = parent.fullPath
			fi.offset = parent.offset + sf.Offset
		} else {
			fi.offset = sf.Offset
		}
		fi.fullPath = joinFieldPath(fi.path, name)

		if sf.Anonymous && sf.Type.Kind() != reflect.Interface {
			fields = append(fields, flattenFields(sf.Type, &fi, namer)...)
		} else {
			fields = append(fields, fi)
		}
	}
	return fields
}

// buildFieldMappings builds field mappings between two struct types.
// Returns nil if either type is not a struct.
func buildFieldMappings(destType, srcType reflect.Type, nameFunc NameFunc) []*fieldMapping {
	if destType.Kind() != reflect.Struct || srcType.Kind() != reflect.Struct {
		return nil
	}

	fieldsByName := make(map[string][]fieldInfo)
	fieldsByPath := make(map[string]fieldInfo)
	for _, fi := range flattenFields(srcType, nil, nameFunc) {
		fieldsByName[fi.name] = append(fieldsByName[fi.name], fi)
		fieldsByPath[fi.fullPath] = fi
	}

	var mappings []*fieldMapping
	for _, destField := range flattenFields(destType, nil, nameFunc) {
		srcFields, ok := fieldsByName[destField.name]
		if !ok {
			continue
		}

		var srcField *fieldInfo
		if len(srcFields) == 1 {
			srcField = &srcFields[0]
		} else if matched, ok := fieldsByPath[destField.fullPath]; ok {
			srcField = &matched
		}

		if srcField != nil {
			mappings = append(mappings, &fieldMapping{
				direct:     isAssignable(destField.raw.Type, srcField.raw.Type),
				destType:   destField.raw.Type,
				srcType:    srcField.raw.Type,
				destOffset: destField.offset,
				srcOffset:  srcField.offset,
				destSize:   destField.raw.Type.Size(),
				srcSize:    srcField.raw.Type.Size(),
			})
		}
	}

	return mergeContiguous(mappings)
}

// mergeContiguous merges adjacent field mappings that occupy contiguous
// memory in both source and destination, enabling a single memcpy.
func mergeContiguous(mappings []*fieldMapping) []*fieldMapping {
	for i := 0; i < len(mappings); i++ {
		fm := mappings[i]
		if !fm.direct || fm.merged {
			continue
		}
		for j, other := range mappings {
			if i == j {
				continue
			}
			if fm.tryMerge(other) {
				i--
				break
			}
		}
	}

	result := make([]*fieldMapping, 0, len(mappings))
	for _, fm := range mappings {
		if !fm.merged {
			result = append(result, fm)
		}
	}
	return result
}

// joinFieldPath joins parent and child path segments with ".".
func joinFieldPath(parent, child string) string {
	if parent == "" {
		return child
	}
	return parent + "." + child
}

// ---------------------------------------------------------------------------
// Internal — type helpers
// ---------------------------------------------------------------------------

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Pointer || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}
	return typ
}

func isAssignable(dest, src reflect.Type) bool {
	if src.AssignableTo(dest) {
		return true
	}
	if isNumberKind(src.Kind()) && isNumberKind(dest.Kind()) {
		return true
	}
	if src.Kind() == reflect.Interface && dest.Kind() == reflect.Interface {
		return dest.Implements(src)
	}
	if dest.Kind() == src.Kind() && src.Kind() != reflect.Struct {
		return true
	}
	return false
}

func isNumberKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isSignedKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// ---------------------------------------------------------------------------
// Internal — pointer extraction
// ---------------------------------------------------------------------------

// eface mirrors the runtime layout of an empty interface (any).
type eface struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}

// dataPtr extracts the data pointer from an interface value.
func dataPtr(v any) unsafe.Pointer {
	return (*eface)(unsafe.Pointer(&v)).data
}

// ---------------------------------------------------------------------------
// Built-in handlers — time.Time ↔ int64
// ---------------------------------------------------------------------------

type timeToInt64Handler struct{}

func (timeToInt64Handler) TypeSample() (dest, src any) {
	return int64(0), time.Time{}
}

func (timeToInt64Handler) Copy(destType, srcType reflect.Type, dest, src unsafe.Pointer) error {
	t := reflect.NewAt(srcType, src).Elem().Interface().(time.Time)
	*(*int64)(dest) = t.Unix()
	return nil
}

type int64ToTimeHandler struct{}

func (int64ToTimeHandler) TypeSample() (dest, src any) {
	return time.Time{}, int64(0)
}

func (int64ToTimeHandler) Copy(destType, srcType reflect.Type, dest, src unsafe.Pointer) error {
	*(*time.Time)(dest) = time.Unix(*(*int64)(src), 0)
	return nil
}

// ---------------------------------------------------------------------------
// Internal — numeric type conversion
// ---------------------------------------------------------------------------

var kindTypes = map[reflect.Kind]reflect.Type{
	reflect.Int:     reflect.TypeFor[int](),
	reflect.Int8:    reflect.TypeFor[int8](),
	reflect.Int16:   reflect.TypeFor[int16](),
	reflect.Int32:   reflect.TypeFor[int32](),
	reflect.Int64:   reflect.TypeFor[int64](),
	reflect.Uint:    reflect.TypeFor[uint](),
	reflect.Uint8:   reflect.TypeFor[uint8](),
	reflect.Uint16:  reflect.TypeFor[uint16](),
	reflect.Uint32:  reflect.TypeFor[uint32](),
	reflect.Uint64:  reflect.TypeFor[uint64](),
	reflect.Uintptr: reflect.TypeFor[uintptr](),
	reflect.Float32: reflect.TypeFor[float32](),
	reflect.Float64: reflect.TypeFor[float64](),
}

func copyNumber(dest, src unsafe.Pointer, destKind, srcKind reflect.Kind) bool {
	srcType, ok1 := kindTypes[srcKind]
	destType, ok2 := kindTypes[destKind]
	if !ok1 || !ok2 {
		return false
	}
	srcVal := reflect.NewAt(srcType, src).Elem()
	destVal := reflect.NewAt(destType, dest).Elem()
	destVal.Set(srcVal.Convert(destType))
	return true
}
