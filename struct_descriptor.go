package copy

import (
	"github.com/modern-go/reflect2"
)

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
