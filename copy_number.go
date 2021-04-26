package copy

import (
	"reflect"
	"unsafe"
)

func copyNumber(dest, src unsafe.Pointer, destKind, srcKind reflect.Kind) bool {
	switch destKind {
 		case reflect.Int8:
			switch srcKind {
			case reflect.Int8:
				*((*int8)(dest)) = int8(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*int8)(dest)) = int8(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*int8)(dest)) = int8(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*int8)(dest)) = int8(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*int8)(dest)) = int8(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*int8)(dest)) = int8(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*int8)(dest)) = int8(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*int8)(dest)) = int8(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*int8)(dest)) = int8(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*int8)(dest)) = int8(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*int8)(dest)) = int8(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*int8)(dest)) = int8(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*int8)(dest)) = int8(*((*float64)(src)))
				return true

			}
 		case reflect.Int16:
			switch srcKind {
			case reflect.Int8:
				*((*int16)(dest)) = int16(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*int16)(dest)) = int16(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*int16)(dest)) = int16(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*int16)(dest)) = int16(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*int16)(dest)) = int16(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*int16)(dest)) = int16(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*int16)(dest)) = int16(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*int16)(dest)) = int16(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*int16)(dest)) = int16(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*int16)(dest)) = int16(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*int16)(dest)) = int16(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*int16)(dest)) = int16(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*int16)(dest)) = int16(*((*float64)(src)))
				return true

			}
 		case reflect.Int32:
			switch srcKind {
			case reflect.Int8:
				*((*int32)(dest)) = int32(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*int32)(dest)) = int32(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*int32)(dest)) = int32(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*int32)(dest)) = int32(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*int32)(dest)) = int32(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*int32)(dest)) = int32(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*int32)(dest)) = int32(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*int32)(dest)) = int32(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*int32)(dest)) = int32(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*int32)(dest)) = int32(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*int32)(dest)) = int32(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*int32)(dest)) = int32(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*int32)(dest)) = int32(*((*float64)(src)))
				return true

			}
 		case reflect.Int64:
			switch srcKind {
			case reflect.Int8:
				*((*int64)(dest)) = int64(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*int64)(dest)) = int64(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*int64)(dest)) = int64(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*int64)(dest)) = int64(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*int64)(dest)) = int64(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*int64)(dest)) = int64(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*int64)(dest)) = int64(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*int64)(dest)) = int64(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*int64)(dest)) = int64(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*int64)(dest)) = int64(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*int64)(dest)) = int64(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*int64)(dest)) = int64(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*int64)(dest)) = int64(*((*float64)(src)))
				return true

			}
 		case reflect.Int:
			switch srcKind {
			case reflect.Int8:
				*((*int)(dest)) = int(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*int)(dest)) = int(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*int)(dest)) = int(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*int)(dest)) = int(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*int)(dest)) = int(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*int)(dest)) = int(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*int)(dest)) = int(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*int)(dest)) = int(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*int)(dest)) = int(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*int)(dest)) = int(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*int)(dest)) = int(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*int)(dest)) = int(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*int)(dest)) = int(*((*float64)(src)))
				return true

			}
 		case reflect.Uint8:
			switch srcKind {
			case reflect.Int8:
				*((*uint8)(dest)) = uint8(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uint8)(dest)) = uint8(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uint8)(dest)) = uint8(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uint8)(dest)) = uint8(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uint8)(dest)) = uint8(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uint8)(dest)) = uint8(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uint8)(dest)) = uint8(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uint8)(dest)) = uint8(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uint8)(dest)) = uint8(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uint8)(dest)) = uint8(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uint8)(dest)) = uint8(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uint8)(dest)) = uint8(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uint8)(dest)) = uint8(*((*float64)(src)))
				return true

			}
 		case reflect.Uint16:
			switch srcKind {
			case reflect.Int8:
				*((*uint16)(dest)) = uint16(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uint16)(dest)) = uint16(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uint16)(dest)) = uint16(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uint16)(dest)) = uint16(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uint16)(dest)) = uint16(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uint16)(dest)) = uint16(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uint16)(dest)) = uint16(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uint16)(dest)) = uint16(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uint16)(dest)) = uint16(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uint16)(dest)) = uint16(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uint16)(dest)) = uint16(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uint16)(dest)) = uint16(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uint16)(dest)) = uint16(*((*float64)(src)))
				return true

			}
 		case reflect.Uint32:
			switch srcKind {
			case reflect.Int8:
				*((*uint32)(dest)) = uint32(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uint32)(dest)) = uint32(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uint32)(dest)) = uint32(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uint32)(dest)) = uint32(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uint32)(dest)) = uint32(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uint32)(dest)) = uint32(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uint32)(dest)) = uint32(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uint32)(dest)) = uint32(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uint32)(dest)) = uint32(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uint32)(dest)) = uint32(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uint32)(dest)) = uint32(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uint32)(dest)) = uint32(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uint32)(dest)) = uint32(*((*float64)(src)))
				return true

			}
 		case reflect.Uint64:
			switch srcKind {
			case reflect.Int8:
				*((*uint64)(dest)) = uint64(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uint64)(dest)) = uint64(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uint64)(dest)) = uint64(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uint64)(dest)) = uint64(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uint64)(dest)) = uint64(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uint64)(dest)) = uint64(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uint64)(dest)) = uint64(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uint64)(dest)) = uint64(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uint64)(dest)) = uint64(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uint64)(dest)) = uint64(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uint64)(dest)) = uint64(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uint64)(dest)) = uint64(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uint64)(dest)) = uint64(*((*float64)(src)))
				return true

			}
 		case reflect.Uint:
			switch srcKind {
			case reflect.Int8:
				*((*uint)(dest)) = uint(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uint)(dest)) = uint(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uint)(dest)) = uint(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uint)(dest)) = uint(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uint)(dest)) = uint(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uint)(dest)) = uint(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uint)(dest)) = uint(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uint)(dest)) = uint(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uint)(dest)) = uint(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uint)(dest)) = uint(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uint)(dest)) = uint(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uint)(dest)) = uint(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uint)(dest)) = uint(*((*float64)(src)))
				return true

			}
 		case reflect.Uintptr:
			switch srcKind {
			case reflect.Int8:
				*((*uintptr)(dest)) = uintptr(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*uintptr)(dest)) = uintptr(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*uintptr)(dest)) = uintptr(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*uintptr)(dest)) = uintptr(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*uintptr)(dest)) = uintptr(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*uintptr)(dest)) = uintptr(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*uintptr)(dest)) = uintptr(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*uintptr)(dest)) = uintptr(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*uintptr)(dest)) = uintptr(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*uintptr)(dest)) = uintptr(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*uintptr)(dest)) = uintptr(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*uintptr)(dest)) = uintptr(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*uintptr)(dest)) = uintptr(*((*float64)(src)))
				return true

			}
 		case reflect.Float32:
			switch srcKind {
			case reflect.Int8:
				*((*float32)(dest)) = float32(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*float32)(dest)) = float32(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*float32)(dest)) = float32(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*float32)(dest)) = float32(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*float32)(dest)) = float32(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*float32)(dest)) = float32(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*float32)(dest)) = float32(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*float32)(dest)) = float32(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*float32)(dest)) = float32(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*float32)(dest)) = float32(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*float32)(dest)) = float32(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*float32)(dest)) = float32(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*float32)(dest)) = float32(*((*float64)(src)))
				return true

			}
 		case reflect.Float64:
			switch srcKind {
			case reflect.Int8:
				*((*float64)(dest)) = float64(*((*int8)(src)))
				return true
			case reflect.Int16:
				*((*float64)(dest)) = float64(*((*int16)(src)))
				return true
			case reflect.Int32:
				*((*float64)(dest)) = float64(*((*int32)(src)))
				return true
			case reflect.Int64:
				*((*float64)(dest)) = float64(*((*int64)(src)))
				return true
			case reflect.Int:
				*((*float64)(dest)) = float64(*((*int)(src)))
				return true
			case reflect.Uint8:
				*((*float64)(dest)) = float64(*((*uint8)(src)))
				return true
			case reflect.Uint16:
				*((*float64)(dest)) = float64(*((*uint16)(src)))
				return true
			case reflect.Uint32:
				*((*float64)(dest)) = float64(*((*uint32)(src)))
				return true
			case reflect.Uint64:
				*((*float64)(dest)) = float64(*((*uint64)(src)))
				return true
			case reflect.Uint:
				*((*float64)(dest)) = float64(*((*uint)(src)))
				return true
			case reflect.Uintptr:
				*((*float64)(dest)) = float64(*((*uintptr)(src)))
				return true
			case reflect.Float32:
				*((*float64)(dest)) = float64(*((*float32)(src)))
				return true
			case reflect.Float64:
				*((*float64)(dest)) = float64(*((*float64)(src)))
				return true

			}

	}
	return false
}
