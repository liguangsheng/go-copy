package copy

import (
	"encoding/binary"
	"github.com/modern-go/reflect2"
)

type hash [16]byte

func hashType(dest, src interface{}) hash {
	return hashRType(
		reflect2.TypeOf(dest).RType(),
		reflect2.TypeOf(src).RType(),
	)
}

func hashRType(dest, src uintptr) hash {
	var h hash
	binary.LittleEndian.PutUint64(h[8:], uint64(dest))
	binary.LittleEndian.PutUint64(h[0:], uint64(src))
	return h
}
