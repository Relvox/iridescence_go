package sources_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/relvox/iridescence_go/random/sources"
	"github.com/stretchr/testify/assert"
)

const (
	BUCKETS = 100000
	SIZE    = 100_000_000
)

func getBucket(min, max, buckets, current int64) int {
	if max < min || current < min || current > max || buckets <= 0 {
		panic(fmt.Errorf("not %d(min) <= %d(current) <= %d(max) | %d(buckets)", min, current, max, buckets))
	}
	if current == min {
		return 0
	}
	if current == max {
		return int(buckets) - 1
	}

	bucketSize := (max/buckets - min/buckets + 1/buckets)
	res := current/bucketSize - min/bucketSize
	return int(res)
}

func Test_GetBucket(t *testing.T) {
	// Normal range
	t.Run("Normal range", func(t *testing.T) {
		min, max, buckets := int64(0), int64(100), int64(10)
		current := int64(50)
		expectedBucket := 5
		assert.Equal(t, expectedBucket, getBucket(min, max, buckets, current), "Bucket for mid-range value")
	})

	// Edge cases
	t.Run("Edge cases", func(t *testing.T) {
		min, max, buckets := int64(0), int64(100), int64(10)

		// Test for minimum edge
		assert.Equal(t, 0, getBucket(min, max, buckets, min), "Bucket for minimum value")

		// Test for maximum edge
		assert.Equal(t, 9, getBucket(min, max, buckets, max-1), "Bucket for maximum value (excluding max)")
	})

	// Invalid inputs
	t.Run("Invalid inputs", func(t *testing.T) {
		// Testing with min >= max
		assert.Panics(t, func() { getBucket(100, 100, 10, 50) }, "Should panic when min >= max")

		// Testing with buckets <= 0
		assert.Panics(t, func() { getBucket(0, 100, 0, 50) }, "Should panic when buckets <= 0")

		// Testing with current outside min to max range
		assert.Panics(t, func() { getBucket(0, 100, 10, 150) }, "Should panic when current is outside the min to max range")
	})

	// Special values
	t.Run("Special values", func(t *testing.T) {
		// Test for special value
		assert.Equal(t, 5, getBucket(math.MinInt64, math.MaxInt64, 10, 0), "Bucket for zero with full int64 range")
	})
}

func Test_Rands(t *testing.T) {
	t.Run("go_src_Int63", func(t *testing.T) {
		src := rand.NewSource(-1)
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxInt64, BUCKETS, src.Int63())]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)
	})
	t.Run("go_rnd_Uint32", func(t *testing.T) {
		rng := rand.New(rand.NewSource(-1))
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxUint32, BUCKETS, int64(rng.Uint32()))]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)
	})
	t.Run("go_wrap_Int63", func(t *testing.T) {
		src := sources.NewGoSource(rand.NewSource(-1))
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxInt64, BUCKETS, src.Int63())]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)
	})
	t.Run("go_wrap_Uint32", func(t *testing.T) {
		rng := rand.New(sources.NewGoSource(rand.NewSource(-1)))
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxUint32, BUCKETS, int64(rng.Uint32()))]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)

	})
	t.Run("my_src_Int63", func(t *testing.T) {
		src := sources.NewWELL512(0)
		src.Seed(-1)
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxInt64, BUCKETS, src.Int63())]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}

		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)

	})
	t.Run("my_src_Uint32", func(t *testing.T) {
		rng := sources.NewWELL512(0)
		rng.Seed(-1)
		buckets := make(map[int]int, BUCKETS)
		for i := 0; i < SIZE; i++ {
			buckets[getBucket(0, math.MaxUint32, BUCKETS, int64(rng.Uint32()))]++
		}
		min, max := buckets[0], buckets[0]
		for _, v := range buckets {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		// chi := sadistics.ChiSquared(buckets, SIZE)
		// log.Println(min, max, chi)

	})
}

func Benchmark(b *testing.B) {

	b.Run("go_src_Int63", func(b *testing.B) {
		src := rand.NewSource(-1)
		for i := 0; i < b.N; i++ {
			_ = src.Int63()
		}
	})

	b.Run("go_rnd_Uint32", func(b *testing.B) {
		rng := rand.New(rand.NewSource(-1))
		for i := 0; i < b.N; i++ {
			_ = rng.Uint32()
		}
	})

	b.Run("go_wrap_Int63", func(b *testing.B) {
		src := sources.NewGoSource(rand.NewSource(-1))
		for i := 0; i < b.N; i++ {
			_ = src.Int63()
		}
	})

	b.Run("go_wrap_Uint32", func(b *testing.B) {
		rng := rand.New(sources.NewGoSource(rand.NewSource(-1)))
		for i := 0; i < b.N; i++ {
			_ = rng.Uint32()
		}
	})

	b.Run("my_src_Int63", func(b *testing.B) {
		src := sources.NewWELL512(0)
		src.Seed(-1)
		for i := 0; i < b.N; i++ {
			_ = src.Int63()
		}
	})

	b.Run("my_src_Uint32", func(b *testing.B) {
		rng := sources.NewWELL512(0)
		rng.Seed(-1)
		for i := 0; i < b.N; i++ {
			_ = rng.Uint32()
		}
	})
}
