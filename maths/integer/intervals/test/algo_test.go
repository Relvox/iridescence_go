package test

import (
	"fmt"
	"testing"

	"github.com/relvox/iridescence_go/asserts"
)

func Test_FindCover(t *testing.T) {
	tests := []struct {
		intervals [][2]int
		expected  []int
	}{
		{intervals: [][2]int{},
			expected: []int{}},
		{intervals: [][2]int{{2, 3}},
			expected: []int{2, 3}},
		{intervals: [][2]int{{13, 15}, {2, 4}},
			expected: []int{2, 4, 13, 15}},
		{intervals: [][2]int{{2, 3}, {4, 4}},
			expected: []int{2, 4}},
		{intervals: [][2]int{{7, 8}, {3, 4}, {5, 6}, {10, 12}, {1, 2}},
			expected: []int{1, 8, 10, 12}},
		{intervals: [][2]int{{3, 5}, {2, 4}},
			expected: []int{2, 5}},
		{intervals: [][2]int{{10, 12}, {4, 6}, {0, 8}, {7, 9}, {1, 3}},
			expected: []int{0, 12}},
		{intervals: [][2]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			expected: []int{1, 10}},
		{intervals: [][2]int{{128, 191}, {192, 223}},
			expected: []int{128, 223}},
		{intervals: [][2]int{{32, 127}, {128, 159}},
			expected: []int{32, 159}},
		{intervals: [][2]int{{64, 159}, {160, 191}},
			expected: []int{64, 191}},
		// {intervals: [][2]int{
		// 	1 {Structural Support [[160 191] [192 223]]} [[32 95] [96 127]]} [[64 127] [128 159]]}}
		// 	1 {Structural Support [160 223] [32 127] [64 159]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	2 {Specialized Construction Element [192 223] [[32 63] [64 95]]} [[128 159] [160 191]]}}
		// 	2 {Specialized Construction Element [192 223] [32 95] [128 191]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	3 {Simple Ration [[64 159] [160 191]]} [[64 159] [160 191]]} [[128 191] [192 223]]}}
		// 	3 {Simple Ration [64 191] [64 191] [128 223]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	4 {Nutritional Supplement [[64 127] [128 159]]} [[64 127] [128 159]]} [[160 191] [192 223]]}}
		// 	4 {Nutritional Supplement [64 159] [64 159] [160 223]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	5 {Emergency Medicine [[64 95] [128 159]]} [[96 127] [128 159]]} [192 223]}
		// 	5 {Emergency Medicine [64 95 128 159] [96 159] [192 223]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	6 {Fodder [[64 159] [160 191]]} [[96 191] [192 223]]} [[64 159] [160 191]]}}
		// 	6 {Fire Fodder [64 191] [96 223] [64 191]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	7 {Fuel [[64 127] [128 159]]} [[160 191] [192 223]]} [[32 95] [160 191]]}}
		// 	7 {Fuel Cell [128 223] [160 223] [32 95 160 191]}
		// 	expected:
		// 	[]int{1, 10}},
		// {intervals: [][2]int{
		// 	8 {Power Core [[64 95] [192 223]]} [192 223] [[32 63] [192 223]]}}
		// 	8 {Power Core [64 95 192 223] [192 223] [32 63 192 223]}
		// 	expected:
		// 	[]int{1, 10}},
	}
	for i, tt := range tests {
		for si, solver := range solvers {
			t.Run(fmt.Sprintf("Solver%d/Test%d", si, i), func(t *testing.T) {
				var testIntervals [][2]int = make([][2]int, len(tt.intervals))
				copy(testIntervals, tt.intervals)
				actual := solver(testIntervals)
				fmt.Println(tt.expected, " vs. ", actual)
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
	for n := 1; n <= 100_000; n *= 10 {
		for fi, F := range gen_funcs {
			testIntervals := F(n)
			for i := 0; i < len(solvers)-1; i++ {
				baseSolver := solvers[i]
				for j := i + 1; j < len(solvers); j++ {
					otherSolver := solvers[j]
					t.Run(fmt.Sprintf("n=%d/%sGen/Solver_%d_%d", n, gen_names[fi], i, j), func(t *testing.T) {
						baseResult := baseSolver(testIntervals)
						otherResult := otherSolver(testIntervals)
						if !asserts.SameElements(t, baseResult, otherResult) {
							t.FailNow()
						}
					})
				}
			}
		}
	}
}
