package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/relvox/iridescence_go/maths"
)

func Test_GeometricMean(t *testing.T) {
	t.Run("uint", func(t *testing.T) {
		var a, b, c uint = 1, 2, 4
		if !assert.Equal(t, uint(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint8", func(t *testing.T) {
		var a, b, c uint8 = 1, 2, 4
		if !assert.Equal(t, uint8(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint16", func(t *testing.T) {
		var a, b, c uint16 = 1, 2, 4
		if !assert.Equal(t, uint16(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var a, b, c uint32 = 1, 2, 4
		if !assert.Equal(t, uint32(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("uint64", func(t *testing.T) {
		var a, b, c uint64 = 1, 2, 4
		if !assert.Equal(t, uint64(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int", func(t *testing.T) {
		var a, b, c int = 1, 2, 4
		if !assert.Equal(t, int(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int8", func(t *testing.T) {
		var a, b, c int8 = 1, 2, 4
		if !assert.Equal(t, int8(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int16", func(t *testing.T) {
		var a, b, c int16 = 1, 2, 4
		if !assert.Equal(t, int16(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int32", func(t *testing.T) {
		var a, b, c int32 = 1, 2, 4
		if !assert.Equal(t, int32(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("int64", func(t *testing.T) {
		var a, b, c int64 = 1, 2, 4
		if !assert.Equal(t, int64(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float32", func(t *testing.T) {
		var a, b, c float32 = 1.0, 2.0, 4.0
		if !assert.Equal(t, float32(2.0), maths.GeometricMean(a, b, c), "Geometric mean of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

	t.Run("float64", func(t *testing.T) {
		var a, b, c float64 = 1.0, 2.0, 4.0
		if !assert.Equal(t, float64(2.0), maths.GeometricMean(a, b, c), "Geometric mean of %f, %f, %f", a, b, c) {
			t.FailNow()
		}
	})

}
