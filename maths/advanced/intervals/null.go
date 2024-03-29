package intervals

import (
	"golang.org/x/exp/constraints"

	"github.com/relvox/iridescence_go/sets"
)

type nullInterval[T constraints.Ordered] sets.Unit

func NewNull[T constraints.Ordered]() Interval[T] {
	return nullInterval[T]{}
}

func (ni nullInterval[T]) Min() T                  { return *new(T) }
func (ni nullInterval[T]) Max() T                  { return *new(T) }
func (ni nullInterval[T]) IncludeLR() (bool, bool) { return false, false }

func (ni nullInterval[T]) IsEmpty() bool     { return true }
func (ni nullInterval[T]) IsSingleton() bool { return false }

func (ni nullInterval[T]) Enumerate(step T) []T     { return nil }
func (ni nullInterval[T]) Intervals() []Interval[T] { return nil }

func (ni nullInterval[T]) Contains(value T) bool           { return false }
func (ni nullInterval[T]) Overlaps(other Interval[T]) bool { return false }
func (ni nullInterval[T]) Equals(other Interval[T]) bool   { return other == nil || other.IsEmpty() }

func (ni nullInterval[T]) Union(other Interval[T]) Interval[T]        { return other }
func (ni nullInterval[T]) Intersection(other Interval[T]) Interval[T] { return ni }
func (ni nullInterval[T]) Difference(other Interval[T]) Interval[T]   { return ni }
