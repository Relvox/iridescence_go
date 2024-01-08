package utils_test

import (
	"testing"
)

func foo(n int) int {
	n2 := uint64(n)
	n2 <<= 10
	return int(n2 / 1024)
}

const SIZE = 100

func initMatrix() [][]int {
	matrix := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		matrix[i] = make([]int, SIZE)
		for j := 0; j < SIZE; j++ {
			matrix[i][j] = i + j*SIZE
		}
	}
	return matrix
}

func Benchmark_EXP_Foo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var n int = foo(i)
		_ = n
	}
}

func Benchmark_EXP_Loop(b *testing.B) {
	m := initMatrix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < SIZE; x++ {
			for y := 0; y < SIZE; y++ {
				m[x][y] = foo(m[x][y])
			}
		}
	}
}

func FF(mat [][]int, x, y int) {
	mat[x][y] = foo(mat[x][y])
}

func Benchmark_EXP_Loop_Func(b *testing.B) {
	m := initMatrix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < SIZE; x++ {
			for y := 0; y < SIZE; y++ {
				FF(m, x, y)
			}
		}
	}
}

func FLOOP(mat [][]int, x, y int) {
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			FF(mat, x, y)
		}
	}
}

func Benchmark_EXP_FLoop_Func(b *testing.B) {
	m := initMatrix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < SIZE; x++ {
			for y := 0; y < SIZE; y++ {
				FF(m, x, y)
			}
		}
	}
}

func FFLOOP(mat [][]int, fop func(mat [][]int, x, y int)) {
	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			fop(mat, x, y)
		}
	}
}

func Benchmark_EXP_FFLoop_Func(b *testing.B) {
	m := initMatrix()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		FFLOOP(m, FF)
	}
}
