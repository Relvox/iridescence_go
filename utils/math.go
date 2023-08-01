package utils

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

func Abs[N constraints.Integer](i N) N {
	if i >= 0 {
		return i
	}
	return -i

}

func Max[N constraints.Integer](vs ...N) N {
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

func Min[N constraints.Integer](vs ...N) N {
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

func Sum[N constraints.Integer](vs ...N) N {
	var sum N = 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func GeometricMean[N constraints.Integer](vals ...N) N {
	var prod float64 = 1.0
	for _, v := range vals {
		prod *= float64(v)
	}
	return N(math.Pow(prod, 1.0/float64(len(vals))))
}

func XenoSum[N constraints.Integer](vals ...N) N {
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
