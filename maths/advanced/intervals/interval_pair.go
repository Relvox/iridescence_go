package intervals

import (
	"golang.org/x/exp/constraints"
)

func NewInterval[T constraints.Ordered](min T, includeMin bool, max T, includeMax bool) Interval[T] {
	if max == min && includeMin && includeMax {
		return NewSingleton(min)
	}
	if max <= min {
		return NewNull[T]()
	}
	switch {
	case includeMin && includeMax:
		return closedInterval[T]{min, max}
	case includeMin && !includeMax:
		return ROpenInterval[T]{min, max}
	case !includeMin && includeMax:
		return LOpenInterval[T]{min, max}
	case !includeMin && !includeMax:
		return OpenInterval[T]{min, max}

	default:
		panic("impossible configuration")
	}
}

func enumerateFromToStep[T constraints.Ordered](from, to, step T) []T {
	var res []T
	for i := from; i < to; i += step {
		res = append(res, i)
	}
	return res
}

func overlapAtomic[T constraints.Ordered](one, other Interval[T]) bool {
	inclOneMin, inclOneMax := one.IncludeLR()
	inclOtherMin, inclOtherMax := other.IncludeLR()

	if one.Min() > other.Max() ||
		one.Max() < other.Min() ||
		(one.Min() == other.Max() && !(inclOneMin && inclOtherMax)) ||
		(one.Max() == other.Min() && !(inclOneMax && inclOtherMin)) {
		return false
	}

	return true
}
