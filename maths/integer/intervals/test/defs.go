package test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/relvox/iridescence_go/maths/integer/intervals"
)

var N = 10_000_000

var (
	gen_funcs = []func(n int) [][2]int{
		generateChallengingIntervals,
		generateRandomBlobIntervals,
		generateAlternatingSizeIntervals,
		generateDenseOverlappingIntervals,
		generateSparseIntervals,
	}
	gen_names = []string{
		"Challenging",
		"RandomBlob",
		"AlternatingSize",
		"DenseOverlapping",
		"Sparse",
	}
	solvers = []func(ints [][2]int) []int{
		intervals.FindCover[int],
		intervals.FindCover1[int],
		intervals.FindCover2[int],
		intervals.FindCover3[int],
		intervals.FindCover4[int],
	}
)

func generateChallengingIntervals(n int) [][2]int {
	rng := rand.New(rand.NewSource(int64(n)))
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		start := rng.Intn(n * n)
		end := start + rng.Intn(n) // ensuring some overlap
		intervals[i] = [2]int{start, end}
	}
	return intervals
}

func generateSparseIntervals(n int) [][2]int {
	rng := rand.New(rand.NewSource(int64(n)))
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		start := rng.Intn(n * 10)
		end := start + rng.Intn(n)
		intervals[i] = [2]int{start, end}
	}
	return intervals
}

func generateDenseOverlappingIntervals(n int) [][2]int {
	rng := rand.New(rand.NewSource(int64(n)))
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		start := rng.Intn(n)
		end := start + rng.Intn(n/2+1) // higher chance of overlap
		intervals[i] = [2]int{start, end}
	}
	return intervals
}

func generateAlternatingSizeIntervals(n int) [][2]int {
	rng := rand.New(rand.NewSource(int64(n)))
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		start := rng.Intn(n * 3)
		end := start
		if i%2 == 0 {
			end += rng.Intn(n/10 + 1) // small intervals
		} else {
			end += rng.Intn(n) // large intervals
		}
		intervals[i] = [2]int{start, end}
	}
	return intervals
}

func generateRandomBlobIntervals(n int) [][2]int {
	rng := rand.New(rand.NewSource(int64(n)))
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		start := rng.Intn(n * n)
		end := start + rng.Intn(n*n)
		intervals[i] = [2]int{start, end}
	}
	return intervals
}

func Test_Generators(t *testing.T) {
	var data [][2]int
	data = generateChallengingIntervals(20)
	log.Println("Challenging", data, len(data), "\n->\n", intervals.FindCover(data), len(intervals.FindCover(data))/2)
	data = generateSparseIntervals(20)
	log.Println("Sparse", data, len(data), "\n->\n", intervals.FindCover(data), len(intervals.FindCover(data))/2)
	data = generateDenseOverlappingIntervals(20)
	log.Println("DenseOverlapping", data, len(data), "\n->\n", intervals.FindCover(data), len(intervals.FindCover(data))/2)
	data = generateAlternatingSizeIntervals(20)
	log.Println("AlternatingSize", data, len(data), "\n->\n", intervals.FindCover(data), len(intervals.FindCover(data))/2)
	data = generateRandomBlobIntervals(20)
	log.Println("RandomSizeAndPosition", data, len(data), "\n->\n", intervals.FindCover(data), len(intervals.FindCover(data))/2)
}
