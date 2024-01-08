package main_test

import (
	"testing"
	"math"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/relvox/iridescence_go/maths"
)

func Test_MinMax(t *testing.T) {
t.Run("uint", func(t *testing.T) {
	var a, b, c uint = 10, 20, 30
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("uint8", func(t *testing.T) {
	var a, b, c uint8 = 10, 20, 30
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("uint16", func(t *testing.T) {
	var a, b, c uint16 = 10, 20, 30
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("uint32", func(t *testing.T) {
	var a, b, c uint32 = 10, 20, 30
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("uint64", func(t *testing.T) {
	var a, b, c uint64 = 10, 20, 30
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})


t.Run("int", func(t *testing.T) {
	var a, b, c int = -10, 0, 10
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("int8", func(t *testing.T) {
	var a, b, c int8 = -10, 0, 10
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("int16", func(t *testing.T) {
	var a, b, c int16 = -10, 0, 10
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("int32", func(t *testing.T) {
	var a, b, c int32 = -10, 0, 10
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})

t.Run("int64", func(t *testing.T) {
	var a, b, c int64 = -10, 0, 10
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
		t.FailNow()
	}
})


t.Run("float32", func(t *testing.T) {
	var a, b, c float32 = -1.5, 0.0, 1.5
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %f, %f, %f", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %f, %f, %f", a, b, c) {
		t.FailNow()
	}
	// Testing with extreme values
	var max, min float32 = math.MaxFloat32, -math.MaxFloat32
	if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	
	max, min = math.SmallestNonzeroFloat32, -math.SmallestNonzeroFloat32
	if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
})

t.Run("float64", func(t *testing.T) {
	var a, b, c float64 = -1.5, 0.0, 1.5
	if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %f, %f, %f", a, b, c) {
		t.FailNow()
	}
	if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %f, %f, %f", a, b, c) {
		t.FailNow()
	}
	// Testing with extreme values
	var max, min float64 = math.MaxFloat64, -math.MaxFloat64
	if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	
	max, min = math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64
	if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
	if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
		t.FailNow()
	}
})


}
