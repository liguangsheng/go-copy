package copy

import (
	"github.com/modern-go/reflect2"
	"reflect"
)

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}
	return typ
}

func assignable(dest, src reflect2.Type) bool {
	if src.AssignableTo(dest) {
		return true
	}

	if isNumberKind(src.Kind()) && isNumberKind(dest.Kind()) {
		return true
	}

	if src.Kind() == reflect.Interface && src.Kind() == reflect.Interface {
		return dest.Implements(src)
	}

	if dest.Kind() == src.Kind() && src.Kind() != reflect.Struct {
		return true
	}

	return false
}

func isNumberKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64:
		return true
	default:
		return false
	}
}

func isSignedIntsKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return true
	}
	return false
}
