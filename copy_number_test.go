package copy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyNumber(t *testing.T) {

	t.Run("TestInt8ToInt8", func(t *testing.T) {
		var src int8 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestInt8ToInt16", func(t *testing.T) {
		var src int8 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestInt8ToInt32", func(t *testing.T) {
		var src int8 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestInt8ToInt64", func(t *testing.T) {
		var src int8 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestInt8ToInt", func(t *testing.T) {
		var src int8 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestInt8ToUint8", func(t *testing.T) {
		var src int8 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestInt8ToUint16", func(t *testing.T) {
		var src int8 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestInt8ToUint32", func(t *testing.T) {
		var src int8 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestInt8ToUint64", func(t *testing.T) {
		var src int8 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestInt8ToUint", func(t *testing.T) {
		var src int8 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestInt8ToUintptr", func(t *testing.T) {
		var src int8 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestInt8ToFloat32", func(t *testing.T) {
		var src int8 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestInt8ToFloat64", func(t *testing.T) {
		var src int8 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestInt16ToInt8", func(t *testing.T) {
		var src int16 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestInt16ToInt16", func(t *testing.T) {
		var src int16 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestInt16ToInt32", func(t *testing.T) {
		var src int16 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestInt16ToInt64", func(t *testing.T) {
		var src int16 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestInt16ToInt", func(t *testing.T) {
		var src int16 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestInt16ToUint8", func(t *testing.T) {
		var src int16 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestInt16ToUint16", func(t *testing.T) {
		var src int16 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestInt16ToUint32", func(t *testing.T) {
		var src int16 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestInt16ToUint64", func(t *testing.T) {
		var src int16 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestInt16ToUint", func(t *testing.T) {
		var src int16 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestInt16ToUintptr", func(t *testing.T) {
		var src int16 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestInt16ToFloat32", func(t *testing.T) {
		var src int16 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestInt16ToFloat64", func(t *testing.T) {
		var src int16 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestInt32ToInt8", func(t *testing.T) {
		var src int32 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestInt32ToInt16", func(t *testing.T) {
		var src int32 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestInt32ToInt32", func(t *testing.T) {
		var src int32 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestInt32ToInt64", func(t *testing.T) {
		var src int32 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestInt32ToInt", func(t *testing.T) {
		var src int32 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestInt32ToUint8", func(t *testing.T) {
		var src int32 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestInt32ToUint16", func(t *testing.T) {
		var src int32 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestInt32ToUint32", func(t *testing.T) {
		var src int32 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestInt32ToUint64", func(t *testing.T) {
		var src int32 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestInt32ToUint", func(t *testing.T) {
		var src int32 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestInt32ToUintptr", func(t *testing.T) {
		var src int32 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestInt32ToFloat32", func(t *testing.T) {
		var src int32 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestInt32ToFloat64", func(t *testing.T) {
		var src int32 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestInt64ToInt8", func(t *testing.T) {
		var src int64 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestInt64ToInt16", func(t *testing.T) {
		var src int64 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestInt64ToInt32", func(t *testing.T) {
		var src int64 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestInt64ToInt64", func(t *testing.T) {
		var src int64 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestInt64ToInt", func(t *testing.T) {
		var src int64 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestInt64ToUint8", func(t *testing.T) {
		var src int64 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestInt64ToUint16", func(t *testing.T) {
		var src int64 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestInt64ToUint32", func(t *testing.T) {
		var src int64 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestInt64ToUint64", func(t *testing.T) {
		var src int64 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestInt64ToUint", func(t *testing.T) {
		var src int64 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestInt64ToUintptr", func(t *testing.T) {
		var src int64 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestInt64ToFloat32", func(t *testing.T) {
		var src int64 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestInt64ToFloat64", func(t *testing.T) {
		var src int64 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestIntToInt8", func(t *testing.T) {
		var src int = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestIntToInt16", func(t *testing.T) {
		var src int = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestIntToInt32", func(t *testing.T) {
		var src int = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestIntToInt64", func(t *testing.T) {
		var src int = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestIntToInt", func(t *testing.T) {
		var src int = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestIntToUint8", func(t *testing.T) {
		var src int = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestIntToUint16", func(t *testing.T) {
		var src int = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestIntToUint32", func(t *testing.T) {
		var src int = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestIntToUint64", func(t *testing.T) {
		var src int = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestIntToUint", func(t *testing.T) {
		var src int = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestIntToUintptr", func(t *testing.T) {
		var src int = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestIntToFloat32", func(t *testing.T) {
		var src int = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestIntToFloat64", func(t *testing.T) {
		var src int = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUint8ToInt8", func(t *testing.T) {
		var src uint8 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUint8ToInt16", func(t *testing.T) {
		var src uint8 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUint8ToInt32", func(t *testing.T) {
		var src uint8 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUint8ToInt64", func(t *testing.T) {
		var src uint8 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUint8ToInt", func(t *testing.T) {
		var src uint8 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUint8ToUint8", func(t *testing.T) {
		var src uint8 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUint8ToUint16", func(t *testing.T) {
		var src uint8 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUint8ToUint32", func(t *testing.T) {
		var src uint8 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUint8ToUint64", func(t *testing.T) {
		var src uint8 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUint8ToUint", func(t *testing.T) {
		var src uint8 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUint8ToUintptr", func(t *testing.T) {
		var src uint8 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUint8ToFloat32", func(t *testing.T) {
		var src uint8 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUint8ToFloat64", func(t *testing.T) {
		var src uint8 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUint16ToInt8", func(t *testing.T) {
		var src uint16 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUint16ToInt16", func(t *testing.T) {
		var src uint16 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUint16ToInt32", func(t *testing.T) {
		var src uint16 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUint16ToInt64", func(t *testing.T) {
		var src uint16 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUint16ToInt", func(t *testing.T) {
		var src uint16 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUint16ToUint8", func(t *testing.T) {
		var src uint16 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUint16ToUint16", func(t *testing.T) {
		var src uint16 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUint16ToUint32", func(t *testing.T) {
		var src uint16 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUint16ToUint64", func(t *testing.T) {
		var src uint16 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUint16ToUint", func(t *testing.T) {
		var src uint16 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUint16ToUintptr", func(t *testing.T) {
		var src uint16 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUint16ToFloat32", func(t *testing.T) {
		var src uint16 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUint16ToFloat64", func(t *testing.T) {
		var src uint16 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUint32ToInt8", func(t *testing.T) {
		var src uint32 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUint32ToInt16", func(t *testing.T) {
		var src uint32 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUint32ToInt32", func(t *testing.T) {
		var src uint32 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUint32ToInt64", func(t *testing.T) {
		var src uint32 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUint32ToInt", func(t *testing.T) {
		var src uint32 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUint32ToUint8", func(t *testing.T) {
		var src uint32 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUint32ToUint16", func(t *testing.T) {
		var src uint32 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUint32ToUint32", func(t *testing.T) {
		var src uint32 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUint32ToUint64", func(t *testing.T) {
		var src uint32 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUint32ToUint", func(t *testing.T) {
		var src uint32 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUint32ToUintptr", func(t *testing.T) {
		var src uint32 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUint32ToFloat32", func(t *testing.T) {
		var src uint32 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUint32ToFloat64", func(t *testing.T) {
		var src uint32 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUint64ToInt8", func(t *testing.T) {
		var src uint64 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUint64ToInt16", func(t *testing.T) {
		var src uint64 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUint64ToInt32", func(t *testing.T) {
		var src uint64 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUint64ToInt64", func(t *testing.T) {
		var src uint64 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUint64ToInt", func(t *testing.T) {
		var src uint64 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUint64ToUint8", func(t *testing.T) {
		var src uint64 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUint64ToUint16", func(t *testing.T) {
		var src uint64 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUint64ToUint32", func(t *testing.T) {
		var src uint64 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUint64ToUint64", func(t *testing.T) {
		var src uint64 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUint64ToUint", func(t *testing.T) {
		var src uint64 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUint64ToUintptr", func(t *testing.T) {
		var src uint64 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUint64ToFloat32", func(t *testing.T) {
		var src uint64 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUint64ToFloat64", func(t *testing.T) {
		var src uint64 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUintToInt8", func(t *testing.T) {
		var src uint = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUintToInt16", func(t *testing.T) {
		var src uint = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUintToInt32", func(t *testing.T) {
		var src uint = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUintToInt64", func(t *testing.T) {
		var src uint = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUintToInt", func(t *testing.T) {
		var src uint = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUintToUint8", func(t *testing.T) {
		var src uint = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUintToUint16", func(t *testing.T) {
		var src uint = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUintToUint32", func(t *testing.T) {
		var src uint = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUintToUint64", func(t *testing.T) {
		var src uint = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUintToUint", func(t *testing.T) {
		var src uint = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUintToUintptr", func(t *testing.T) {
		var src uint = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUintToFloat32", func(t *testing.T) {
		var src uint = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUintToFloat64", func(t *testing.T) {
		var src uint = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestUintptrToInt8", func(t *testing.T) {
		var src uintptr = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestUintptrToInt16", func(t *testing.T) {
		var src uintptr = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestUintptrToInt32", func(t *testing.T) {
		var src uintptr = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestUintptrToInt64", func(t *testing.T) {
		var src uintptr = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestUintptrToInt", func(t *testing.T) {
		var src uintptr = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestUintptrToUint8", func(t *testing.T) {
		var src uintptr = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestUintptrToUint16", func(t *testing.T) {
		var src uintptr = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestUintptrToUint32", func(t *testing.T) {
		var src uintptr = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestUintptrToUint64", func(t *testing.T) {
		var src uintptr = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestUintptrToUint", func(t *testing.T) {
		var src uintptr = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestUintptrToUintptr", func(t *testing.T) {
		var src uintptr = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestUintptrToFloat32", func(t *testing.T) {
		var src uintptr = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestUintptrToFloat64", func(t *testing.T) {
		var src uintptr = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestFloat32ToInt8", func(t *testing.T) {
		var src float32 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestFloat32ToInt16", func(t *testing.T) {
		var src float32 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestFloat32ToInt32", func(t *testing.T) {
		var src float32 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestFloat32ToInt64", func(t *testing.T) {
		var src float32 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestFloat32ToInt", func(t *testing.T) {
		var src float32 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestFloat32ToUint8", func(t *testing.T) {
		var src float32 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestFloat32ToUint16", func(t *testing.T) {
		var src float32 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestFloat32ToUint32", func(t *testing.T) {
		var src float32 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestFloat32ToUint64", func(t *testing.T) {
		var src float32 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestFloat32ToUint", func(t *testing.T) {
		var src float32 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestFloat32ToUintptr", func(t *testing.T) {
		var src float32 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestFloat32ToFloat32", func(t *testing.T) {
		var src float32 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestFloat32ToFloat64", func(t *testing.T) {
		var src float32 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

	t.Run("TestFloat64ToInt8", func(t *testing.T) {
		var src float64 = 18
		var dest int8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int8(src), dest)
	})

	t.Run("TestFloat64ToInt16", func(t *testing.T) {
		var src float64 = 18
		var dest int16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int16(src), dest)
	})

	t.Run("TestFloat64ToInt32", func(t *testing.T) {
		var src float64 = 18
		var dest int32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int32(src), dest)
	})

	t.Run("TestFloat64ToInt64", func(t *testing.T) {
		var src float64 = 18
		var dest int64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int64(src), dest)
	})

	t.Run("TestFloat64ToInt", func(t *testing.T) {
		var src float64 = 18
		var dest int = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, int(src), dest)
	})

	t.Run("TestFloat64ToUint8", func(t *testing.T) {
		var src float64 = 18
		var dest uint8 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint8(src), dest)
	})

	t.Run("TestFloat64ToUint16", func(t *testing.T) {
		var src float64 = 18
		var dest uint16 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint16(src), dest)
	})

	t.Run("TestFloat64ToUint32", func(t *testing.T) {
		var src float64 = 18
		var dest uint32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint32(src), dest)
	})

	t.Run("TestFloat64ToUint64", func(t *testing.T) {
		var src float64 = 18
		var dest uint64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint64(src), dest)
	})

	t.Run("TestFloat64ToUint", func(t *testing.T) {
		var src float64 = 18
		var dest uint = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uint(src), dest)
	})

	t.Run("TestFloat64ToUintptr", func(t *testing.T) {
		var src float64 = 18
		var dest uintptr = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, uintptr(src), dest)
	})

	t.Run("TestFloat64ToFloat32", func(t *testing.T) {
		var src float64 = 18
		var dest float32 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float32(src), dest)
	})

	t.Run("TestFloat64ToFloat64", func(t *testing.T) {
		var src float64 = 18
		var dest float64 = 10
		assert.NoError(t, Copy(&dest, src))
		assert.Equal(t, float64(src), dest)
	})

}
