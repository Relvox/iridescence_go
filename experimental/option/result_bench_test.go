package option_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/relvox/iridescence_go/experimental/option"
)

func Benchmark_Result1(b *testing.B) {
	run_Benchmark_Result1[int](b, 7, errors.New("error"))
	run_Benchmark_Result1[uint](b, 7, errors.New("error"))
	run_Benchmark_Result1[float64](b, 7, errors.New("error"))
	run_Benchmark_Result1[dummy](b, dummy{"a", 1}, errors.New("error"))
}

func run_Benchmark_Result1[TOk any](b *testing.B, okVal TOk, errVal error) {
	b.Run(fmt.Sprintf("%T/%T", okVal, errVal), func(b *testing.B) {
		b.Run("Ok", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := Ok[TOk](okVal)
				_ = res.IsOk()
				res.Try(func(val TOk) {})
				res.Handle(func(val TOk) {}, func(err error) {})
			}
		})

		b.Run("Err", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := Err[TOk](errVal)
				_ = res.IsErr()
				res.Try(func(val TOk) {})
				res.Handle(func(val TOk) {}, func(err error) {})
			}
		})
	})
}
