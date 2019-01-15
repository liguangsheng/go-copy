package copy

import (
	"time"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Descriptor interface {
	Copy(dstPtr, srcPtr unsafe.Pointer)
}

// -------------
type assignableDescriptor struct {
	DstType reflect2.Type
	SrcType reflect2.Type
}

func (d *assignableDescriptor) Copy(dstPtr, srcPtr unsafe.Pointer) {
	d.DstType.UnsafeSet(dstPtr, srcPtr)
}

// -------------
type StructFieldDescriptor struct {
	SrcOffset uintptr
	DstOffset uintptr
	DstType   reflect2.Type
	SrcType   reflect2.Type
}

type structDescriptor struct {
	DstType          reflect2.Type
	SrcType          reflect2.Type
	FieldDescriptors []StructFieldDescriptor
}

func (d *structDescriptor) Copy(dstPtr, srcPtr unsafe.Pointer) {}

// -------------
type TimeToInt64Descriptor struct{}

func (d *TimeToInt64Descriptor) Copy(dstPtr, srcPtr unsafe.Pointer) {
	unix := TypeTime.PackEFace(srcPtr).(*time.Time).Unix()
	TypeInt64.UnsafeSet(dstPtr, reflect2.PtrOf(unix))
}
