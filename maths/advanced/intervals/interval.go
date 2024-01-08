package intervals

import "golang.org/x/exp/constraints"

type Interval[T constraints.Ordered] interface {
	Min() T
	Max() T
	IncludeLR() (bool, bool)

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
