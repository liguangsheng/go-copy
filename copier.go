package copy

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Copier = *copier

type copier struct {
	cache map[string]Descriptor
}

func NewCopier(descriptors ...CustomDescriptor) *copier {
	c := &copier{
		cache: make(map[string]Descriptor),
	}
	for _, d := range DefaultDescriptors() {
		c.Register(d)
	}
	for _, d := range descriptors {
		c.Register(d)
	}
	return c
}

// Register add descriptor to cache
func (c *copier) Register(descriptor CustomDescriptor) {
	cacheKey := c.cacheKey(descriptor.DstType().RType(), descriptor.SrcType().RType())
	c.cache[cacheKey] = descriptor
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
	var (
		dstRType = dstType.RType()
		srcRType = srcType.RType()
		cacheKey = c.cacheKey(dstRType, srcRType)
	)
	descriptor, ok := c.cache[cacheKey]
	if !ok {
		descriptor = c.describe(dstType, srcType)
		if descriptor == nil {
			return errors.New("unsupported copy")
		}
	}

	if sd, ok := descriptor.(*structDescriptor); ok {
		for _, i := range sd.FieldDescriptors {
			c.copy(i.DstType, i.SrcType, unsafe.Pointer(i.DstOffset+uintptr(dstPtr)), unsafe.Pointer(i.SrcOffset+uintptr(srcPtr)))
		}
	} else {
		descriptor.Copy(dstType, srcType, unsafe.Pointer(dstPtr), unsafe.Pointer(srcPtr))
	}

	return nil
}

func (c *copier) describe(dstType, srcType reflect2.Type) Descriptor {
	var (
		dstRType = dstType.RType()
		srcRType = srcType.RType()
		cacheKey = c.cacheKey(dstRType, srcRType)
	)

	if des, ok := c.cache[cacheKey]; ok {
		return des
	}

	if d := c.describeAssignable(dstType, srcType); d != nil {
		return c.saveDescriptor(cacheKey, d)
	}

	if d := c.describeStruct(dstType, srcType); d != nil {
		return c.saveDescriptor(cacheKey, d)
	}

	return nil
}

func (c *copier) describeAssignable(dstType, srcType reflect2.Type) *assignableDescriptor {
	if dstType.AssignableTo(srcType) {
		return &assignableDescriptor{
			DstType: dstType,
			SrcType: srcType,
		}
	}
	return nil
}

func (c *copier) describeStruct(dstType, srcType reflect2.Type) *structDescriptor {
	if dstType.Kind() != reflect.Struct || srcType.Kind() != reflect.Struct {
		return nil
	}

	sd := &structDescriptor{
		DstType: dstType,
		SrcType: srcType,
	}

	dstFields := make(map[string]reflect.StructField)
	for _, field := range deepFields(dstType.Type1()) {
		dstFields[parseFiledName(field)] = field
	}

	srcFields := make(map[string]reflect.StructField)
	for _, field := range deepFields(srcType.Type1()) {
		srcFields[parseFiledName(field)] = field
	}

	for name, dstField := range dstFields {
		if srcField, ok := srcFields[name]; ok {
			c.describe(reflect2.Type2(dstField.Type), reflect2.Type2(srcField.Type))
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

func (c *copier) saveDescriptor(key string, d Descriptor) Descriptor {
	if d != nil {
		c.cache[key] = d
	}
	return d
}

func (c *copier) cacheKey(dstRType, srcRType uintptr) string {
	return fmt.Sprintf("%d-%d", dstRType, srcRType)
}
