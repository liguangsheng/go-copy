package copy

import "unsafe"

func memcpy(dest, src unsafe.Pointer, n uintptr) unsafe.Pointer {
	cnt := n >> 3

	for i := uintptr(0); i < cnt; i++ {
		var destPtr = (*uint64)(unsafe.Pointer(uintptr(dest) + uintptr(8*i)))
		var srcPtr = (*uint64)(unsafe.Pointer(uintptr(src) + uintptr(8*i)))
		*destPtr = *srcPtr
	}
	left := n & 7
	for i := uintptr(0); i < left; i++ {
		var destPtr = (*uint8)(unsafe.Pointer(uintptr(dest) + uintptr(8*cnt+i)))
		var srcPtr = (*uint8)(unsafe.Pointer(uintptr(src) + uintptr(8*cnt+i)))

		*destPtr = *srcPtr
	}
	return dest
}

func memset(dest unsafe.Pointer, c int8, n uintptr) unsafe.Pointer {
	left := n & 7
	cnt := n >> 3
	if cnt > 0 {
		left += 8
	}
	for i := uintptr(0); i < left; i++ {
		var destPtr = (*int8)(unsafe.Pointer(uintptr(dest) + uintptr(i)))
		*destPtr = c
	}
	if cnt < 2 {
		return dest
	}
	var firstPtr = (*int64)(dest)

	for i := uintptr(0); i < cnt-1; i++ {
		var destPtr = (*int64)(unsafe.Pointer(uintptr(dest) + uintptr(left+8*i)))
		*destPtr = *firstPtr
	}

	return dest
}
