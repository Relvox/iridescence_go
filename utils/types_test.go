package utils_test

import (
	"testing"

	"github.com/relvox/iridescence_go/utils"
)

func TestIsNilOrZero(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want bool
	}{
		{"NilPointer", (*int)(nil), true},
		{"NonNilPointer", new(int), false},
		{"NilSlice", []int(nil), true},
		{"NonNilSlice", []int{}, false},
		{"NonNilSliceWithData", []int{1, 2, 3}, false},
		{"NilMap", map[string]int(nil), true},
		{"NonNilMap", map[string]int{}, false},
		{"NonNilMapWithData", map[string]int{"a": 1}, false},
		{"NilChan", (chan int)(nil), true},
		{"NonNilChan", make(chan int), false},
		{"NilFunc", (func())(nil), true},
		{"NonNilFunc", func() {}, false},
		{"IntZero", 0, true},
		{"IntNonZero", 1, false},
		{"StringEmpty", "", true},
		{"StringNonEmpty", "a", false},
		{"NilInterface", (any)(nil), true},
		{"NonNilInterface", any(0), true},
		{"NonNilInterface", any(1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsNilOrZero(tt.val); got != tt.want {
				t.Errorf("IsNilOrZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsNilOrZero(b *testing.B) {

	b.Run("(*int)(nil)", func(b *testing.B) {
		z := (*int)(nil)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("new(int)", func(b *testing.B) {
		z := new(int)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("[]int(nil)", func(b *testing.B) {
		z := []int(nil)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("[]int{}", func(b *testing.B) {
		z := []int{}
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("map[string]int(nil)", func(b *testing.B) {
		z := map[string]int(nil)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("map[string]int{}", func(b *testing.B) {
		z := map[string]int{}
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("make(chan int)", func(b *testing.B) {
		z := make(chan int)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("func() {}", func(b *testing.B) {
		z := func() {}
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("0", func(b *testing.B) {
		z := 0
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("1", func(b *testing.B) {
		z := 1
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("\"\"", func(b *testing.B) {
		z := ""
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("\"a\"", func(b *testing.B) {
		z := "a"
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("(any)(nil)", func(b *testing.B) {
		z := (any)(nil)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
	b.Run("any(0)", func(b *testing.B) {
		z := any(0)
		for i := 0; i < b.N; i++ {
			_ = utils.IsNilOrZero(z)
		}
	})
}
