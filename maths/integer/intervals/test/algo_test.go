package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindCover(t *testing.T) {
	tests := []struct {
		intervals [][2]int
		expected  []int
	}{
		{
			intervals: [][2]int{},
			expected:  []int{},
		}, {
			intervals: [][2]int{{2, 3}},
			expected:  []int{2, 3},
		}, {
			intervals: [][2]int{{13, 15}, {2, 4}},
			expected:  []int{2, 4, 13, 15},
		}, {
			intervals: [][2]int{{2, 3}, {4, 4}},
			expected:  []int{2, 4},
		}, {
			intervals: [][2]int{{7, 8}, {3, 4}, {5, 6}, {10, 12}, {1, 2}},
			expected:  []int{1, 8, 10, 12},
		}, {
			intervals: [][2]int{{3, 5}, {2, 4}},
			expected:  []int{2, 5},
		}, {
			intervals: [][2]int{{10, 12}, {4, 6}, {0, 8}, {7, 9}, {1, 3}},
			expected:  []int{0, 12},
		}, {
			intervals: [][2]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			expected:  []int{1, 10},
		}, {
			intervals: [][2]int{{128, 191}, {192, 223}},
			expected:  []int{128, 223},
		}, {
			intervals: [][2]int{{32, 127}, {128, 159}},
			expected:  []int{32, 159},
		}, {
			intervals: [][2]int{{64, 159}, {160, 191}},
			expected:  []int{64, 191},
		}}
	for i, tt := range tests {
		for si, solver := range solvers {
			t.Run(fmt.Sprintf("Solver%d/Test%d", si, i), func(t *testing.T) {
				var testIntervals [][2]int = make([][2]int, len(tt.intervals))
				copy(testIntervals, tt.intervals)
				actual := solver(testIntervals)

				if len(actual) != len(tt.expected) {
					t.FailNow()
				}
				for i, v := range tt.expected {
					if v != actual[i] {
						t.FailNow()
					}
				}
			})
		}
	}
}
func Test_FindCover_Group(t *testing.T) {
	for n := 1; n <= 1_000_000; n *= 10 {
		for fi, F := range gen_funcs {
			testIntervals := F(n)
			for i := 0; i < len(solvers)-1; i++ {
				baseSolver := solvers[i]
				for j := i + 1; j < len(solvers); j++ {
					otherSolver := solvers[j]
					t.Run(fmt.Sprintf("n=%d/%sGen/Solver_%d_%d", n, gen_names[fi], i, j), func(t *testing.T) {
						baseResult := baseSolver(testIntervals)
						otherResult := otherSolver(testIntervals)
						if baseResult == nil || otherResult == nil {
							return
						}
						if !assert.ElementsMatch(t, baseResult, otherResult) {
							t.FailNow()
						}
					})
				}
			}
		}
	}
}
