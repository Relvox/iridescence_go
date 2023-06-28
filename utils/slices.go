package utils

// Contains checks whether a slice contains an item
func Contains[T comparable](items []T, test T) bool {
	for _, t := range items {
		if t == test {
			return true
		}
	}
	return false
}

// Equal checks if two slices have the same elements in the same order
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

// Distinct narrows down a slice to only distinct items.
func Distinct[T comparable](items []T) []T {
	if len(items) == 0 {
		return items
	}
	result := []T{}
	var test Set[T] = make(Set[T])
	for i := 0; i < len(items); i++ {
		item := items[i]
		if _, ok := test[item]; ok {
			continue
		}
		result = append(result, item)
		test[item] = U
	}
	return result
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
