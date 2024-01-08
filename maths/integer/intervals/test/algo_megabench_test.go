package test

import (
	"fmt"
	"log"
	"math"
	"sort"
	"testing"
	"time"
)

// SolverResults : Solver -> Gen -> N -> time
type SolverResults map[int]map[int]map[int]time.Duration

func (r SolverResults) SetResult(solver, function, N, runs int, dur time.Duration) {
	solverResult, ok := r[solver]
	if !ok {
		solverResult = make(map[int]map[int]time.Duration)
		r[solver] = solverResult
	}

	if _, ok := solverResult[function]; !ok {
		r[solver][function] = make(map[int]time.Duration)
	}

	r[solver][function][N] = dur / time.Duration(runs)
}

type Statistics struct {
	Average time.Duration
	Median  time.Duration
	StdDev  time.Duration
}

func (r SolverResults) CalculateStatistics() {
	stats := make(map[int]map[int]Statistics)

	for solver, functions := range r {
		if _, exists := stats[solver]; !exists {
			stats[solver] = make(map[int]Statistics)
		}
		for function, results := range functions {
			var total time.Duration
			var times []float64

			for _, result := range results {
				total += result
				times = append(times, result.Seconds())
			}

			average := total / time.Duration(len(results))
			sort.Float64s(times)

			var median time.Duration
			mid := len(times) / 2
			if len(times)%2 == 0 {
				median = time.Duration((times[mid-1] + times[mid]) / 2 * float64(time.Second))
			} else {
				median = time.Duration(times[mid] * float64(time.Second))
			}

			var sumSquares float64
			for _, t := range times {
				diff := t - average.Seconds()
				sumSquares += diff * diff
			}
			stdDev := time.Duration(math.Sqrt(sumSquares/float64(len(times))) * float64(time.Second))

			stats[solver][function] = Statistics{Average: average, Median: median, StdDev: stdDev}
		}
	}

	for solver := range solvers {
		var sumAvg, sumMed, sumStd time.Duration
		for function := range gen_funcs {
			fmt.Printf("Solver%d - %s(%d): %+v\n", solver+1, gen_names[function], function, stats[solver][function])
			sumAvg += stats[solver][function].Average
			sumMed += stats[solver][function].Median
			sumStd += stats[solver][function].StdDev
		}
		fmt.Println(sumAvg/time.Duration(len(gen_funcs)), sumAvg, "|", sumMed/time.Duration(len(gen_funcs)), sumMed, "|", sumMed/time.Duration(len(gen_funcs)), sumMed)
		fmt.Println()
	}

}

func Benchmark_SuperFindCover(b *testing.B) {
	var results SolverResults = make(SolverResults)
	for n := 20000; n <= 10000000; n = 2 * n {
		for fi, F := range gen_funcs {
			for si, solver := range solvers {
				testIntervals := F(n)
				b.ResetTimer()
				b.Run(fmt.Sprintf("N=%d/Solver%d/%sGen", n, si, gen_names[fi]), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = solver(testIntervals)
					}
					b.StopTimer()
					if b.N > 10 {
						results.SetResult(si, fi, n, b.N, b.Elapsed())
					}
				})
			}
			log.Println()
		}
	}

	results.CalculateStatistics()
}
