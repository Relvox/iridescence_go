package utils

type Cloner[T any] interface {
	Clone() T
}

func CloneSlice[T Cloner[T]](srcs []T) []T {
	result := make([]T, len(srcs))
	for i, src := range srcs {
		result[i] = src.Clone()
	}
	return result
}

func CloneSliceWithCloner[T any](srcs []T, clone func(T) T) []T {
	result := make([]T, len(srcs))
	for i, src := range srcs {
		result[i] = clone(src)
	}
	return result
}

func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range src {
		result[k] = v
	}
	return result
}
