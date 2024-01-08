package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/relvox/iridescence_go/maths"
)

func Test_Sum(t *testing.T) {
	t.Run("uint", func(t *testing.T) {
		var a, b, c uint = 10, 20, 30
		if !assert.Equal(t, uint(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint8", func(t *testing.T) {
		var a, b, c uint8 = 10, 20, 30
		if !assert.Equal(t, uint8(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint16", func(t *testing.T) {
		var a, b, c uint16 = 10, 20, 30
		if !assert.Equal(t, uint16(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var a, b, c uint32 = 10, 20, 30
		if !assert.Equal(t, uint32(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint64", func(t *testing.T) {
		var a, b, c uint64 = 10, 20, 30
		if !assert.Equal(t, uint64(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int", func(t *testing.T) {
		var a, b, c int = -10, 0, 10
		if !assert.Equal(t, int(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int8", func(t *testing.T) {
		var a, b, c int8 = -10, 0, 10
		if !assert.Equal(t, int8(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int16", func(t *testing.T) {
		var a, b, c int16 = -10, 0, 10
		if !assert.Equal(t, int16(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int32", func(t *testing.T) {
		var a, b, c int32 = -10, 0, 10
		if !assert.Equal(t, int32(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int64", func(t *testing.T) {
		var a, b, c int64 = -10, 0, 10
		if !assert.Equal(t, int64(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float32", func(t *testing.T) {
		var a, b, c float32 = -1.5, 0.0, 1.5
		if !assert.Equal(t, float32(0.0), maths.Sum(a, b, c), "Sum of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float64", func(t *testing.T) {
		var a, b, c float64 = -1.5, 0.0, 1.5
		if !assert.Equal(t, float64(0.0), maths.Sum(a, b, c), "Sum of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

}
