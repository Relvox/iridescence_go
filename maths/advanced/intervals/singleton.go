package intervals

import (
	"golang.org/x/exp/constraints"
)

type singletonInterval[T constraints.Ordered] struct{ Value T }

func NewSingleton[T constraints.Ordered](v T) Interval[T] {
	return singletonInterval[T]{v}
}

func (si singletonInterval[T]) Min() T                  { return si.Value }
func (si singletonInterval[T]) Max() T                  { return si.Value }
func (ni singletonInterval[T]) IncludeLR() (bool, bool) { return true, true }

func (si singletonInterval[T]) IsEmpty() bool     { return false }
func (si singletonInterval[T]) IsSingleton() bool { return true }

func (si singletonInterval[T]) Enumerate(_ T) []T        { return []T{si.Value} }
func (si singletonInterval[T]) Intervals() []Interval[T] { return nil }

func (si singletonInterval[T]) Contains(value T) bool           { return si.Value == value }
func (si singletonInterval[T]) Overlaps(other Interval[T]) bool { return other.Contains(si.Value) }
func (si singletonInterval[T]) Equals(other Interval[T]) bool {
	return other != nil && other.IsSingleton() && other.Contains(si.Value)
}

func (si singletonInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return si
	}
	if other.Contains(si.Value) {
		return other
	}
	return NewMerged(si, other)
}

func (si singletonInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Contains(si.Value) {
		return NewNull[T]()
	}
	return si
}

func (si singletonInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Contains(si.Value) {
		return si
	}
	return NewNull[T]()
}

type foo32 float32
type foo64 float64

type flap interface {
	foo32 | float64
}