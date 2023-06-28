package utils

import (
	"math"
	"sort"
)

type IntLike interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Abs[N IntLike](i N) N {
	if i >= 0 {
		return i
	}
	return -i
}

func Max[N IntLike](vs ...N) N {
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

func Min[N IntLike](vs ...N) N {
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

func Sum[N IntLike](vs ...N) N {
	var sum N = 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func GeometricMean[N IntLike](vals ...N) N {
	var prod float64 = 1.0
	for _, v := range vals {
		prod *= float64(v)
	}
	return N(math.Pow(prod, 1.0/float64(len(vals))))
}

func XenoSum[N IntLike](vals ...N) N {
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
