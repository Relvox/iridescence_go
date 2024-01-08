package intervals

import "golang.org/x/exp/constraints"

type MergedIntervals[T constraints.Ordered] []Interval[T]

func NewMerged[T constraints.Ordered](intervals ...Interval[T]) Interval[T] {
	return MergedIntervals[T]{}
}

func (ni MergedIntervals[T]) Min() T                  { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Max() T                  { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) IncludeLR() (bool, bool) { panic("NOT IMPLEMENTED") }

func (ni MergedIntervals[T]) IsEmpty() bool     { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) IsSingleton() bool { panic("NOT IMPLEMENTED") }

func (ni MergedIntervals[T]) Enumerate(step T) []T     { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Boundaries() []T          { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Intervals() []Interval[T] { panic("NOT IMPLEMENTED") }

func (ni MergedIntervals[T]) Contains(value T) bool           { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Overlaps(other Interval[T]) bool { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Equals(other Interval[T]) bool   { panic("NOT IMPLEMENTED") }

func (ni MergedIntervals[T]) Union(other Interval[T]) Interval[T]        { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Intersection(other Interval[T]) Interval[T] { panic("NOT IMPLEMENTED") }
func (ni MergedIntervals[T]) Difference(other Interval[T]) Interval[T]   { panic("NOT IMPLEMENTED") }
