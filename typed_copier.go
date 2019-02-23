package copy

import (
	"time"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type TypePair struct {
	SrcType uintptr
	DstType uintptr
}

type TypedCopier interface {
	Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer)
	Pairs() []TypePair
}

// -------------
type TimeToInt64Copier struct{}

func (d TimeToInt64Copier) Pairs() []TypePair {
	return []TypePair{{
		SrcType: reflect2.TypeOf(time.Time{}).RType(),
		DstType: reflect2.TypeOf(int64(0)).RType(),
	}}
}

func (d TimeToInt64Copier) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	unix := srcType.PackEFace(srcPtr).(*time.Time).Unix()
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(unix))
}

// -------------
type Int64ToTimeCopier struct{}

func (d Int64ToTimeCopier) Pairs() []TypePair {
	return []TypePair{{
		SrcType: reflect2.TypeOf(int64(0)).RType(),
		DstType: reflect2.TypeOf(time.Time{}).RType(),
	}}
}

func (d Int64ToTimeCopier) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	unix := srcType.PackEFace(srcPtr).(*int64)
	dstType.UnsafeSet(dstPtr, reflect2.PtrOf(time.Unix(*unix, 0)))
}
