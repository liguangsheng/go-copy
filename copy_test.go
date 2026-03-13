package copy

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringToString(t *testing.T) {
	var src = "sudo i love you"
	var dest = "i love you"
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, "sudo i love you", dest)
}

func TestInt32ToEnum(t *testing.T) {
	type enum int32
	var src int32 = 2242
	var dest enum
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, enum(2242), dest)
}

func TestEnumToInt32(t *testing.T) {
	type enum int32
	var src enum = 2242
	var dest int32
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, int32(2242), dest)
}

func TestStringToEnum(t *testing.T) {
	type enum string
	var src string = "enum string"
	var dest enum
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, enum("enum string"), dest)
}

func TestEnumToString(t *testing.T) {
	type enum string
	var src enum = "enum string"
	var dest string
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, string("enum string"), dest)
}

func TestStructToStruct(t *testing.T) {
	type (
		FieldStruct struct {
			FieldInt int
		}
		SrcStruct struct {
			FieldInterface any
			FieldInt       int
			FieldInt64     int64
			FieldFloat64   float64
			FieldString    string
			FieldStruct    FieldStruct
		}
		DestStruct struct {
			FieldInterface any
			FieldInt       int
			FieldInt64     int64
			FieldFloat64   float64
			FieldString    string
			FieldStruct    FieldStruct
		}
	)

	var (
		src = SrcStruct{
			FieldInterface: "interface value",
			FieldInt:       1,
			FieldString:    "you are a good guy",
			FieldInt64:     3,
			FieldFloat64:   3.141592654,
			FieldStruct:    FieldStruct{FieldInt: 42},
		}
		dest = DestStruct{}
		cpr  = New()
		a    = assert.New(t)
	)
	a.NoError(cpr.Copy(&dest, src))
	a.Equal("interface value", dest.FieldInterface)
	a.Equal(int(1), dest.FieldInt)
	a.Equal("you are a good guy", dest.FieldString)
	a.Equal(int64(3), dest.FieldInt64)
	a.Equal(float64(3.141592654), dest.FieldFloat64)
	a.Equal(42, dest.FieldStruct.FieldInt)
}

func TestEmbeddingStruct(t *testing.T) {
	type Embedding struct {
		FieldString          string
		DuplicateField       string
		UniqueEmbeddingField string
	}
	type Embedding2 struct {
		FieldString    string
		DuplicateField string
	}
	var src = struct {
		DuplicateField string
		Embedding
		Embedding2
	}{
		DuplicateField: "outer duplicate",
		Embedding: Embedding{
			FieldString:          "embedding string",
			DuplicateField:       "inner duplicate",
			UniqueEmbeddingField: "unique embedding string",
		},
		Embedding2: Embedding2{
			FieldString:    "embedding string - 2",
			DuplicateField: "inner duplicate - 2",
		},
	}

	var dest struct {
		Embedding
		DuplicateField string
		Embedding2
		UniqueEmbeddingField string
		FieldEmpty           string
	}

	a := assert.New(t)
	a.NoError(Copy(&dest, src))
	a.Equal("embedding string", dest.Embedding.FieldString)
	a.Equal("inner duplicate", dest.Embedding.DuplicateField)
	a.Equal("outer duplicate", dest.DuplicateField)
	a.Equal("embedding string - 2", dest.Embedding2.FieldString)
	a.Equal("inner duplicate - 2", dest.Embedding2.DuplicateField)
	a.Equal("unique embedding string", dest.UniqueEmbeddingField)
	a.Equal("", dest.FieldEmpty)
}

type Interface interface {
	Foo()
}

type StringInterface string

func (StringInterface) Foo() { fmt.Println("foo") }

type Interface2 interface {
	Bar()
}

func TestInterfaceNotImplemented(t *testing.T) {
	type (
		SrcStruct struct {
			InterfaceToInterface2 Interface
		}
		DestStruct struct {
			InterfaceToInterface2 Interface2
		}
	)

	var (
		src = SrcStruct{
			InterfaceToInterface2: StringInterface("string interface"),
		}
		dest = DestStruct{}
		cpr  = New()
		a    = assert.New(t)
	)

	a.NotNil(cpr.Copy(&dest, src))
	a.Nil(dest.InterfaceToInterface2)
}

