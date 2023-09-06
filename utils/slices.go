package utils

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

func Crop[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	if len(slice) == 1 {
		return []T{}
	}
	if index == 0 {
		return slice[1:]
	}
	if index == len(slice)-1 {
		return slice[:index]
	}
	left := slice[:index]
	right := slice[index+1:]
	return append(left, right...)
}

func FindPred[E ~[]T, T any](list E, pred func(T) bool) (T, bool) {
	for i := range list {
		if pred(list[i]) {
			return list[i], true
		}
	}
	var t T
	return t, false
}

func Last[E ~[]T, T any](list E) (T, bool) {
	if len(list) == 0 {
		var t T
		return t, false
	}
	return list[len(list)-1], true
}
