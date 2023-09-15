package sets_test

import (
	"strconv"
	"testing"

	"github.com/relvox/iridescence_go/sets"
)

func generateSet(size, start int) sets.Set[int] {
	set := sets.NewSet[int]()
	for i := 0; i < size; i++ {
		set[start+i] = sets.U
	}
	return set
}

func generateElements(size, start int) []int {
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = i + start
	}
	return elements
}

func BenchmarkSetUnion(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			set2 := generateSet(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.SetUnion(set2)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			elements := generateElements(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.Union(elements...)
			}
		})
	}
}

func BenchmarkSetIntersection(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			set2 := generateSet(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.SetIntersection(set2)
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			elements := generateElements(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.Intersection(elements...)
			}
		})
	}
}

func BenchmarkSetDifference(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			set2 := generateSet(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.SetDifference(set2)
			}
		})
	}
}

func BenchmarkDifference(b *testing.B) {
	for size := 100; size < 10_000_00; size *= 10 {
		b.Run("Size"+strconv.Itoa(size), func(b *testing.B) {
			set1 := generateSet(size, 0)
			elements := generateElements(size, size/2)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set1.Difference(elements...)
			}
		})
	}
}