func TestTimeToInt64(t *testing.T) {
	var src = time.Now()
	var dest int64 = 0
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, src.Unix(), dest)
}

func TestInt64ToTime(t *testing.T) {
	var src = time.Now().Unix()
	var dest time.Time
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, dest.Unix(), src)
}

func TestEmbeddedInterfaceToDifferentStruct(t *testing.T) {
	type Iface interface{}
	type SrcStruct struct {
		Iface
		FieldStr string
		FieldInt int
	}
	type DestStruct struct {
		FieldStr string
		FieldInt int
	}

	src := SrcStruct{FieldStr: "test", FieldInt: 100}
	var dest DestStruct
	a := assert.New(t)
	a.NoError(New().Copy(&dest, &src))
	a.Equal("test", dest.FieldStr)
	a.Equal(100, dest.FieldInt)

	// Also test with non-pointer src
	var dest2 DestStruct
	a.NoError(New().Copy(&dest2, src))
	a.Equal("test", dest2.FieldStr)
	a.Equal(100, dest2.FieldInt)

	// Test with non-nil interface value
	src2 := SrcStruct{Iface: "iface value", FieldStr: "test2", FieldInt: 200}
	var dest3 DestStruct
	a.NoError(New().Copy(&dest3, &src2))
	a.Equal("test2", dest3.FieldStr)
	a.Equal(200, dest3.FieldInt)
}

func TestEmbeddedMethodInterfaceToDifferentStruct(t *testing.T) {
	type SrcStruct struct {
		Interface
		FieldStr string
		FieldInt int
	}
	type DestStruct struct {
		FieldStr string
		FieldInt int
	}

	src := SrcStruct{FieldStr: "test", FieldInt: 100}
	var dest DestStruct
	a := assert.New(t)
	a.NoError(New().Copy(&dest, &src))
	a.Equal("test", dest.FieldStr)
	a.Equal(100, dest.FieldInt)

	// With non-nil interface value
	src2 := SrcStruct{Interface: StringInterface("hello"), FieldStr: "test2", FieldInt: 200}
	var dest2 DestStruct
	a.NoError(New().Copy(&dest2, &src2))
	a.Equal("test2", dest2.FieldStr)
	a.Equal(200, dest2.FieldInt)
}

func TestEmbeddedInterfaceBothSides(t *testing.T) {
	type Iface interface{}
	type SrcStruct struct {
		Iface
		FieldStr string
		FieldInt int
	}
	type DestStruct struct {
		Iface
		FieldStr string
		FieldInt int
	}

	src := SrcStruct{Iface: "iface", FieldStr: "test", FieldInt: 100}
	var dest DestStruct
	a := assert.New(t)
	a.NoError(New().Copy(&dest, &src))
	a.Equal("iface", dest.Iface)
	a.Equal("test", dest.FieldStr)
	a.Equal(100, dest.FieldInt)
}

func TestCopyNumber(t *testing.T) {
	types := []struct {
		name string
		typ  reflect.Type
	}{
		{"int8", reflect.TypeFor[int8]()},
		{"int16", reflect.TypeFor[int16]()},
		{"int32", reflect.TypeFor[int32]()},
		{"int64", reflect.TypeFor[int64]()},
		{"int", reflect.TypeFor[int]()},
		{"uint8", reflect.TypeFor[uint8]()},
		{"uint16", reflect.TypeFor[uint16]()},
		{"uint32", reflect.TypeFor[uint32]()},
		{"uint64", reflect.TypeFor[uint64]()},
		{"uint", reflect.TypeFor[uint]()},
		{"uintptr", reflect.TypeFor[uintptr]()},
		{"float32", reflect.TypeFor[float32]()},
		{"float64", reflect.TypeFor[float64]()},
	}

	for _, st := range types {
		for _, dt := range types {
			t.Run(st.name+"To"+dt.name, func(t *testing.T) {
				src := reflect.New(st.typ)
				src.Elem().Set(reflect.ValueOf(18).Convert(st.typ))
				dest := reflect.New(dt.typ)

				assert.NoError(t, Copy(dest.Interface(), src.Elem().Interface()))
				expected := src.Elem().Convert(dt.typ).Interface()
				assert.Equal(t, expected, dest.Elem().Interface())
			})
		}
	}
}
