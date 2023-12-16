package utils

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sign[N Number](n N) N {
	if n == 0 {
		return n
	}
	return n / Abs(n)
}

func Abs[N Number](i N) N {
	if i >= 0 {
		return i
	}
	return -i
}

func Max[N Number](vs ...N) N {
	if len(vs) == 0 {
		return 0
	}
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > max {
			max = vs[i]
		}
	}
	return max
}

func Min[N Number](vs ...N) N {
	if len(vs) == 0 {
		return 0
	}
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] < min {
			min = vs[i]
		}
	}
	return min
}

func Sum[N constraints.Ordered](vs ...N) N {
	var sum N
	for _, v := range vs {
		sum += v
	}
	return sum
}

func GeometricMean[N constraints.Integer | constraints.Float](vals ...N) N {
	var prod float64 = 1.0
	for _, v := range vals {
		prod *= float64(v)
	}
	return N(math.Pow(prod, 1.0/float64(len(vals))))
}

func XenoSum[N constraints.Integer | constraints.Float](vals ...N) N {
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})
	var sum, factor N
	sum, factor = 0, 1
	for i := len(vals) - 1; i >= 0; i-- {
		sum += vals[i] / factor
		factor *= 2
	}
	return sum
}
