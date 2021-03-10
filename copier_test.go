package copy

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToString(t *testing.T) {
	var src = "sudo i love you"
	var dest = "i love you"
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, "sudo i love you", dest)
}

func TestInt64ToInt64(t *testing.T) {
	var src int64 = 64
	var dest int64 = 32
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, int64(64), dest)
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
			FieldInterface: "interface field",
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
	a.Equal("interface field", dest.FieldInterface)
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
		DuplicateField: "outer duplicate copyable",
		Embedding: Embedding{
			FieldString:          "embedding string copyable",
			DuplicateField:       "inner duplicate copyable",
			UniqueEmbeddingField: "unique embedding string copyable",
		},
		Embedding2: Embedding2{
			FieldString:    "embedding string copyable - 2",
			DuplicateField: "inner duplicate copyable - 2",
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
	a.Equal("embedding string copyable", dest.Embedding.FieldString)
	a.Equal("inner duplicate copyable", dest.Embedding.DuplicateField)
	a.Equal("outer duplicate copyable", dest.DuplicateField)
	a.Equal("embedding string copyable - 2", dest.Embedding2.FieldString)
	a.Equal("inner duplicate copyable - 2", dest.Embedding2.DuplicateField)
	a.Equal("unique embedding string copyable", dest.UniqueEmbeddingField)
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

type StringInterface2 string

func (StringInterface2) Bar() { fmt.Println("bar") }

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

	cpr.Copy(&dest, src)
	a.Nil(dest.InterfaceToInterface2)
}

