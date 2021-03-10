package copy

import "reflect"

type _field struct {
	raw      reflect.StructField
	path     string
	name     string
	fullPath string
	offset   uintptr
}

func deepFields(typ reflect.Type, parent *_field, namer NameFunc) []_field {
	typ = indirectType(typ)
	num := typ.NumField()
	var fields []_field
	for i := 0; i < num; i++ {
		field := _field{
			raw: typ.Field(i),
		}

		name := namer(field.raw)
		if name == "" {
			continue
		}
		field.name = name

		if parent != nil {
			field.path = parent.fullPath
			field.offset = parent.offset + field.raw.Offset
		} else {
			field.offset = field.raw.Offset
		}
		field.fullPath = pathJoin(field.path, name)

		if field.raw.Anonymous && field.raw.Type.Kind() != reflect.Interface {
			fields = append(fields, deepFields(field.raw.Type, &field, namer)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func pathJoin(parent string, child string) string {
	if parent == "" {
		return child
	}
	return parent + "." + child
}
