package copy

import (
	"unsafe"

	"github.com/modern-go/reflect2"
)

const _mask uint8 = 0b10000000

type _pair struct {
	ban        bool
	assignable bool
	destSize   uintptr
	srcSize    uintptr
	destOffset uintptr
	srcOffset  uintptr
	destType   reflect2.Type
	srcType    reflect2.Type
}

func (p *_pair) copy(dest, src unsafe.Pointer) {
	if p.copyNumber(unsafe.Pointer(uintptr(dest)+p.destOffset), unsafe.Pointer(uintptr(src)+p.srcOffset)) {
		return
	}

	max := p.destSize
	min := p.srcSize
	if max < min {
		max = p.srcSize
		min = p.destSize
	}
	memcpy(
		unsafe.Pointer(uintptr(dest)+p.destOffset),
		unsafe.Pointer(uintptr(src)+p.srcOffset),
		min,
	)

	if p.destType != nil && p.srcType != nil {
		var c int8 = 0
		if isSignedIntsKind(p.srcType.Kind()) && isSignedIntsKind(p.destType.Kind()) && _mask&*(*uint8)(src) > 0 {
			c = -0b1
		}

		if max > min {
			memset(unsafe.Pointer(uintptr(dest)+p.destOffset+min), c, max-min)
		}
	}
}

func (p *_pair) merge(p2 *_pair) bool {
	dest1End := p.destOffset + p.destSize
	src1End := p.srcOffset + p.destSize
	dest2Start := p2.destOffset
	src2Start := p2.srcOffset
	if src2Start != src1End || dest2Start != dest1End {
		return false
	}

	p.destSize = p.destSize + p2.destType.Type1().Size()
	p.srcSize = p.destSize
	p.destType = nil
	p.srcType = nil
	p2.ban = true
	return true
}

func (p *_pair) copyNumber(dest, src unsafe.Pointer) bool {
	if p.destType == nil || p.srcType == nil {
		return false
	}

	return copyNumber(dest, src, p.destType.Kind(), p.srcType.Kind())
}
