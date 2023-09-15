package utils

type Pipe[T any] func(T) T

func Pipeline[T any](pipes ...Pipe[T]) Pipe[T] {
	return func(t T) T {
		for _, p := range pipes {
			t = p(t)
		}
		return t
	}
}
