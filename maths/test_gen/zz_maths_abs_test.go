package main_test

import (
	"testing"
	"math"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/relvox/iridescence_go/maths"
)

func Test_Abs(t *testing.T) {
t.Run("uint", func(t *testing.T) {
	var i uint = 0
	for ; i < 100; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxUint - 100; i != 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("uint8", func(t *testing.T) {
	var i uint8 = 0
	for ; i < 100; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxUint8 - 100; i != 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("uint16", func(t *testing.T) {
	var i uint16 = 0
	for ; i < 100; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxUint16 - 100; i != 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("uint32", func(t *testing.T) {
	var i uint32 = 0
	for ; i < 100; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxUint32 - 100; i != 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("uint64", func(t *testing.T) {
	var i uint64 = 0
	for ; i < 100; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxUint64 - 100; i != 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})


t.Run("int", func(t *testing.T) {
	var i int = math.MinInt
	for ; i < math.MinInt+100; i++ {
		if !assert.Equal(t, int(-i), maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = -50; i < 50; i++ {
		expected := int(i)
		if i < 0 {
			expected = -expected
		}
		if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxInt - 100; i > 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("int8", func(t *testing.T) {
	var i int8 = math.MinInt8
	for ; i < math.MinInt8+100; i++ {
		if !assert.Equal(t, int8(-i), maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = -50; i < 50; i++ {
		expected := int8(i)
		if i < 0 {
			expected = -expected
		}
		if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxInt8 - 100; i > 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("int16", func(t *testing.T) {
	var i int16 = math.MinInt16
	for ; i < math.MinInt16+100; i++ {
		if !assert.Equal(t, int16(-i), maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = -50; i < 50; i++ {
		expected := int16(i)
		if i < 0 {
			expected = -expected
		}
		if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxInt16 - 100; i > 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("int32", func(t *testing.T) {
	var i int32 = math.MinInt32
	for ; i < math.MinInt32+100; i++ {
		if !assert.Equal(t, int32(-i), maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = -50; i < 50; i++ {
		expected := int32(i)
		if i < 0 {
			expected = -expected
		}
		if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxInt32 - 100; i > 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})

t.Run("int64", func(t *testing.T) {
	var i int64 = math.MinInt64
	for ; i < math.MinInt64+100; i++ {
		if !assert.Equal(t, int64(-i), maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = -50; i < 50; i++ {
		expected := int64(i)
		if i < 0 {
			expected = -expected
		}
		if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
	for i = math.MaxInt64 - 100; i > 0; i++ {
		if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
			t.FailNow()
		}
	}
})


t.Run("float32", func(t *testing.T) {
	for sign := float32(-1); sign <= 1; sign += 2 {
		for f, j := float32(0), 0; j < 100; j++ {
			expected := f
			if f < 0 {
				expected = -f
			}
			if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
				t.FailNow()
			}
			f = f*2 + math.SmallestNonzeroFloat32*sign
		}

		for f, j := (math.MaxFloat32-100)*sign, 0; j < 100; j++ {
			expected := f
			if f < 0 {
				expected = -f
			}
			if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
				t.FailNow()
			}
			f += 1.0
		}
	}
	for f, j := float32(-50), 0; j < 100; j++ {
		expected := f
		if f < 0 {
			expected = -f
		}
		if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
			t.FailNow()
		}
		f += 1.0
	}
})

t.Run("float64", func(t *testing.T) {
	for sign := float64(-1); sign <= 1; sign += 2 {
		for f, j := float64(0), 0; j < 100; j++ {
			expected := f
			if f < 0 {
				expected = -f
			}
			if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
				t.FailNow()
			}
			f = f*2 + math.SmallestNonzeroFloat64*sign
		}

		for f, j := (math.MaxFloat64-100)*sign, 0; j < 100; j++ {
			expected := f
			if f < 0 {
				expected = -f
			}
			if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
				t.FailNow()
			}
			f += 1.0
		}
	}
	for f, j := float64(-50), 0; j < 100; j++ {
		expected := f
		if f < 0 {
			expected = -f
		}
		if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
			t.FailNow()
		}
		f += 1.0
	}
})


}
