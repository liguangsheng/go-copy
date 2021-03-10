package copy

import (
	"github.com/modern-go/reflect2"
	"reflect"
	"unsafe"
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
		if isSignedInt(p.srcType.Kind()) && isSignedInt(p.destType.Kind()) && _mask&*(*uint8)(src) > 0 {
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
