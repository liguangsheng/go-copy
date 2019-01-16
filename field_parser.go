package copy

import (
	"reflect"
	"strings"
)

type FieldParser func(field reflect.StructField) string

func ParseFiledByName(field reflect.StructField) string {
	return field.Name
}

func ParseFieldByJSONTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

func ParseFieldByCopyTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("copy"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}
