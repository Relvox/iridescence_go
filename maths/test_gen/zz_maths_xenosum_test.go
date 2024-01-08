package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/relvox/iridescence_go/maths"
)

func Test_XenoSum(t *testing.T) {
	t.Run("uint", func(t *testing.T) {
		var a, b, c uint = 4, 4, 4
		if !assert.Equal(t, uint(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint8", func(t *testing.T) {
		var a, b, c uint8 = 4, 4, 4
		if !assert.Equal(t, uint8(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint16", func(t *testing.T) {
		var a, b, c uint16 = 4, 4, 4
		if !assert.Equal(t, uint16(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var a, b, c uint32 = 4, 4, 4
		if !assert.Equal(t, uint32(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint64", func(t *testing.T) {
		var a, b, c uint64 = 4, 4, 4
		if !assert.Equal(t, uint64(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int", func(t *testing.T) {
		var a, b, c int = 4, 4, 4
		if !assert.Equal(t, int(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int8", func(t *testing.T) {
		var a, b, c int8 = 4, 4, 4
		if !assert.Equal(t, int8(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int16", func(t *testing.T) {
		var a, b, c int16 = 4, 4, 4
		if !assert.Equal(t, int16(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int32", func(t *testing.T) {
		var a, b, c int32 = 4, 4, 4
		if !assert.Equal(t, int32(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int64", func(t *testing.T) {
		var a, b, c int64 = 4, 4, 4
		if !assert.Equal(t, int64(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float32", func(t *testing.T) {
		var a, b, c float32 = 4.0, 4.0, 4.0
		if !assert.Equal(t, float32(7.0), maths.XenoSum(a, b, c), "Xeno sum of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float64", func(t *testing.T) {
		var a, b, c float64 = 4.0, 4.0, 4.0
		if !assert.Equal(t, float64(7.0), maths.XenoSum(a, b, c), "Xeno sum of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

}
