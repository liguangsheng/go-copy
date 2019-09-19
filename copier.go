package copy

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
	"github.com/golang/groupcache/lru"
)

type Copier = *copier

type copier struct {
	cacheSize   int
	typeCache   *lru.Cache
	fieldParser FieldParseFunc
}

func NewCopier(opts ...Option) *copier {
	c := &copier{
		typeCache:   lru.New(1000),
		fieldParser: ParseFiledByName,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Register(
		TimeToInt64Copier{},
		Int64ToTimeCopier{},
	)

	return c
}

// Register add typed copier to cache
func (c *copier) Register(copiers ...TypedCopier) {
	for _, co := range copiers {
		for _, pair := range co.Pairs() {
			c.typeCache.Add(pair, co)
		}
	}
}

func (c *copier) Copy(dst, src interface{}) error {
	var (
		dstType  = indirectType(reflect.TypeOf(dst))
		srcType  = indirectType(reflect.TypeOf(src))
		dstType2 = reflect2.Type2(dstType)
		srcType2 = reflect2.Type2(srcType)
		dstPtr   = reflect2.PtrOf(dst)
		srcPtr   = reflect2.PtrOf(src)
	)

	return c.copy(dstType2, srcType2, dstPtr, srcPtr)
}

func (c *copier) copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) error {
	cpr := c.parse(dstType, srcType)
	if cpr == nil {
		return errors.New("unsupported copy")
	}

	switch cpr.(type) {
	case *assignCopier:
		cpr.(*assignCopier).Copy(dstType, srcType, dstPtr, srcPtr)
	case *structDescriptor:
		for _, i := range cpr.(*structDescriptor).FieldDescriptors {
			c.copy(i.DstType, i.SrcType, unsafe.Pointer(i.DstOffset+uintptr(dstPtr)), unsafe.Pointer(i.SrcOffset+uintptr(srcPtr)))
		}
	default:
		cpr.(TypedCopier).Copy(dstType, srcType, dstPtr, srcPtr)
	}
	return nil
}

func (c *copier) parse(dstType, srcType reflect2.Type) interface{} {
	pair := TypePair{
		DstType: dstType.RType(),
		SrcType: srcType.RType(),
	}

	if cpr, ok := c.typeCache.Get(pair); ok {
		return cpr
	}

	if d := c.parseAssignable(dstType, srcType); d != nil {
		return c.save(pair, d)
	}

	if d := c.parseStructs(dstType, srcType); d != nil {
		return c.save(pair, d)
	}

	return nil
}

func (c *copier) parseAssignable(dstType, srcType reflect2.Type) *assignCopier {
	if dstType.AssignableTo(srcType) {
		return &assignCopier{}
	}
	return nil
}

func (c *copier) parseStructs(dstType, srcType reflect2.Type) *structDescriptor {
	if dstType.Kind() != reflect.Struct || srcType.Kind() != reflect.Struct {
		return nil
	}

	sd := &structDescriptor{
		DstType: dstType,
		SrcType: srcType,
	}

	dstFields := make(map[string]reflect.StructField)
	for _, field := range deepFields(dstType.Type1()) {
		name := c.fieldParser(field)
		if name != "" {
			dstFields[name] = field
		}
	}

	srcFields := make(map[string]reflect.StructField)
	for _, field := range deepFields(srcType.Type1()) {
		name := c.fieldParser(field)
		if name != "" {
			srcFields[name] = field
		}
	}

	for name, dstField := range dstFields {
		if srcField, ok := srcFields[name]; ok {
			c.parse(reflect2.Type2(dstField.Type), reflect2.Type2(srcField.Type))
			sd.FieldDescriptors = append(sd.FieldDescriptors, structFieldDescriptor{
				DstType:   reflect2.Type2(dstField.Type),
				SrcType:   reflect2.Type2(srcField.Type),
				DstOffset: dstField.Offset,
				SrcOffset: srcField.Offset,
			})
		}
	}

	return sd
}

func (c *copier) save(pair TypePair, d interface{}) interface{} {
	if d != nil {
		c.typeCache.Add(pair, d)
	}
	return d
}
