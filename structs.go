package copy

import (
	"github.com/modern-go/reflect2"
	"reflect"
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

func parseStructs(destType, srcType reflect2.Type, nameFunc NameFunc) []*_pair {
	if destType.Kind() != reflect.Struct || srcType.Kind() != reflect.Struct {
		return nil
	}

	nameMap := make(map[string][]_field)
	fullNameMap := make(map[string]_field)
	for _, field := range deepFields(srcType.Type1(), nil, nameFunc) {
		nameMap[field.name] = append(nameMap[field.name], field)
		fullNameMap[field.fullPath] = field
	}

	var fields []*_pair
	for _, destField := range deepFields(destType.Type1(), nil, nameFunc) {
		srcFields, ok := nameMap[destField.name]
		if !ok {
			continue
		}

		var srcField *_field
		if len(srcFields) == 1 {
			srcField = &srcFields[0]
		} else {
			matched, ok := fullNameMap[destField.fullPath]
			if ok {
				srcField = &matched
			}
		}

		if srcField != nil {
			fields = append(fields, &_pair{
				assignable: assignable(reflect2.Type2(destField.raw.Type), reflect2.Type2(srcField.raw.Type)),
				srcType:    reflect2.Type2(srcField.raw.Type),
				srcOffset:  srcField.offset,
				destType:   reflect2.Type2(destField.raw.Type),
				destOffset: destField.offset,
				destSize:   destField.raw.Type.Size(),
				srcSize:    srcField.raw.Type.Size(),
			})
		}
	}

	return merge(fields)
}

func merge(pairs []*_pair) []*_pair {
	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		if !p.assignable || p.ban {
			continue
		}

		for j, p2 := range pairs {
			if i == j {
				continue
			}

			if p.merge(p2) {
				i--
				break
			}
		}
	}

	var validFields []*_pair
	for _, f := range pairs {
		if !f.ban {
			validFields = append(validFields, f)
		}
	}
	return validFields
}
