package main

import (
	"os"
	"reflect"
	"strings"
	"text/template"
)

func main() {
	check(generateCopyNumberFunc())
	check(generateCopyNumberTests())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var kinds = []reflect.Kind{
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Int,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uint,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
}

var funcMap = template.FuncMap{
	"title": strings.Title,
}

func generateCopyNumberFunc() error {
	tplString := `{{$destKinds := .kinds}}{{$srcKinds := .kinds}}package copy

import (
	"reflect"
	"unsafe"
)

func copyNumber(dest, src unsafe.Pointer, destKind, srcKind reflect.Kind) bool {
	switch destKind {
{{range $destkind := $destKinds}} {{$DestKind := $destkind.String|title}}		case reflect.{{$DestKind}}:
			switch srcKind {
{{range $srckind := $srcKinds}}{{$SrcKind := $srckind.String|title}}			case reflect.{{$SrcKind}}:
				*((*{{$destkind}})(dest)) = {{$destkind}}(*((*{{$srckind}})(src)))
				return true
{{end}}
			}
{{end}}
	}
	return false
}
`
	file, err := os.OpenFile("copy_number.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	tpl := template.Must(
		template.New("generateCopyNumberFunc").Funcs(funcMap).Parse(tplString),
	)

	err = tpl.Execute(file, map[string]interface{}{
		"kinds": kinds,
	})
	if err != nil {
		return err
	}
	return nil
}

func generateCopyNumberTests() error {
	tplString := `{{$destKinds := .kinds}}{{$srcKinds := .kinds}}package copy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyNumber(t *testing.T){
{{range $srckind := $srcKinds}}{{$SrcKind := $srckind.String|title}}
{{range $destkind := $destKinds}}{{$DestKind := $destkind.String|title}}
	t.Run("Test{{$SrcKind}}To{{$DestKind}}", func(t *testing.T) {
		var src {{$srckind}} = 18
		var dest {{$destkind}} = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, {{$destkind}}(src), dest)
	})
{{end}}
{{end}}
}
`
	file, err := os.OpenFile("copy_number_test.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	tpl := template.Must(
		template.New("generateCopyNumberTests").Funcs(funcMap).Parse(tplString),
	)

	err = tpl.Execute(file, map[string]interface{}{
		"kinds": kinds,
	})
	if err != nil {
		return err
	}
	return nil
}
