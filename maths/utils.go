package maths

import (
	"github.com/toolvox/utilgo/pkg/mathutil"
)

type Number = mathutil.Number

func Sign[N Number](n N) N {
	switch {
	case n > 0:
		return n / n
	case n < 0:
		return 0 - (n / n)
	default:
		return 0
	}
}

func Abs[N Number](n N) N {
	if n >= 0 {
		return n
	}
	return -n
}

func Bounds[N Number](vs ...N) (N, N) {
	if len(vs) == 0 {
		return 0, 0
	}
	min := vs[0]
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > max {
			max = vs[i]
		}
		if vs[i] < min {
			min = vs[i]
		}
	}
	return min, max
}

func Clamp[N Number](min, val, max N) N {
	switch {
	case val < min:
		val = min
	case val > max:
		val = max
	}
	return val
}
