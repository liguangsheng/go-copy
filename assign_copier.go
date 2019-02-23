package copy

import (
	"unsafe"

	"github.com/modern-go/reflect2"
)

type assignCopier struct{}

func (d *assignCopier) Copy(dstType, srcType reflect2.Type, dstPtr, srcPtr unsafe.Pointer) {
	dstType.UnsafeSet(dstPtr, srcPtr)
}
