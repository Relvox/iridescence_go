package utils

import "golang.org/x/exp/constraints"

func BitIndices[T constraints.Integer](num T) []int {
	var results []int
	for i := 0; num != 0; i++ {
		digit := num % 2
		num = num / 2
		if digit == 0 {
			continue
		}
		results = append(results, i)
	}
	return results
}
