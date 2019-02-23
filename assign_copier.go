package copy

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

type assignCopier struct{}

func (d *assignCopier) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	dstType.UnsafeSet(dstPtr, srcPtr)
}
