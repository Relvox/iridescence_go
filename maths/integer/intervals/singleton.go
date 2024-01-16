package intervals

import "github.com/relvox/iridescence_go/maths"

type singletonInterval[T Number] [1]T

func NewSingleton[T Number](v T) Interval[T] {
	return singletonInterval[T]{v}
}

func (si singletonInterval[T]) Min() T { return si[0] }
func (si singletonInterval[T]) Max() T { return si[0] }
func (si singletonInterval[T]) Len() T { return 1 }

func (si singletonInterval[T]) IsEmpty() bool     { return false }
func (si singletonInterval[T]) IsSingleton() bool { return true }
func (si singletonInterval[T]) IsCompound() bool  { return false }

func (si singletonInterval[T]) Enumerate(_ T) []T        { return []T{si[0]} }
func (si singletonInterval[T]) Intervals() []Interval[T] { return []Interval[T]{si} }

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
	otherInts := other.Intervals()
	var res []Interval[T]
	var firstAfter int
	for oi, otherInt := range otherInts {
		if otherInt.Max() < si.Min() {
			res = append(res, otherInt)
			continue
		}
		firstAfter = oi
		break
	}
	res = append(res, si)
	res = append(res, otherInts[firstAfter:]...)
	return RawMerged(res...)
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

func (si singletonInterval[T]) Resize(newSize T, growMode GrowFlags) Interval[T] {
	if newSize <= 1 {
		return si
	}

	left, right := Growths[T](growMode,
		maths.MinValue[T](), si.Min(), newSize-1, si.Max(), maths.MaxValue[T]())

	return NewClosed(left, right)
}

func (si singletonInterval[T]) Scale(scale float64, growMode GrowFlags) Interval[T] {
	newSize := T(scale * float64(si.Len()))
	return si.Resize(newSize, growMode)
}

func (si singletonInterval[T]) Translate(offset T, back bool) Interval[T] {
	minT := maths.MinValue[T]()
	if back {
		leftSlack := si[0] - minT
		if offset >= leftSlack {
			return NewSingleton(minT)
		}
		return NewSingleton(si[0] - offset)
	}
	maxT := maths.MaxValue[T]()
	rightSlack := maxT - si[0]
	if offset >= rightSlack {
		return NewSingleton(maxT)
	}
	return NewSingleton(si[0] + offset)
}
