package option_test

import (
	"fmt"
	"testing"

	. "github.com/relvox/iridescence_go/experimental/option"
)

func Benchmark_Option(b *testing.B) {
	run_Benchmark_Option[int](b, 0, 1)
	run_Benchmark_Option[uint](b, 0, 1)
	run_Benchmark_Option[float64](b, 0, 1)

	run_Benchmark_Option[dummy](b, dummy{}, dummy{"A", 1})
	run_Benchmark_Option[*dummy](b, nil, &dummy{"A", 1})
	run_Benchmark_Option[any](b, dummy{}, dummy{"A", 1})
	run_Benchmark_Option[any](b, nil, dummy{"A", 1})
	run_Benchmark_Option[any](b, nil, &dummy{"A", 1})
}

func run_Benchmark_Option[T any](b *testing.B, noneVal, someVal T) {
	b.Run(fmt.Sprintf("%T(%#v,%#v)", someVal, noneVal, someVal), func(b *testing.B) {
		b.Run("None", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var opt Option[T] = None[T]()
				_ = opt.HasValue()
				opt.Try(func(val T) {})
				opt.Do(func() {}, func(val T) {})
			}
		})

		b.Run("Some", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var opt Option[T] = Some(someVal)
				_ = opt.HasValue()
				opt.Try(func(val T) {})
				opt.Do(func() {}, func(val T) {})
			}
		})

		b.Run("Maybe", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var opt Option[T] = Maybe(noneVal)
				_ = opt.HasValue()
				opt.Try(func(val T) {})
				opt.Do(func() {}, func(val T) {})

				opt = Maybe(someVal)
				_ = opt.HasValue()
				opt.Try(func(val T) {})
				opt.Do(func() {}, func(val T) {})

			}
		})
	})
}
