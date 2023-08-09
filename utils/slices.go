package utils

import "golang.org/x/exp/slices"

func init() {
	slices.Contains([]int{1, 2, 3}, 1)
	slices.Equal([]int{1, 2}, []int{1, 2})
	
}

// Same checks if two slices have the same elements in any order
func Same[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	has := map[T]int{}
	for i := 0; i < len(a); i++ {
		has[a[i]]++
		has[b[i]]--
	}
	for _, v := range has {
		if v != 0 {
			return false
		}
	}
	return true
}

func Repeat[T any](item T, count int) []T {
	var result []T = make([]T, count)
	for i := range result {
		result[i] = item
	}
	return result
}

func Intersect[T comparable](list1, list2 []T) []T {
	var set Set[T] = make(Set[T])
	for _, v := range list1 {
		set[v] = U
	}
	var result []T
	for _, v := range list2 {
		if _, ok := set[v]; !ok {
			continue
		}
		result = append(result, v)
	}
	return result
}
