package test

import (
	"testing"

	"github.com/relvox/iridescence_go/maths/integer/intervals"
)

func Test_Intervals(t *testing.T) {
	Intervals_Tester[uint8](t)
	Intervals_Tester[uint16](t)
	Intervals_Tester[uint32](t)
	Intervals_Tester[uint64](t)
	Intervals_Tester[uint](t)

	Intervals_Tester[int8](t)
	Intervals_Tester[int16](t)
	Intervals_Tester[int32](t)
	Intervals_Tester[int64](t)
	Intervals_Tester[int](t)

	// Intervals_Tester[float32](t)
	// Intervals_Tester[float64](t)
}

func Intervals_Tester[T intervals.Number](t *testing.T) {
	NullInterval_Tester[T](t)
	SingletonInterval_Tester[T](t)
	// ClosedInterval_Tester[T](t)
}
