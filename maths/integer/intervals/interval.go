package intervals

import (
	"golang.org/x/exp/constraints"
)

type Number = constraints.Integer

type GrowFlags int

const (
	GROW_BOTH_OVERFLOW_RIGHT GrowFlags = 0
	GROW_LEFT_OVERFLOW_RIGHT GrowFlags = 1
	GROW_RIGHT_OVERFLOW_LEFT GrowFlags = 2
	GROW_BOTH_OVERFLOW_LEFT  GrowFlags = 3
	GROW_NO_OVERFLOW         GrowFlags = 4
)

func initialGrowths[T Number](gf GrowFlags, growth T) (T, T) {
	switch gf & 3 {
	case 0:
		return growth / 2, growth - (growth / 2)
	case 1:
		return growth, 0
	case 2:
		return 0, growth
	case 3:
		return growth - (growth / 2), growth / 2
	default:
		panic("what how!")
	}
}

func Growths[T Number](gf GrowFlags, leftBound, leftEdge, growth, rightEdge, rightBound T) (T, T) {
	leftGrow, rightGrow := initialGrowths(gf, growth)
	leftSlack, rightSlack := leftEdge-leftBound, rightBound-rightEdge
	var doOverflow bool = gf&4 != 0

	if !doOverflow {
		leftGrow = min(leftSlack, leftGrow)
		rightGrow = min(rightSlack, rightGrow)
	}

	if leftGrow > 0 {
		if leftGrow > leftSlack {
			leftEdge = leftBound
			if doOverflow {
				rightGrow += leftGrow - leftSlack
			}
			leftSlack = 0
		} else {
			leftEdge -= leftGrow
			leftSlack -= leftGrow
		}
		leftGrow = 0
	}

	if rightGrow > 0 {
		if rightGrow > rightSlack {
			rightEdge = rightBound
			if doOverflow {
				leftGrow += rightGrow - rightSlack
			}
			rightSlack = 0
		} else {
			rightEdge += rightGrow
			rightSlack -= rightGrow
		}
		rightGrow = 0
	}

	if leftGrow > 0 {
		if leftGrow > leftSlack {
			return leftBound, rightBound
		}
		leftEdge -= leftGrow
		leftSlack -= leftGrow
		leftGrow = 0
	}

	return leftEdge, rightEdge
}

type Interval[T Number] interface {
	Min() T
	Max() T
	Len() T

	IsEmpty() bool
	IsSingleton() bool
	IsCompound() bool

	Enumerate(step T) []T
	Intervals() []Interval[T]

	Contains(value T) bool
	Overlaps(other Interval[T]) bool
	Equals(other Interval[T]) bool

	Union(other Interval[T]) Interval[T]
	Intersection(other Interval[T]) Interval[T]
	Difference(other Interval[T]) Interval[T]

	Resize(newSize T, growMode GrowFlags) Interval[T]
	Scale(scale float64, growMode GrowFlags) Interval[T]
	Translate(offset T, back bool) Interval[T]
}
