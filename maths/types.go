package maths

import (
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func MaxValue[T any]() T {
	var val T
	switch any(val).(type) {
	case uint8, uint16, uint32, uint64, uint:
		var maxUintVal uint64 = ^uint64(0)
		return *(*T)(unsafe.Pointer(&maxUintVal))

	case int8, int16, int32, int64, int:
		var maxIntVal int64 = math.MaxInt8
		size := sizeof[T]()
		for i := 1; i < size; i++ {
			maxIntVal <<= 8
			maxIntVal += 0xFF
		}
		return *(*T)(unsafe.Pointer(&maxIntVal))

	case float32:
		var maxFloat32Val float32 = math.MaxFloat32
		return *(*T)(unsafe.Pointer(&maxFloat32Val))

	case float64:
		var maxFloat64Val float64 = math.MaxFloat64
		return *(*T)(unsafe.Pointer(&maxFloat64Val))
	default:
		panic("unsupported type")
	}
}

func MinValue[T any]() T {
	var val T
	switch any(val).(type) {
	case uint8, uint16, uint32, uint64, uint:
		var minUintVal uint64 = uint64(0)
		return *(*T)(unsafe.Pointer(&minUintVal))

	case int8, int16, int32, int64, int:
		var minIntVal int64 = math.MinInt8
		size := sizeof[T]()
		for i := 1; i < size; i++ {
			minIntVal <<= 8
		}
		return *(*T)(unsafe.Pointer(&minIntVal))

	case float32:
		var minFloat32Val float32 = -math.MaxFloat32
		return *(*T)(unsafe.Pointer(&minFloat32Val))

	case float64:
		var minFloat64Val float64 = -math.MaxFloat64
		return *(*T)(unsafe.Pointer(&minFloat64Val))

	default:
		panic("unsupported type")
	}
}

func EpsilonValue[T any]() T {
	var val T
	switch any(val).(type) {
	case uint8, uint16, uint32, uint64, uint:
		var epsilonUintVal uint64 = 1
		return *(*T)(unsafe.Pointer(&epsilonUintVal))

	case int8, int16, int32, int64, int:
		var epsilonIntVal uint64 = 1
		return *(*T)(unsafe.Pointer(&epsilonIntVal))

	case float32:
		var epsilonFloat32Val float32 = math.SmallestNonzeroFloat32
		return *(*T)(unsafe.Pointer(&epsilonFloat32Val))

	case float64:
		var epsilonFloat64Val float64 = math.SmallestNonzeroFloat64
		return *(*T)(unsafe.Pointer(&epsilonFloat64Val))

	default:
		panic("unsupported type")
	}
}

func One[T any]() T {
	var val T
	switch any(val).(type) {
	case uint8, uint16, uint32, uint64, uint:
		var oneUintVal uint64 = uint64(1)
		return *(*T)(unsafe.Pointer(&oneUintVal))

	case int8, int16, int32, int64, int:
		var oneIntVal int64 = 1
		return *(*T)(unsafe.Pointer(&oneIntVal))

	case float32:
		var oneFloat32Val float32 = 1
		return *(*T)(unsafe.Pointer(&oneFloat32Val))

	case float64:
		var oneFloat64Val float64 = 1
		return *(*T)(unsafe.Pointer(&oneFloat64Val))

	default:
		panic("unsupported type")
	}
}

func sizeof[T any]() int {
	var val T
	return int(unsafe.Sizeof(val))
}
