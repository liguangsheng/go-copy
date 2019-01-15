package copy

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

func deepFields(typ reflect.Type) []reflect.StructField {
	typ = indirectType(typ)
	num := typ.NumField()
	var fields []reflect.StructField
	for i := 0; i < num; i++ {
		field := typ.Field(i)
		if field.Anonymous {
			fields = append(fields, deepFields(field.Type)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func indirect(o interface{}) interface{} {
	for {
		typ := reflect2.TypeOf(o)
		if typ.Kind() == reflect.Ptr {
			o = typ.Indirect(o)
			return indirect(o)
		}
		return o
	}
}

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}
	return typ
}

func indirectValue(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func parseFiledName(field reflect.StructField) string {
	return field.Name
}