func TestCopyInts(t *testing.T) {
	//var val = 18
	var (
		src = struct {
			IntToInt         int
			IntToInt8        int
			IntToInt16       int
			IntToInt32       int
			IntToInt64       int
			IntToUInt        int
			IntToUInt8       int
			IntToUInt16      int
			IntToUInt32      int
			IntToUInt64      int
			IntToUintptr     int
			Int8ToInt        int8
			Int8ToInt8       int8
			Int8ToInt16      int8
			Int8ToInt32      int8
			Int8ToInt64      int8
			Int8ToUInt       int8
			Int8ToUInt8      int8
			Int8ToUInt16     int8
			Int8ToUInt32     int8
			Int8ToUInt64     int8
			Int8ToUintptr    int8
			Int16ToInt       int16
			Int16ToInt8      int16
			Int16ToInt16     int16
			Int16ToInt32     int16
			Int16ToInt64     int16
			Int16ToUInt      int16
			Int16ToUInt8     int16
			Int16ToUInt16    int16
			Int16ToUInt32    int16
			Int16ToUInt64    int16
			Int16ToUintptr   int16
			Int32ToInt       int32
			Int32ToInt8      int32
			Int32ToInt16     int32
			Int32ToInt32     int32
			Int32ToInt64     int32
			Int32ToUInt      int32
			Int32ToUInt8     int32
			Int32ToUInt16    int32
			Int32ToUInt32    int32
			Int32ToUInt64    int32
			Int32ToUintptr   int32
			Int64ToInt       int64
			Int64ToInt8      int64
			Int64ToInt16     int64
			Int64ToInt32     int64
			Int64ToInt64     int64
			Int64ToUInt      int64
			Int64ToUInt8     int64
			Int64ToUInt16    int64
			Int64ToUInt32    int64
			Int64ToUInt64    int64
			Int64ToUintptr   int64
			UIntToInt        uint
			UIntToInt8       uint
			UIntToInt16      uint
			UIntToInt32      uint
			UIntToInt64      uint
			UIntToUInt       uint
			UIntToUInt8      uint
			UIntToUInt16     uint
			UIntToUInt32     uint
			UIntToUInt64     uint
			UIntToUintptr    uint
			UInt8ToInt       uint8
			UInt8ToInt8      uint8
			UInt8ToInt16     uint8
			UInt8ToInt32     uint8
			UInt8ToInt64     uint8
			UInt8ToUInt      uint8
			UInt8ToUInt8     uint8
			UInt8ToUInt16    uint8
			UInt8ToUInt32    uint8
			UInt8ToUInt64    uint8
			UInt8ToUintptr   uint8
			UInt16ToInt      uint16
			UInt16ToInt8     uint16
			UInt16ToInt16    uint16
			UInt16ToInt32    uint16
			UInt16ToInt64    uint16
			UInt16ToUInt     uint16
			UInt16ToUInt8    uint16
			UInt16ToUInt16   uint16
			UInt16ToUInt32   uint16
			UInt16ToUInt64   uint16
			UInt16ToUintptr  uint16
			UInt32ToInt      uint32
			UInt32ToInt8     uint32
			UInt32ToInt16    uint32
			UInt32ToInt32    uint32
			UInt32ToInt64    uint32
			UInt32ToUInt     uint32
			UInt32ToUInt8    uint32
			UInt32ToUInt16   uint32
			UInt32ToUInt32   uint32
			UInt32ToUInt64   uint32
			UInt32ToUintptr  uint32
			UInt64ToInt      uint64
			UInt64ToInt8     uint64
			UInt64ToInt16    uint64
			UInt64ToInt32    uint64
			UInt64ToInt64    uint64
			UInt64ToUInt     uint64
			UInt64ToUInt8    uint64
			UInt64ToUInt16   uint64
			UInt64ToUInt32   uint64
			UInt64ToUInt64   uint64
			UInt64ToUintptr  uint64
			UintptrToInt     uintptr
			UintptrToInt8    uintptr
			UintptrToInt16   uintptr
			UintptrToInt32   uintptr
			UintptrToInt64   uintptr
			UintptrToUInt    uintptr
			UintptrToUInt8   uintptr
			UintptrToUInt16  uintptr
			UintptrToUInt32  uintptr
			UintptrToUInt64  uintptr
			UintptrToUintptr uintptr
		}{
			IntToInt:         18,
			IntToInt8:        18,
			IntToInt16:       18,
			IntToInt32:       18,
			IntToInt64:       18,
			IntToUInt:        18,
			IntToUInt8:       18,
			IntToUInt16:      18,
			IntToUInt32:      18,
			IntToUInt64:      18,
			IntToUintptr:     18,
			Int8ToInt:        18,
			Int8ToInt8:       18,
			Int8ToInt16:      18,
			Int8ToInt32:      18,
			Int8ToInt64:      18,
			Int8ToUInt:       18,
			Int8ToUInt8:      18,
			Int8ToUInt16:     18,
			Int8ToUInt32:     18,
			Int8ToUInt64:     18,
			Int8ToUintptr:    18,
			Int16ToInt:       18,
			Int16ToInt8:      18,
			Int16ToInt16:     18,
			Int16ToInt32:     18,
			Int16ToInt64:     18,
			Int16ToUInt:      18,
			Int16ToUInt8:     18,
			Int16ToUInt16:    18,
			Int16ToUInt32:    18,
			Int16ToUInt64:    18,
			Int16ToUintptr:   18,
			Int32ToInt:       18,
			Int32ToInt8:      18,
			Int32ToInt16:     18,
			Int32ToInt32:     18,
			Int32ToInt64:     18,
			Int32ToUInt:      18,
			Int32ToUInt8:     18,
			Int32ToUInt16:    18,
			Int32ToUInt32:    18,
			Int32ToUInt64:    18,
			Int32ToUintptr:   18,
			Int64ToInt:       18,
			Int64ToInt8:      18,
			Int64ToInt16:     18,
			Int64ToInt32:     18,
			Int64ToInt64:     18,
			Int64ToUInt:      18,
			Int64ToUInt8:     18,
			Int64ToUInt16:    18,
			Int64ToUInt32:    18,
			Int64ToUInt64:    18,
			Int64ToUintptr:   18,
			UIntToInt:        18,
			UIntToInt8:       18,
			UIntToInt16:      18,
			UIntToInt32:      18,
			UIntToInt64:      18,
			UIntToUInt:       18,
			UIntToUInt8:      18,
			UIntToUInt16:     18,
			UIntToUInt32:     18,
			UIntToUInt64:     18,
			UIntToUintptr:    18,
			UInt8ToInt:       18,
			UInt8ToInt8:      18,
			UInt8ToInt16:     18,
			UInt8ToInt32:     18,
			UInt8ToInt64:     18,
			UInt8ToUInt:      18,
			UInt8ToUInt8:     18,
			UInt8ToUInt16:    18,
			UInt8ToUInt32:    18,
			UInt8ToUInt64:    18,
			UInt8ToUintptr:   18,
			UInt16ToInt:      18,
			UInt16ToInt8:     18,
			UInt16ToInt16:    18,
			UInt16ToInt32:    18,
			UInt16ToInt64:    18,
			UInt16ToUInt:     18,
			UInt16ToUInt8:    18,
			UInt16ToUInt16:   18,
			UInt16ToUInt32:   18,
			UInt16ToUInt64:   18,
			UInt16ToUintptr:  18,
			UInt32ToInt:      18,
			UInt32ToInt8:     18,
			UInt32ToInt16:    18,
			UInt32ToInt32:    18,
			UInt32ToInt64:    18,
			UInt32ToUInt:     18,
			UInt32ToUInt8:    18,
			UInt32ToUInt16:   18,
			UInt32ToUInt32:   18,
			UInt32ToUInt64:   18,
			UInt32ToUintptr:  18,
			UInt64ToInt:      18,
			UInt64ToInt8:     18,
			UInt64ToInt16:    18,
			UInt64ToInt32:    18,
			UInt64ToInt64:    18,
			UInt64ToUInt:     18,
			UInt64ToUInt8:    18,
			UInt64ToUInt16:   18,
			UInt64ToUInt32:   18,
			UInt64ToUInt64:   18,
			UInt64ToUintptr:  18,
			UintptrToInt:     18,
			UintptrToInt8:    18,
			UintptrToInt16:   18,
			UintptrToInt32:   18,
			UintptrToInt64:   18,
			UintptrToUInt:    18,
			UintptrToUInt8:   18,
			UintptrToUInt16:  18,
			UintptrToUInt32:  18,
			UintptrToUInt64:  18,
			UintptrToUintptr: 18,
		}

		dest = struct {
			IntToIntBefore         int
			IntToInt               int
			IntToIntAfter          int
			IntToInt8Before        int8
			IntToInt8              int8
			IntToInt8After         int8
			IntToInt16Before       int16
			IntToInt16             int16
			IntToInt16After        int16
			IntToInt32Before       int32
			IntToInt32             int32
			IntToInt32After        int32
			IntToInt64Before       int64
			IntToInt64             int64
			IntToInt64After        int64
			IntToUIntBefore        uint
			IntToUInt              uint
			IntToUIntAfter         uint
			IntToUInt8Before       uint8
			IntToUInt8             uint8
			IntToUInt8After        uint8
			IntToUInt16Before      uint16
			IntToUInt16            uint16
			IntToUInt16After       uint16
			IntToUInt32Before      uint32
			IntToUInt32            uint32
			IntToUInt32After       uint32
			IntToUInt64Before      uint64
			IntToUInt64            uint64
			IntToUInt64After       uint64
			IntToUintptrBefore     uintptr
			IntToUintptr           uintptr
			IntToUintptrAfter      uintptr
			Int8ToIntBefore        int
			Int8ToInt              int
			Int8ToIntAfter         int
			Int8ToInt8Before       int8
			Int8ToInt8             int8
			Int8ToInt8After        int8
			Int8ToInt16Before      int16
			Int8ToInt16            int16
			Int8ToInt16After       int16
			Int8ToInt32Before      int32
			Int8ToInt32            int32
			Int8ToInt32After       int32
			Int8ToInt64Before      int64
			Int8ToInt64            int64
			Int8ToInt64After       int64
			Int8ToUIntBefore       uint
			Int8ToUInt             uint
			Int8ToUIntAfter        uint
			Int8ToUInt8Before      uint8
			Int8ToUInt8            uint8
			Int8ToUInt8After       uint8
			Int8ToUInt16Before     uint16
			Int8ToUInt16           uint16
			Int8ToUInt16After      uint16
			Int8ToUInt32Before     uint32
			Int8ToUInt32           uint32
			Int8ToUInt32After      uint32
			Int8ToUInt64Before     uint64
			Int8ToUInt64           uint64
			Int8ToUInt64After      uint64
			Int8ToUintptrBefore    uintptr
			Int8ToUintptr          uintptr
			Int8ToUintptrAfter     uintptr
			Int16ToIntBefore       int
			Int16ToInt             int
			Int16ToIntAfter        int
			Int16ToInt8Before      int8
			Int16ToInt8            int8
			Int16ToInt8After       int8
			Int16ToInt16Before     int16
			Int16ToInt16           int16
			Int16ToInt16After      int16
			Int16ToInt32Before     int32
			Int16ToInt32           int32
			Int16ToInt32After      int32
			Int16ToInt64Before     int64
			Int16ToInt64           int64
			Int16ToInt64After      int64
			Int16ToUIntBefore      uint
			Int16ToUInt            uint
			Int16ToUIntAfter       uint
			Int16ToUInt8Before     uint8
			Int16ToUInt8           uint8
			Int16ToUInt8After      uint8
			Int16ToUInt16Before    uint16
			Int16ToUInt16          uint16
			Int16ToUInt16After     uint16
			Int16ToUInt32Before    uint32
			Int16ToUInt32          uint32
			Int16ToUInt32After     uint32
			Int16ToUInt64Before    uint64
			Int16ToUInt64          uint64
			Int16ToUInt64After     uint64
			Int16ToUintptrBefore   uintptr
			Int16ToUintptr         uintptr
			Int16ToUintptrAfter    uintptr
			Int32ToIntBefore       int
			Int32ToInt             int
			Int32ToIntAfter        int
			Int32ToInt8Before      int8
			Int32ToInt8            int8
			Int32ToInt8After       int8
			Int32ToInt16Before     int16
			Int32ToInt16           int16
			Int32ToInt16After      int16
			Int32ToInt32Before     int32
			Int32ToInt32           int32
			Int32ToInt32After      int32
			Int32ToInt64Before     int64
			Int32ToInt64           int64
			Int32ToInt64After      int64
			Int32ToUIntBefore      uint
			Int32ToUInt            uint
			Int32ToUIntAfter       uint
			Int32ToUInt8Before     uint8
			Int32ToUInt8           uint8
			Int32ToUInt8After      uint8
			Int32ToUInt16Before    uint16
			Int32ToUInt16          uint16
			Int32ToUInt16After     uint16
			Int32ToUInt32Before    uint32
			Int32ToUInt32          uint32
			Int32ToUInt32After     uint32
			Int32ToUInt64Before    uint64
			Int32ToUInt64          uint64
			Int32ToUInt64After     uint64
			Int32ToUintptrBefore   uintptr
			Int32ToUintptr         uintptr
			Int32ToUintptrAfter    uintptr
			Int64ToIntBefore       int
			Int64ToInt             int
			Int64ToIntAfter        int
			Int64ToInt8Before      int8
			Int64ToInt8            int8
			Int64ToInt8After       int8
			Int64ToInt16Before     int16
			Int64ToInt16           int16
			Int64ToInt16After      int16
			Int64ToInt32Before     int32
			Int64ToInt32           int32
			Int64ToInt32After      int32
			Int64ToInt64Before     int64
			Int64ToInt64           int64
			Int64ToInt64After      int64
			Int64ToUIntBefore      uint
			Int64ToUInt            uint
			Int64ToUIntAfter       uint
			Int64ToUInt8Before     uint8
			Int64ToUInt8           uint8
			Int64ToUInt8After      uint8
			Int64ToUInt16Before    uint16
			Int64ToUInt16          uint16
			Int64ToUInt16After     uint16
			Int64ToUInt32Before    uint32
			Int64ToUInt32          uint32
			Int64ToUInt32After     uint32
			Int64ToUInt64Before    uint64
			Int64ToUInt64          uint64
			Int64ToUInt64After     uint64
			Int64ToUintptrBefore   uintptr
			Int64ToUintptr         uintptr
			Int64ToUintptrAfter    uintptr
			UIntToIntBefore        int
			UIntToInt              int
			UIntToIntAfter         int
			UIntToInt8Before       int8
			UIntToInt8             int8
			UIntToInt8After        int8
			UIntToInt16Before      int16
			UIntToInt16            int16
			UIntToInt16After       int16
			UIntToInt32Before      int32
			UIntToInt32            int32
			UIntToInt32After       int32
			UIntToInt64Before      int64
			UIntToInt64            int64
			UIntToInt64After       int64
			UIntToUIntBefore       uint
			UIntToUInt             uint
			UIntToUIntAfter        uint
			UIntToUInt8Before      uint8
			UIntToUInt8            uint8
			UIntToUInt8After       uint8
			UIntToUInt16Before     uint16
			UIntToUInt16           uint16
			UIntToUInt16After      uint16
			UIntToUInt32Before     uint32
			UIntToUInt32           uint32
			UIntToUInt32After      uint32
			UIntToUInt64Before     uint64
			UIntToUInt64           uint64
			UIntToUInt64After      uint64
			UIntToUintptrBefore    uintptr
			UIntToUintptr          uintptr
			UIntToUintptrAfter     uintptr
			UInt8ToIntBefore       int
			UInt8ToInt             int
			UInt8ToIntAfter        int
			UInt8ToInt8Before      int8
			UInt8ToInt8            int8
			UInt8ToInt8After       int8
			UInt8ToInt16Before     int16
			UInt8ToInt16           int16
			UInt8ToInt16After      int16
			UInt8ToInt32Before     int32
			UInt8ToInt32           int32
			UInt8ToInt32After      int32
			UInt8ToInt64Before     int64
			UInt8ToInt64           int64
			UInt8ToInt64After      int64
			UInt8ToUIntBefore      uint
			UInt8ToUInt            uint
			UInt8ToUIntAfter       uint
			UInt8ToUInt8Before     uint8
			UInt8ToUInt8           uint8
			UInt8ToUInt8After      uint8
			UInt8ToUInt16Before    uint16
			UInt8ToUInt16          uint16
			UInt8ToUInt16After     uint16
			UInt8ToUInt32Before    uint32
			UInt8ToUInt32          uint32
			UInt8ToUInt32After     uint32
			UInt8ToUInt64Before    uint64
			UInt8ToUInt64          uint64
			UInt8ToUInt64After     uint64
			UInt8ToUintptrBefore   uintptr
			UInt8ToUintptr         uintptr
			UInt8ToUintptrAfter    uintptr
			UInt16ToIntBefore      int
			UInt16ToInt            int
			UInt16ToIntAfter       int
			UInt16ToInt8Before     int8
			UInt16ToInt8           int8
			UInt16ToInt8After      int8
			UInt16ToInt16Before    int16
			UInt16ToInt16          int16
			UInt16ToInt16After     int16
			UInt16ToInt32Before    int32
			UInt16ToInt32          int32
			UInt16ToInt32After     int32
			UInt16ToInt64Before    int64
			UInt16ToInt64          int64
			UInt16ToInt64After     int64
			UInt16ToUIntBefore     uint
			UInt16ToUInt           uint
			UInt16ToUIntAfter      uint
			UInt16ToUInt8Before    uint8
			UInt16ToUInt8          uint8
			UInt16ToUInt8After     uint8
			UInt16ToUInt16Before   uint16
			UInt16ToUInt16         uint16
			UInt16ToUInt16After    uint16
			UInt16ToUInt32Before   uint32
			UInt16ToUInt32         uint32
			UInt16ToUInt32After    uint32
			UInt16ToUInt64Before   uint64
			UInt16ToUInt64         uint64
			UInt16ToUInt64After    uint64
			UInt16ToUintptrBefore  uintptr
			UInt16ToUintptr        uintptr
			UInt16ToUintptrAfter   uintptr
			UInt32ToIntBefore      int
			UInt32ToInt            int
			UInt32ToIntAfter       int
			UInt32ToInt8Before     int8
			UInt32ToInt8           int8
			UInt32ToInt8After      int8
			UInt32ToInt16Before    int16
			UInt32ToInt16          int16
			UInt32ToInt16After     int16
			UInt32ToInt32Before    int32
			UInt32ToInt32          int32
			UInt32ToInt32After     int32
			UInt32ToInt64Before    int64
			UInt32ToInt64          int64
			UInt32ToInt64After     int64
			UInt32ToUIntBefore     uint
			UInt32ToUInt           uint
			UInt32ToUIntAfter      uint
			UInt32ToUInt8Before    uint8
			UInt32ToUInt8          uint8
			UInt32ToUInt8After     uint8
			UInt32ToUInt16Before   uint16
			UInt32ToUInt16         uint16
			UInt32ToUInt16After    uint16
			UInt32ToUInt32Before   uint32
			UInt32ToUInt32         uint32
			UInt32ToUInt32After    uint32
			UInt32ToUInt64Before   uint64
			UInt32ToUInt64         uint64
			UInt32ToUInt64After    uint64
			UInt32ToUintptrBefore  uintptr
			UInt32ToUintptr        uintptr
			UInt32ToUintptrAfter   uintptr
			UInt64ToIntBefore      int
			UInt64ToInt            int
			UInt64ToIntAfter       int
			UInt64ToInt8Before     int8
			UInt64ToInt8           int8
			UInt64ToInt8After      int8
			UInt64ToInt16Before    int16
			UInt64ToInt16          int16
			UInt64ToInt16After     int16
			UInt64ToInt32Before    int32
			UInt64ToInt32          int32
			UInt64ToInt32After     int32
			UInt64ToInt64Before    int64
			UInt64ToInt64          int64
			UInt64ToInt64After     int64
			UInt64ToUIntBefore     uint
			UInt64ToUInt           uint
			UInt64ToUIntAfter      uint
			UInt64ToUInt8Before    uint8
			UInt64ToUInt8          uint8
			UInt64ToUInt8After     uint8
			UInt64ToUInt16Before   uint16
			UInt64ToUInt16         uint16
			UInt64ToUInt16After    uint16
			UInt64ToUInt32Before   uint32
			UInt64ToUInt32         uint32
			UInt64ToUInt32After    uint32
			UInt64ToUInt64Before   uint64
			UInt64ToUInt64         uint64
			UInt64ToUInt64After    uint64
			UInt64ToUintptrBefore  uintptr
			UInt64ToUintptr        uintptr
			UInt64ToUintptrAfter   uintptr
			UintptrToIntBefore     int
			UintptrToInt           int
			UintptrToIntAfter      int
			UintptrToInt8Before    int8
			UintptrToInt8          int8
			UintptrToInt8After     int8
			UintptrToInt16Before   int16
			UintptrToInt16         int16
			UintptrToInt16After    int16
			UintptrToInt32Before   int32
			UintptrToInt32         int32
			UintptrToInt32After    int32
			UintptrToInt64Before   int64
			UintptrToInt64         int64
			UintptrToInt64After    int64
			UintptrToUIntBefore    uint
			UintptrToUInt          uint
			UintptrToUIntAfter     uint
			UintptrToUInt8Before   uint8
			UintptrToUInt8         uint8
			UintptrToUInt8After    uint8
			UintptrToUInt16Before  uint16
			UintptrToUInt16        uint16
			UintptrToUInt16After   uint16
			UintptrToUInt32Before  uint32
			UintptrToUInt32        uint32
			UintptrToUInt32After   uint32
			UintptrToUInt64Before  uint64
			UintptrToUInt64        uint64
			UintptrToUInt64After   uint64
			UintptrToUintptrBefore uintptr
			UintptrToUintptr       uintptr
			UintptrToUintptrAfter  uintptr
		}{}
	)

	a := assert.New(t)
	a.NoError(Copy(&dest, src))
	a.Equal(int(0), dest.IntToIntBefore)
	a.Equal(int(18), dest.IntToInt)
	a.Equal(int(0), dest.IntToIntAfter)
	a.Equal(int8(0), dest.IntToInt8Before)
	a.Equal(int8(18), dest.IntToInt8)
	a.Equal(int8(0), dest.IntToInt8After)
	a.Equal(int16(0), dest.IntToInt16Before)
	a.Equal(int16(18), dest.IntToInt16)
	a.Equal(int16(0), dest.IntToInt16After)
	a.Equal(int32(0), dest.IntToInt32Before)
	a.Equal(int32(18), dest.IntToInt32)
	a.Equal(int32(0), dest.IntToInt32After)
	a.Equal(int64(0), dest.IntToInt64Before)
	a.Equal(int64(18), dest.IntToInt64)
	a.Equal(int64(0), dest.IntToInt64After)
	a.Equal(uint(0), dest.IntToUIntBefore)
	a.Equal(uint(18), dest.IntToUInt)
	a.Equal(uint(0), dest.IntToUIntAfter)
	a.Equal(uint8(0), dest.IntToUInt8Before)
	a.Equal(uint8(18), dest.IntToUInt8)
	a.Equal(uint8(0), dest.IntToUInt8After)
	a.Equal(uint16(0), dest.IntToUInt16Before)
	a.Equal(uint16(18), dest.IntToUInt16)
	a.Equal(uint16(0), dest.IntToUInt16After)
	a.Equal(uint32(0), dest.IntToUInt32Before)
	a.Equal(uint32(18), dest.IntToUInt32)
	a.Equal(uint32(0), dest.IntToUInt32After)
	a.Equal(uint64(0), dest.IntToUInt64Before)
	a.Equal(uint64(18), dest.IntToUInt64)
	a.Equal(uint64(0), dest.IntToUInt64After)
	a.Equal(uintptr(0), dest.IntToUintptrBefore)
	a.Equal(uintptr(18), dest.IntToUintptr)
	a.Equal(uintptr(0), dest.IntToUintptrAfter)
	a.Equal(int(0), dest.Int8ToIntBefore)
	a.Equal(int(18), dest.Int8ToInt)
	a.Equal(int(0), dest.Int8ToIntAfter)
	a.Equal(int8(0), dest.Int8ToInt8Before)
	a.Equal(int8(18), dest.Int8ToInt8)
	a.Equal(int8(0), dest.Int8ToInt8After)
	a.Equal(int16(0), dest.Int8ToInt16Before)
	a.Equal(int16(18), dest.Int8ToInt16)
	a.Equal(int16(0), dest.Int8ToInt16After)
	a.Equal(int32(0), dest.Int8ToInt32Before)
	a.Equal(int32(18), dest.Int8ToInt32)
	a.Equal(int32(0), dest.Int8ToInt32After)
	a.Equal(int64(0), dest.Int8ToInt64Before)
	a.Equal(int64(18), dest.Int8ToInt64)
	a.Equal(int64(0), dest.Int8ToInt64After)
	a.Equal(uint(0), dest.Int8ToUIntBefore)
	a.Equal(uint(18), dest.Int8ToUInt)
	a.Equal(uint(0), dest.Int8ToUIntAfter)
	a.Equal(uint8(0), dest.Int8ToUInt8Before)
	a.Equal(uint8(18), dest.Int8ToUInt8)
	a.Equal(uint8(0), dest.Int8ToUInt8After)
	a.Equal(uint16(0), dest.Int8ToUInt16Before)
	a.Equal(uint16(18), dest.Int8ToUInt16)
	a.Equal(uint16(0), dest.Int8ToUInt16After)
	a.Equal(uint32(0), dest.Int8ToUInt32Before)
	a.Equal(uint32(18), dest.Int8ToUInt32)
	a.Equal(uint32(0), dest.Int8ToUInt32After)
	a.Equal(uint64(0), dest.Int8ToUInt64Before)
	a.Equal(uint64(18), dest.Int8ToUInt64)
	a.Equal(uint64(0), dest.Int8ToUInt64After)
	a.Equal(uintptr(0), dest.Int8ToUintptrBefore)
	a.Equal(uintptr(18), dest.Int8ToUintptr)
	a.Equal(uintptr(0), dest.Int8ToUintptrAfter)
	a.Equal(int(0), dest.Int16ToIntBefore)
	a.Equal(int(18), dest.Int16ToInt)
	a.Equal(int(0), dest.Int16ToIntAfter)
	a.Equal(int8(0), dest.Int16ToInt8Before)
	a.Equal(int8(18), dest.Int16ToInt8)
	a.Equal(int8(0), dest.Int16ToInt8After)
	a.Equal(int16(0), dest.Int16ToInt16Before)
	a.Equal(int16(18), dest.Int16ToInt16)
	a.Equal(int16(0), dest.Int16ToInt16After)
	a.Equal(int32(0), dest.Int16ToInt32Before)
	a.Equal(int32(18), dest.Int16ToInt32)
	a.Equal(int32(0), dest.Int16ToInt32After)
	a.Equal(int64(0), dest.Int16ToInt64Before)
	a.Equal(int64(18), dest.Int16ToInt64)
	a.Equal(int64(0), dest.Int16ToInt64After)
	a.Equal(uint(0), dest.Int16ToUIntBefore)
	a.Equal(uint(18), dest.Int16ToUInt)
	a.Equal(uint(0), dest.Int16ToUIntAfter)
	a.Equal(uint8(0), dest.Int16ToUInt8Before)
	a.Equal(uint8(18), dest.Int16ToUInt8)
	a.Equal(uint8(0), dest.Int16ToUInt8After)
	a.Equal(uint16(0), dest.Int16ToUInt16Before)
	a.Equal(uint16(18), dest.Int16ToUInt16)
	a.Equal(uint16(0), dest.Int16ToUInt16After)
	a.Equal(uint32(0), dest.Int16ToUInt32Before)
	a.Equal(uint32(18), dest.Int16ToUInt32)
	a.Equal(uint32(0), dest.Int16ToUInt32After)
	a.Equal(uint64(0), dest.Int16ToUInt64Before)
	a.Equal(uint64(18), dest.Int16ToUInt64)
	a.Equal(uint64(0), dest.Int16ToUInt64After)
	a.Equal(uintptr(0), dest.Int16ToUintptrBefore)
	a.Equal(uintptr(18), dest.Int16ToUintptr)
	a.Equal(uintptr(0), dest.Int16ToUintptrAfter)
	a.Equal(int(0), dest.Int32ToIntBefore)
	a.Equal(int(18), dest.Int32ToInt)
	a.Equal(int(0), dest.Int32ToIntAfter)
	a.Equal(int8(0), dest.Int32ToInt8Before)
	a.Equal(int8(18), dest.Int32ToInt8)
	a.Equal(int8(0), dest.Int32ToInt8After)
	a.Equal(int16(0), dest.Int32ToInt16Before)
	a.Equal(int16(18), dest.Int32ToInt16)
	a.Equal(int16(0), dest.Int32ToInt16After)
	a.Equal(int32(0), dest.Int32ToInt32Before)
	a.Equal(int32(18), dest.Int32ToInt32)
	a.Equal(int32(0), dest.Int32ToInt32After)
	a.Equal(int64(0), dest.Int32ToInt64Before)
	a.Equal(int64(18), dest.Int32ToInt64)
	a.Equal(int64(0), dest.Int32ToInt64After)
	a.Equal(uint(0), dest.Int32ToUIntBefore)
	a.Equal(uint(18), dest.Int32ToUInt)
	a.Equal(uint(0), dest.Int32ToUIntAfter)
	a.Equal(uint8(0), dest.Int32ToUInt8Before)
	a.Equal(uint8(18), dest.Int32ToUInt8)
	a.Equal(uint8(0), dest.Int32ToUInt8After)
	a.Equal(uint16(0), dest.Int32ToUInt16Before)
	a.Equal(uint16(18), dest.Int32ToUInt16)
	a.Equal(uint16(0), dest.Int32ToUInt16After)
	a.Equal(uint32(0), dest.Int32ToUInt32Before)
	a.Equal(uint32(18), dest.Int32ToUInt32)
	a.Equal(uint32(0), dest.Int32ToUInt32After)
	a.Equal(uint64(0), dest.Int32ToUInt64Before)
	a.Equal(uint64(18), dest.Int32ToUInt64)
	a.Equal(uint64(0), dest.Int32ToUInt64After)
	a.Equal(uintptr(0), dest.Int32ToUintptrBefore)
	a.Equal(uintptr(18), dest.Int32ToUintptr)
	a.Equal(uintptr(0), dest.Int32ToUintptrAfter)
	a.Equal(int(0), dest.Int64ToIntBefore)
	a.Equal(int(18), dest.Int64ToInt)
	a.Equal(int(0), dest.Int64ToIntAfter)
	a.Equal(int8(0), dest.Int64ToInt8Before)
	a.Equal(int8(18), dest.Int64ToInt8)
	a.Equal(int8(0), dest.Int64ToInt8After)
	a.Equal(int16(0), dest.Int64ToInt16Before)
	a.Equal(int16(18), dest.Int64ToInt16)
	a.Equal(int16(0), dest.Int64ToInt16After)
	a.Equal(int32(0), dest.Int64ToInt32Before)
	a.Equal(int32(18), dest.Int64ToInt32)
	a.Equal(int32(0), dest.Int64ToInt32After)
	a.Equal(int64(0), dest.Int64ToInt64Before)
	a.Equal(int64(18), dest.Int64ToInt64)
	a.Equal(int64(0), dest.Int64ToInt64After)
	a.Equal(uint(0), dest.Int64ToUIntBefore)
	a.Equal(uint(18), dest.Int64ToUInt)
	a.Equal(uint(0), dest.Int64ToUIntAfter)
	a.Equal(uint8(0), dest.Int64ToUInt8Before)
	a.Equal(uint8(18), dest.Int64ToUInt8)
	a.Equal(uint8(0), dest.Int64ToUInt8After)
	a.Equal(uint16(0), dest.Int64ToUInt16Before)
	a.Equal(uint16(18), dest.Int64ToUInt16)
	a.Equal(uint16(0), dest.Int64ToUInt16After)
	a.Equal(uint32(0), dest.Int64ToUInt32Before)
	a.Equal(uint32(18), dest.Int64ToUInt32)
	a.Equal(uint32(0), dest.Int64ToUInt32After)
	a.Equal(uint64(0), dest.Int64ToUInt64Before)
	a.Equal(uint64(18), dest.Int64ToUInt64)
	a.Equal(uint64(0), dest.Int64ToUInt64After)
	a.Equal(uintptr(0), dest.Int64ToUintptrBefore)
	a.Equal(uintptr(18), dest.Int64ToUintptr)
	a.Equal(uintptr(0), dest.Int64ToUintptrAfter)
	a.Equal(int(0), dest.UIntToIntBefore)
	a.Equal(int(18), dest.UIntToInt)
	a.Equal(int(0), dest.UIntToIntAfter)
	a.Equal(int8(0), dest.UIntToInt8Before)
	a.Equal(int8(18), dest.UIntToInt8)
	a.Equal(int8(0), dest.UIntToInt8After)
	a.Equal(int16(0), dest.UIntToInt16Before)
	a.Equal(int16(18), dest.UIntToInt16)
	a.Equal(int16(0), dest.UIntToInt16After)
	a.Equal(int32(0), dest.UIntToInt32Before)
	a.Equal(int32(18), dest.UIntToInt32)
	a.Equal(int32(0), dest.UIntToInt32After)
	a.Equal(int64(0), dest.UIntToInt64Before)
	a.Equal(int64(18), dest.UIntToInt64)
	a.Equal(int64(0), dest.UIntToInt64After)
	a.Equal(uint(0), dest.UIntToUIntBefore)
	a.Equal(uint(18), dest.UIntToUInt)
	a.Equal(uint(0), dest.UIntToUIntAfter)
	a.Equal(uint8(0), dest.UIntToUInt8Before)
	a.Equal(uint8(18), dest.UIntToUInt8)
	a.Equal(uint8(0), dest.UIntToUInt8After)
	a.Equal(uint16(0), dest.UIntToUInt16Before)
	a.Equal(uint16(18), dest.UIntToUInt16)
	a.Equal(uint16(0), dest.UIntToUInt16After)
	a.Equal(uint32(0), dest.UIntToUInt32Before)
	a.Equal(uint32(18), dest.UIntToUInt32)
	a.Equal(uint32(0), dest.UIntToUInt32After)
	a.Equal(uint64(0), dest.UIntToUInt64Before)
	a.Equal(uint64(18), dest.UIntToUInt64)
	a.Equal(uint64(0), dest.UIntToUInt64After)
	a.Equal(uintptr(0), dest.UIntToUintptrBefore)
	a.Equal(uintptr(18), dest.UIntToUintptr)
	a.Equal(uintptr(0), dest.UIntToUintptrAfter)
	a.Equal(int(0), dest.UInt8ToIntBefore)
	a.Equal(int(18), dest.UInt8ToInt)
	a.Equal(int(0), dest.UInt8ToIntAfter)
	a.Equal(int8(0), dest.UInt8ToInt8Before)
	a.Equal(int8(18), dest.UInt8ToInt8)
	a.Equal(int8(0), dest.UInt8ToInt8After)
	a.Equal(int16(0), dest.UInt8ToInt16Before)
	a.Equal(int16(18), dest.UInt8ToInt16)
	a.Equal(int16(0), dest.UInt8ToInt16After)
	a.Equal(int32(0), dest.UInt8ToInt32Before)
	a.Equal(int32(18), dest.UInt8ToInt32)
	a.Equal(int32(0), dest.UInt8ToInt32After)
	a.Equal(int64(0), dest.UInt8ToInt64Before)
	a.Equal(int64(18), dest.UInt8ToInt64)
	a.Equal(int64(0), dest.UInt8ToInt64After)
	a.Equal(uint(0), dest.UInt8ToUIntBefore)
	a.Equal(uint(18), dest.UInt8ToUInt)
	a.Equal(uint(0), dest.UInt8ToUIntAfter)
	a.Equal(uint8(0), dest.UInt8ToUInt8Before)
	a.Equal(uint8(18), dest.UInt8ToUInt8)
	a.Equal(uint8(0), dest.UInt8ToUInt8After)
	a.Equal(uint16(0), dest.UInt8ToUInt16Before)
	a.Equal(uint16(18), dest.UInt8ToUInt16)
	a.Equal(uint16(0), dest.UInt8ToUInt16After)
	a.Equal(uint32(0), dest.UInt8ToUInt32Before)
	a.Equal(uint32(18), dest.UInt8ToUInt32)
	a.Equal(uint32(0), dest.UInt8ToUInt32After)
	a.Equal(uint64(0), dest.UInt8ToUInt64Before)
	a.Equal(uint64(18), dest.UInt8ToUInt64)
	a.Equal(uint64(0), dest.UInt8ToUInt64After)
	a.Equal(uintptr(0), dest.UInt8ToUintptrBefore)
	a.Equal(uintptr(18), dest.UInt8ToUintptr)
	a.Equal(uintptr(0), dest.UInt8ToUintptrAfter)
	a.Equal(int(0), dest.UInt16ToIntBefore)
	a.Equal(int(18), dest.UInt16ToInt)
	a.Equal(int(0), dest.UInt16ToIntAfter)
	a.Equal(int8(0), dest.UInt16ToInt8Before)
	a.Equal(int8(18), dest.UInt16ToInt8)
	a.Equal(int8(0), dest.UInt16ToInt8After)
	a.Equal(int16(0), dest.UInt16ToInt16Before)
	a.Equal(int16(18), dest.UInt16ToInt16)
	a.Equal(int16(0), dest.UInt16ToInt16After)
	a.Equal(int32(0), dest.UInt16ToInt32Before)
	a.Equal(int32(18), dest.UInt16ToInt32)
	a.Equal(int32(0), dest.UInt16ToInt32After)
	a.Equal(int64(0), dest.UInt16ToInt64Before)
	a.Equal(int64(18), dest.UInt16ToInt64)
	a.Equal(int64(0), dest.UInt16ToInt64After)
	a.Equal(uint(0), dest.UInt16ToUIntBefore)
	a.Equal(uint(18), dest.UInt16ToUInt)
	a.Equal(uint(0), dest.UInt16ToUIntAfter)
	a.Equal(uint8(0), dest.UInt16ToUInt8Before)
	a.Equal(uint8(18), dest.UInt16ToUInt8)
	a.Equal(uint8(0), dest.UInt16ToUInt8After)
	a.Equal(uint16(0), dest.UInt16ToUInt16Before)
	a.Equal(uint16(18), dest.UInt16ToUInt16)
	a.Equal(uint16(0), dest.UInt16ToUInt16After)
	a.Equal(uint32(0), dest.UInt16ToUInt32Before)
	a.Equal(uint32(18), dest.UInt16ToUInt32)
	a.Equal(uint32(0), dest.UInt16ToUInt32After)
	a.Equal(uint64(0), dest.UInt16ToUInt64Before)
	a.Equal(uint64(18), dest.UInt16ToUInt64)
	a.Equal(uint64(0), dest.UInt16ToUInt64After)
	a.Equal(uintptr(0), dest.UInt16ToUintptrBefore)
	a.Equal(uintptr(18), dest.UInt16ToUintptr)
	a.Equal(uintptr(0), dest.UInt16ToUintptrAfter)
	a.Equal(int(0), dest.UInt32ToIntBefore)
	a.Equal(int(18), dest.UInt32ToInt)
	a.Equal(int(0), dest.UInt32ToIntAfter)
	a.Equal(int8(0), dest.UInt32ToInt8Before)
	a.Equal(int8(18), dest.UInt32ToInt8)
	a.Equal(int8(0), dest.UInt32ToInt8After)
	a.Equal(int16(0), dest.UInt32ToInt16Before)
	a.Equal(int16(18), dest.UInt32ToInt16)
	a.Equal(int16(0), dest.UInt32ToInt16After)
	a.Equal(int32(0), dest.UInt32ToInt32Before)
	a.Equal(int32(18), dest.UInt32ToInt32)
	a.Equal(int32(0), dest.UInt32ToInt32After)
	a.Equal(int64(0), dest.UInt32ToInt64Before)
	a.Equal(int64(18), dest.UInt32ToInt64)
	a.Equal(int64(0), dest.UInt32ToInt64After)
	a.Equal(uint(0), dest.UInt32ToUIntBefore)
	a.Equal(uint(18), dest.UInt32ToUInt)
	a.Equal(uint(0), dest.UInt32ToUIntAfter)
	a.Equal(uint8(0), dest.UInt32ToUInt8Before)
	a.Equal(uint8(18), dest.UInt32ToUInt8)
	a.Equal(uint8(0), dest.UInt32ToUInt8After)
	a.Equal(uint16(0), dest.UInt32ToUInt16Before)
	a.Equal(uint16(18), dest.UInt32ToUInt16)
	a.Equal(uint16(0), dest.UInt32ToUInt16After)
	a.Equal(uint32(0), dest.UInt32ToUInt32Before)
	a.Equal(uint32(18), dest.UInt32ToUInt32)
	a.Equal(uint32(0), dest.UInt32ToUInt32After)
	a.Equal(uint64(0), dest.UInt32ToUInt64Before)
	a.Equal(uint64(18), dest.UInt32ToUInt64)
	a.Equal(uint64(0), dest.UInt32ToUInt64After)
	a.Equal(uintptr(0), dest.UInt32ToUintptrBefore)
	a.Equal(uintptr(18), dest.UInt32ToUintptr)
	a.Equal(uintptr(0), dest.UInt32ToUintptrAfter)
	a.Equal(int(0), dest.UInt64ToIntBefore)
	a.Equal(int(18), dest.UInt64ToInt)
	a.Equal(int(0), dest.UInt64ToIntAfter)
	a.Equal(int8(0), dest.UInt64ToInt8Before)
	a.Equal(int8(18), dest.UInt64ToInt8)
	a.Equal(int8(0), dest.UInt64ToInt8After)
	a.Equal(int16(0), dest.UInt64ToInt16Before)
	a.Equal(int16(18), dest.UInt64ToInt16)
	a.Equal(int16(0), dest.UInt64ToInt16After)
	a.Equal(int32(0), dest.UInt64ToInt32Before)
	a.Equal(int32(18), dest.UInt64ToInt32)
	a.Equal(int32(0), dest.UInt64ToInt32After)
	a.Equal(int64(0), dest.UInt64ToInt64Before)
	a.Equal(int64(18), dest.UInt64ToInt64)
	a.Equal(int64(0), dest.UInt64ToInt64After)
	a.Equal(uint(0), dest.UInt64ToUIntBefore)
	a.Equal(uint(18), dest.UInt64ToUInt)
	a.Equal(uint(0), dest.UInt64ToUIntAfter)
	a.Equal(uint8(0), dest.UInt64ToUInt8Before)
	a.Equal(uint8(18), dest.UInt64ToUInt8)
	a.Equal(uint8(0), dest.UInt64ToUInt8After)
	a.Equal(uint16(0), dest.UInt64ToUInt16Before)
	a.Equal(uint16(18), dest.UInt64ToUInt16)
	a.Equal(uint16(0), dest.UInt64ToUInt16After)
	a.Equal(uint32(0), dest.UInt64ToUInt32Before)
	a.Equal(uint32(18), dest.UInt64ToUInt32)
	a.Equal(uint32(0), dest.UInt64ToUInt32After)
	a.Equal(uint64(0), dest.UInt64ToUInt64Before)
	a.Equal(uint64(18), dest.UInt64ToUInt64)
	a.Equal(uint64(0), dest.UInt64ToUInt64After)
	a.Equal(uintptr(0), dest.UInt64ToUintptrBefore)
	a.Equal(uintptr(18), dest.UInt64ToUintptr)
	a.Equal(uintptr(0), dest.UInt64ToUintptrAfter)
	a.Equal(int(0), dest.UintptrToIntBefore)
	a.Equal(int(18), dest.UintptrToInt)
	a.Equal(int(0), dest.UintptrToIntAfter)
	a.Equal(int8(0), dest.UintptrToInt8Before)
	a.Equal(int8(18), dest.UintptrToInt8)
	a.Equal(int8(0), dest.UintptrToInt8After)
	a.Equal(int16(0), dest.UintptrToInt16Before)
	a.Equal(int16(18), dest.UintptrToInt16)
	a.Equal(int16(0), dest.UintptrToInt16After)
	a.Equal(int32(0), dest.UintptrToInt32Before)
	a.Equal(int32(18), dest.UintptrToInt32)
	a.Equal(int32(0), dest.UintptrToInt32After)
	a.Equal(int64(0), dest.UintptrToInt64Before)
	a.Equal(int64(18), dest.UintptrToInt64)
	a.Equal(int64(0), dest.UintptrToInt64After)
	a.Equal(uint(0), dest.UintptrToUIntBefore)
	a.Equal(uint(18), dest.UintptrToUInt)
	a.Equal(uint(0), dest.UintptrToUIntAfter)
	a.Equal(uint8(0), dest.UintptrToUInt8Before)
	a.Equal(uint8(18), dest.UintptrToUInt8)
	a.Equal(uint8(0), dest.UintptrToUInt8After)
	a.Equal(uint16(0), dest.UintptrToUInt16Before)
	a.Equal(uint16(18), dest.UintptrToUInt16)
	a.Equal(uint16(0), dest.UintptrToUInt16After)
	a.Equal(uint32(0), dest.UintptrToUInt32Before)
	a.Equal(uint32(18), dest.UintptrToUInt32)
	a.Equal(uint32(0), dest.UintptrToUInt32After)
	a.Equal(uint64(0), dest.UintptrToUInt64Before)
	a.Equal(uint64(18), dest.UintptrToUInt64)
	a.Equal(uint64(0), dest.UintptrToUInt64After)
	a.Equal(uintptr(0), dest.UintptrToUintptrBefore)
	a.Equal(uintptr(18), dest.UintptrToUintptr)
	a.Equal(uintptr(0), dest.UintptrToUintptrAfter)
}

func TestGenerateIntsTestingCode(t *testing.T) {
	t.Skip()
	ints := []string{
		"Int",
		"Int8",
		"Int16",
		"Int32",
		"Int64",
		"UInt",
		"UInt8",
		"UInt16",
		"UInt32",
		"UInt64",
		"Uintptr",
	}

	for _, s1 := range ints {
		for _, s2 := range ints {
			typ := strings.ToLower(s2)
			fmt.Printf("a.Equal(%s(0), dest.%sTo%sBefore )\n", typ, s1, s2)
			fmt.Printf("a.Equal(%s(18), dest.%sTo%s)\n", typ, s1, s2)
			fmt.Printf("a.Equal(%s(0), dest.%sTo%sAfter )\n", typ, s1, s2)
		}
	}
}
