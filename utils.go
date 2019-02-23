package copy

import (
	"reflect"
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

func indirectType(typ reflect.Type) reflect.Type {
	for typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}
	return typ
}
