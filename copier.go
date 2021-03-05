package copy

import (
	"encoding/binary"
	"errors"
	"reflect"
	"strings"
	"unsafe"

	"github.com/golang/groupcache/lru"
	"github.com/modern-go/reflect2"
)

var _globalCopier *Copier

func init() {
	_globalCopier = New()
}

// Copy values
func Copy(dest, src interface{}) error {
	return _globalCopier.Copy(dest, src)
}

// Option dest New a Copier
type Option func(cpr *Copier)

// WithCacheSize New Copier with specify cache destSize
func WithCacheSize(size int) Option {
	return func(cpr *Copier) {
		cpr.cap = size
	}
}

// WithNameFunc New Copier with specify name function
func WithNameFunc(fn NameFunc) Option {
	return func(cpr *Copier) {
		cpr.nameFunc = fn
	}
}

// NameFunc function dest get copyable name
type NameFunc func(field reflect.StructField) string

// NameByFieldName get name by copyable name
func NameByFieldName(field reflect.StructField) string {
	return field.Name
}

// NameByJSONTag get name by json tag
func NameByJSONTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

// NameByCopyTag get name by copy tag
func NameByCopyTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("copy"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

// New a *Copier
func New(opts ...Option) *Copier {
	c := &Copier{
		cache:    lru.New(1000),
		nameFunc: NameByFieldName,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Register(
		timeToInt64{},
		int64ToTime{},
	)

	return c
}

// Copier dest copy values
type Copier struct {
	cap      int
	cache    *lru.Cache
	nameFunc NameFunc
}

// Register add typed Copier dest cache
func (c *Copier) Register(handlers ...CustomHandler) {
	for _, h := range handlers {
		dest, src := h.Samples()
		hash := hashType(dest, src)
		c.cache.Add(hash, h)
	}
}

// Unregister remove typed Copier src cache
func (c *Copier) Unregister(handlers ...Samples) {
	for _, h := range handlers {
		dest, src := h.Samples()
		hash := hashType(dest, src)
		c.cache.Remove(hash)
	}
}

// Copy values
func (c *Copier) Copy(dest, src interface{}) error {
	var (
		srcType   = indirectType(reflect.TypeOf(src))
		srcType2  = reflect2.Type2(srcType)
		srcPtr    = reflect2.PtrOf(src)
		destType  = indirectType(reflect.TypeOf(dest))
		destType2 = reflect2.Type2(destType)
		destPtr   = reflect2.PtrOf(dest)
	)

	return c.copy(destPtr, srcPtr, copyable{
		assignable: assignableTo(srcType2, destType2),
		srcType:    srcType2,
		destType:   destType2,
		destSize:   srcType.Size(),
		srcSize:    destType.Size(),
	})
}

func (c *Copier) copy(destPtr, srcPtr unsafe.Pointer, ca copyable) error {
	// memory copy
	if ca.assignable {
		srcSize := ca.destSize
		destSize := ca.srcSize
		min := destSize
		max := destSize
		if destSize > srcSize {
			max = destSize
			min = srcSize
		} else {
			max = srcSize
			min = destSize
		}
		memcpy(
			unsafe.Pointer(uintptr(destPtr)+ca.destOffset),
			unsafe.Pointer(uintptr(srcPtr)+ca.srcOffset),
			min,
		)

		if max > min {
			memset(unsafe.Pointer(uintptr(destPtr)+ca.destOffset+min), 0, max-min)
		}

		return nil
	}

	// use cached copy handler
	hash := hashRType(ca.destType.RType(), ca.srcType.RType())
	if handler, ok := c.cache.Get(hash); ok {
		return handler.(Handler).Copy(ca.destType, ca.srcType, destPtr, srcPtr)
	}

	fields := parseStructs(ca.destType, ca.srcType, c.nameFunc)
	if fields != nil {
		sh := &structsHandler{copier: c, fields: fields}
		c.cache.Add(hash, sh)
		return sh.Copy(ca.destType, ca.srcType, destPtr, srcPtr)
	}

	return errors.New("unsupported copy")
}

func parseStructs(destType, srcType reflect2.Type, nameFunc NameFunc) []copyable {
	if destType.Kind() != reflect.Struct || srcType.Kind() != reflect.Struct {
		return nil
	}

	nameMap := make(map[string][]structField)
	fullNameMap := make(map[string]structField)
	for _, field := range deepFields(srcType.Type1(), nil, nameFunc) {
		nameMap[field.name] = append(nameMap[field.name], field)
		fullNameMap[field.fullPath] = field
	}

	var fields []*copyable
	for _, destField := range deepFields(destType.Type1(), nil, nameFunc) {
		srcFields, ok := nameMap[destField.name]
		if !ok {
			continue
		}
		if len(srcFields) == 1 {
			srcField := srcFields[0]
			fields = append(fields, &copyable{
				assignable: assignableTo(reflect2.Type2(srcField.raw.Type), reflect2.Type2(destField.raw.Type)),
				srcType:    reflect2.Type2(srcField.raw.Type),
				srcOffset:  srcField.offset,
				destType:   reflect2.Type2(destField.raw.Type),
				destOffset: destField.offset,
				destSize:   destField.raw.Type.Size(),
				srcSize:    srcField.raw.Type.Size(),
			})
			continue
		}

		srcField, ok := fullNameMap[destField.fullPath]
		if ok {
			fields = append(fields, &copyable{
				assignable: assignableTo(reflect2.Type2(srcField.raw.Type), reflect2.Type2(destField.raw.Type)),
				srcType:    reflect2.Type2(srcField.raw.Type),
				srcOffset:  srcField.offset,
				destType:   reflect2.Type2(destField.raw.Type),
				destOffset: destField.offset,
				destSize:   destField.raw.Type.Size(),
				srcSize:    srcField.raw.Type.Size(),
			})
		}
	}

	// merge
	for i := 0; i < len(fields); i++ {
		f1 := fields[i]
		if !f1.assignable || f1.ban {
			continue
		}

		dest1End := f1.destOffset + f1.destSize
		src1End := f1.srcOffset + f1.destSize
		for j, f2 := range fields {
			if i == j {
				continue
			}

			dest2Start := f2.destOffset
			src2Start := f2.srcOffset
			if src2Start == src1End &&
				dest2Start == dest1End {

				f1.destSize = f1.destSize + f2.destType.Type1().Size()
				f1.srcSize = f1.destSize
				f1.destType = nil
				f1.srcType = nil
				f2.ban = true
				i--
				break
			}
		}
	}

	var validFields []copyable
	for _, f := range fields {
		if !f.ban {
			validFields = append(validFields, *f)
		}
	}
	return validFields
}

func memcpy(dest, src unsafe.Pointer, n uintptr) unsafe.Pointer {
	cnt := n >> 3

	for i := uintptr(0); i < cnt; i++ {
		var destPtr = (*uint64)(unsafe.Pointer(uintptr(dest) + uintptr(8*i)))
		var srcPtr = (*uint64)(unsafe.Pointer(uintptr(src) + uintptr(8*i)))
		*destPtr = *srcPtr
	}
	left := n & 7
	for i := uintptr(0); i < left; i++ {
		var destPtr = (*uint8)(unsafe.Pointer(uintptr(dest) + uintptr(8*cnt+i)))
		var srcPtr = (*uint8)(unsafe.Pointer(uintptr(src) + uintptr(8*cnt+i)))

		*destPtr = *srcPtr
	}
	return dest
}

func memset(dest unsafe.Pointer, c int8, n uintptr) unsafe.Pointer {
	left := n & 7
	cnt := n >> 3
	if cnt > 0 {
		left += 8
	}
	var i uintptr = 0
	for i = 0; i < left; i++ {
		var destPtr = (*int8)(unsafe.Pointer(uintptr(dest) + uintptr(i)))
		*destPtr = c
	}
	if cnt < 2 {
		return dest
	}
	var firstPtr = (*int64)(dest)

	for i = 0; i < cnt-1; i++ {
		var destPtr = (*int64)(unsafe.Pointer(uintptr(dest) + uintptr(left+8*i)))
		*destPtr = *firstPtr
	}

	return dest
}

func assignableTo(dest, src reflect2.Type) bool {
	if dest.AssignableTo(src) {
		return true
	}

	if isIntKine(src.Kind()) && isIntKine(dest.Kind()) {
		return true
	}

	if dest.Kind() == src.Kind() && src.Kind() != reflect.Struct {
		return true
	}

	return false
}

func isIntKine(k reflect.Kind) bool {
	switch k {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uintptr:
		return true
	default:
		return false
	}
}

type copyable struct {
	ban        bool
	assignable bool
	destSize   uintptr
	srcSize    uintptr
	destOffset uintptr
	srcOffset  uintptr
	destType   reflect2.Type
	srcType    reflect2.Type
}

type structsHandler struct {
	copier *Copier
	fields []copyable
}

func (s *structsHandler) Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error {
	for _, i := range s.fields {
		if err := s.copier.copy(destPtr, srcPtr, i); err != nil {
			return err
		}
	}
	return nil
}

type structField struct {
	raw      reflect.StructField
	path     string
	name     string
	fullPath string
	offset   uintptr
}

func deepFields(typ reflect.Type, parent *structField, namer NameFunc) []structField {
	typ = indirectType(typ)
	num := typ.NumField()
	var fields []structField
	for i := 0; i < num; i++ {
		field := structField{
			raw: typ.Field(i),
		}

		name := namer(field.raw)
		if name == "" {
			continue
		}
		field.name = name

		if parent != nil {
			field.path = parent.fullPath
			field.offset = parent.offset + field.raw.Offset
		} else {
			field.offset = field.raw.Offset
		}
		field.fullPath = pathJoin(field.path, name)

		if field.raw.Anonymous && field.raw.Type.Kind() != reflect.Interface {
			fields = append(fields, deepFields(field.raw.Type, &field, namer)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func pathJoin(parent string, child string) string {
	if parent == "" {
		return child
	}
	return parent + "." + child
}

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}
	return typ
}

type hash [16]byte

func hashType(dest, src interface{}) hash {
	return hashRType(
		reflect2.TypeOf(dest).RType(),
		reflect2.TypeOf(src).RType(),
	)
}

func hashRType(dest, src uintptr) hash {
	var h hash
	binary.LittleEndian.PutUint64(h[8:], uint64(dest))
	binary.LittleEndian.PutUint64(h[0:], uint64(src))
	return h
}
