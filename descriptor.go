package copy

import (
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Descriptor interface {
	Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer)
}

// -------------
type assignableDescriptor struct {
	DstType reflect2.Type
	SrcType reflect2.Type
}

func (d *assignableDescriptor) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	dstType.UnsafeSet(dstPtr, srcPtr)
}

// -------------
type structFieldDescriptor struct {
	SrcOffset uintptr
	DstOffset uintptr
	DstType   reflect2.Type
	SrcType   reflect2.Type
}

type structDescriptor struct {
	DstType          reflect2.Type
	SrcType          reflect2.Type
	FieldDescriptors []structFieldDescriptor
}

func (d *structDescriptor) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	panic("never reach here")
}
