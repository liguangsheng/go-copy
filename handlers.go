package copy

import (
	"github.com/modern-go/reflect2"
	"time"
	"unsafe"
)

// Handler copy handler interface
type Handler interface {
	Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error
}

// Samples interface
type Samples interface {
	Samples() (dest, src interface{})
}

// CustomHandler custom copy handler interfac3
type CustomHandler interface {
	Handler
	Samples
}

type timeToInt64 struct{}

func (d timeToInt64) Samples() (dest, src interface{}) {
	return int64(0), time.Time{}
}

func (d timeToInt64) Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error {
	unix := srcType.PackEFace(srcPtr).(*time.Time).Unix()
	destType.UnsafeSet(destPtr, reflect2.PtrOf(unix))
	return nil
}

type int64ToTime struct{}

func (d int64ToTime) Samples() (dest, src interface{}) {
	return time.Time{}, int64(0)
}

func (d int64ToTime) Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error {
	unix := srcType.PackEFace(srcPtr).(*int64)
	destType.UnsafeSet(destPtr, reflect2.PtrOf(time.Unix(*unix, 0)))
	return nil
}
