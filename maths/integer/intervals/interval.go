package intervals

import (
	"golang.org/x/exp/constraints"
)

type Number = constraints.Integer

type Interval[T Number] interface {
	Min() T
	Max() T

	IsEmpty() bool
	IsSingleton() bool

	Enumerate(step T) []T
	Intervals() []Interval[T]

	Contains(value T) bool
	Overlaps(other Interval[T]) bool
	Equals(other Interval[T]) bool

	Union(other Interval[T]) Interval[T]
	Intersection(other Interval[T]) Interval[T]
	Difference(other Interval[T]) Interval[T]
}
