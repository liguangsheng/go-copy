package copy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
			FieldInterface interface{}
			FieldInt       int
			FieldInt64     int64
			FieldFloat64   float64
			FieldString    string
			FieldStruct    FieldStruct
		}
		DestStruct struct {
			FieldInterface interface{}
			FieldInt       int
			FieldInt64     int64
			FieldFloat64   float64
			FieldString    string
			FieldStruct    FieldStruct
		}
	)

	var (
		src = SrcStruct{
			FieldInterface: "interface _field",
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
	a.Equal("interface _field", dest.FieldInterface)
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
		DuplicateField: "outer duplicate _pair",
		Embedding: Embedding{
			FieldString:          "embedding string _pair",
			DuplicateField:       "inner duplicate _pair",
			UniqueEmbeddingField: "unique embedding string _pair",
		},
		Embedding2: Embedding2{
			FieldString:    "embedding string _pair - 2",
			DuplicateField: "inner duplicate _pair - 2",
		},
	}

	var dest struct {
		Embedding
		DuplicateField string
		Embedding2
		UniqueEmbeddingField string
		FieldEmpty           string
	}
	t.Run("", func(t *testing.T) {

	})

	a := assert.New(t)
	a.NoError(Copy(&dest, src))
	a.Equal("embedding string _pair", dest.Embedding.FieldString)
	a.Equal("inner duplicate _pair", dest.Embedding.DuplicateField)
	a.Equal("outer duplicate _pair", dest.DuplicateField)
	a.Equal("embedding string _pair - 2", dest.Embedding2.FieldString)
	a.Equal("inner duplicate _pair - 2", dest.Embedding2.DuplicateField)
	a.Equal("unique embedding string _pair", dest.UniqueEmbeddingField)
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
