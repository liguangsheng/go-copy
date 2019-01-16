package copy

import (
	"github.com/modern-go/reflect2"
	"time"
	"unsafe"
)

type CustomDescriptor interface {
	Descriptor
	SrcType() reflect2.Type
	DstType() reflect2.Type
}

func DefaultDescriptors() []CustomDescriptor {
	return []CustomDescriptor{
		TimeToInt64Descriptor{},
		Int64ToTimeDescriptor{},
	}
}

// -------------
type TimeToInt64Descriptor struct{}

func (d TimeToInt64Descriptor) SrcType() reflect2.Type {
	return reflect2.TypeOf(time.Time{})
}

func (d TimeToInt64Descriptor) DstType() reflect2.Type {
	return reflect2.TypeOf(int64(0))
}

func (d TimeToInt64Descriptor) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	unix := srcType.PackEFace(srcPtr).(*time.Time).Unix()
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(unix))
}

// -------------
type Int64ToTimeDescriptor struct{}

func (d Int64ToTimeDescriptor) SrcType() reflect2.Type {
	return reflect2.TypeOf(int64(0))
}

func (d Int64ToTimeDescriptor) DstType() reflect2.Type {
	return reflect2.TypeOf(time.Time{})
}

func (d Int64ToTimeDescriptor) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	unix := srcType.PackEFace(srcPtr).(*int64)
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(time.Unix(*unix, 0)))
}
