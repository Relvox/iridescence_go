package utils_test

import (
	"fmt"
	"testing"

	assutil "github.com/relvox/iridescence_go/assert"
	"github.com/relvox/iridescence_go/utils"
	"github.com/stretchr/testify/assert"
)

func gen_map[K comparable, V any](size int, genKey func() K, genVal func() V) map[K]V {
	result := make(map[K]V)
	for i := 0; i < size; i++ {
		result[genKey()] = genVal()
	}
	return result
}

func split_map[K comparable, V any](toSplit map[K]V) (map[K]V, map[K]V) {
	resLeft, resRight := make(map[K]V), make(map[K]V)
	left := true
	for k, v := range toSplit {
		if left {
			resLeft[k] = v
		} else {
			resRight[k] = v
		}
		left = !left
	}
	return resLeft, resRight
}

func Test_MergeMaps(t *testing.T) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }

	for k := 10; k <= 1000000; k *= 10 {
		t.Run(fmt.Sprintf("%d keys", k), func(t *testing.T) {
			expected := gen_map(k, intGen, intGen)
			m1, m2 := split_map(expected)
			actual := utils.MergeMaps(m1, m2)
			assert.InDeltaMapValues(t, expected, actual, 0, "maps should be identical")
		})
	}
}

func Benchmark_MergeMaps(b *testing.B) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }
	for k := 10; k <= 1000000; k *= 10 {
		b.Run(fmt.Sprintf("%d keys", k), func(b *testing.B) {
			expected := gen_map(k, intGen, intGen)
			m1, m2 := split_map(expected)
			for i := 0; i < b.N; i++ {
				utils.MergeMaps(m1, m2)
			}
		})
	}
}

func Test_Values(t *testing.T) {

	for k := 10; k <= 1000000; k *= 10 {
		t.Run(fmt.Sprintf("%d keys", k), func(t *testing.T) {
			__Z := 0
			intGen := func() int { __Z++; return __Z }
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = (i + 1) * 2
			}
			actual := utils.Values(originalMap)
			assutil.SameElements(t, expected, actual)
		})

	}
}

func Benchmark_Values(b *testing.B) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }
	for k := 10; k <= 1000000; k *= 10 {
		b.Run(fmt.Sprintf("%d keys", k), func(b *testing.B) {
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = (i + 1) * 2
			}
			for i := 0; i < b.N; i++ {
				utils.Values(originalMap)
			}
		})
	}
}
