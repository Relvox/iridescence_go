package integer

import (
	"golang.org/x/exp/constraints"
)

func Log2[N constraints.Integer](n N) uint {
	if n <= 1 {
		return 0
	}
	top := uint(1)
	i := uint(0)
	for top < uint(n) {
		top <<= 1
		i++
	}
	return i
}
