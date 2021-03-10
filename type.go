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

	if isIntKind(src.Kind()) && isIntKind(dest.Kind()) {
		return dest.Type1().Bits() >= src.Type1().Bits()
	}

	if src.Kind() == reflect.Interface && src.Kind() == reflect.Interface {
		return dest.Implements(src)
	}

	if dest.Kind() == src.Kind() && src.Kind() != reflect.Struct {
		return true
	}

	return false
}

func isIntKind(k reflect.Kind) bool {
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
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uintptr:
		return true
	default:
		return false
	}
}

func isSignedInt(k reflect.Kind) bool {
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
