package copy

import (
	"github.com/modern-go/reflect2"
	"unsafe"
)

type structsHandler struct {
	copier *Copier
	fields []*_pair
}

func (s *structsHandler) Copy(destType, srcType reflect2.Type, destPtr, srcPtr unsafe.Pointer) error {
	for _, i := range s.fields {
		if err := s.copier.copy(destPtr, srcPtr, i); err != nil {
			return err
		}
	}
	return nil
}
