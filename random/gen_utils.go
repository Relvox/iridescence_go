package random

import "github.com/relvox/iridescence_go/utils"

type Gen[T any] func(p IntnGenerator) T

func OneOf[T any](gens ...Gen[T]) Gen[T] {
	return func(p IntnGenerator) T { return gens[p.Intn(len(gens))](p) }
}

func Aggregate[G Gen[T], T any](agg func(T, T) T, gens ...Gen[T]) Gen[T] {
	if len(gens) == 0 || agg == nil {
		panic("not enough gens or nil predicate")
	}
	return func(p IntnGenerator) T {
		result := gens[0](p)
		for i := 1; i < len(gens); i++ {
			result = agg(result, gens[i](p))
		}
		return result
	}
}

func Generate[R Gen[T], T any](gen R, pipe utils.Pipe[T]) Gen[T] {
	return func(p IntnGenerator) T {
		obj := gen(p)
		obj = pipe(obj)
		return obj
	}
}
