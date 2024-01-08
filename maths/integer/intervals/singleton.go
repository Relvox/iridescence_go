package intervals

type singletonInterval[T Number] [1]T

func NewSingleton[T Number](v T) Interval[T] {
	return singletonInterval[T]{v}
}

func (si singletonInterval[T]) Min() T { return si[0] }
func (si singletonInterval[T]) Max() T { return si[0] }

func (si singletonInterval[T]) IsEmpty() bool     { return false }
func (si singletonInterval[T]) IsSingleton() bool { return true }

func (si singletonInterval[T]) Enumerate(_ T) []T        { return []T{si[0]} }
func (si singletonInterval[T]) Intervals() []Interval[T] { return nil }

func (si singletonInterval[T]) Contains(value T) bool { return si[0] == value }
func (si singletonInterval[T]) Overlaps(other Interval[T]) bool {
	return other != nil && other.Contains(si[0])
}
func (si singletonInterval[T]) Equals(other Interval[T]) bool {
	return other != nil && other.IsSingleton() && other.Contains(si[0])
}

func (si singletonInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return si
	}
	if other.Contains(si[0]) {
		return other
	}
	return NewMerged(si, other)
}

func (si singletonInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Contains(si[0]) {
		return NewEmpty[T]()
	}
	return si
}

func (si singletonInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Contains(si[0]) {
		return si
	}
	return NewEmpty[T]()
}
