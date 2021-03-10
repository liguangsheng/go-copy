package copy

import (
	"reflect"
	"strings"
)

// Option dest New a Copier
type Option func(cpr *Copier)

// WithCacheSize New Copier with specify cache destSize
func WithCacheSize(size int) Option {
	return func(cpr *Copier) {
		cpr.cap = size
	}
}

// WithNameFunc New Copier with specify name function
func WithNameFunc(fn NameFunc) Option {
	return func(cpr *Copier) {
		cpr.nameFunc = fn
	}
}

// NameFunc function dest get _pair name
type NameFunc func(field reflect.StructField) string

// NameByFieldName get name by _pair name
func NameByFieldName(field reflect.StructField) string {
	return field.Name
}

// NameByJSONTag get name by json tag
func NameByJSONTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}

// NameByCopyTag get name by copy tag
func NameByCopyTag(field reflect.StructField) string {
	tag := strings.SplitN(field.Tag.Get("copy"), ",", 2)[0]
	if tag == "-" {
		return ""
	}
	return tag
}
