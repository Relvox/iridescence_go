package intervals

import (
	"github.com/relvox/iridescence_go/utils"
)

type mergedIntervals[T Number] []Interval[T]

func RawMerged[T Number](intervals ...Interval[T]) Interval[T] {
	intervals = utils.CropElements(intervals, nil, NewEmpty[T]())
	if len(intervals) == 0 {
		return NewEmpty[T]()
	}
	if len(intervals) == 1 {
		return NewClosed(intervals[0].Min(), intervals[0].Max())
	}
	return mergedIntervals[T](intervals)
}

func NewMerged[T Number](intervals ...Interval[T]) Interval[T] {
	intervals = utils.CropElements(intervals, nil, NewEmpty[T]())
	if len(intervals) == 0 {
		return NewEmpty[T]()
	}
	if len(intervals) == 1 {
		return intervals[0]
	}

	rawIntervals := make([][2]T, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		subInts := intervals[i].Intervals()
		if subInts == nil {
			rawIntervals = append(rawIntervals, [2]T{intervals[i].Min(), intervals[i].Max()})
			continue
		}
		for _, subInt := range subInts {
			rawIntervals = append(rawIntervals, [2]T{subInt.Min(), subInt.Max()})
		}
	}
	rawResult := FindCover(rawIntervals)
	var res []Interval[T] = make([]Interval[T], len(rawResult)/2)
	for i := 0; i < len(res); i++ {
		res[i] = NewClosed(rawResult[i*2], rawResult[i*2+1])
	}
	return mergedIntervals[T](res)
}

func (mis mergedIntervals[T]) Min() T { return mis[0].Min() }
func (mis mergedIntervals[T]) Max() T { return mis[len(mis)-1].Max() }
func (mis mergedIntervals[T]) Len() T {
	var result T
	for _, v := range mis.Intervals() {
		result += v.Len()
	}
	return result
}

func (mis mergedIntervals[T]) IsEmpty() bool     { return false }
func (mis mergedIntervals[T]) IsSingleton() bool { return false }
func (mis mergedIntervals[T]) IsCompound() bool  { return true }

func (mis mergedIntervals[T]) Enumerate(step T) []T {
	var res []T
	for _, subInt := range mis {
		res = append(res, subInt.Enumerate(step)...)
	}
	return res
}

func (mis mergedIntervals[T]) Intervals() []Interval[T] {
	for _, v := range mis {
		if v.IsCompound() {
			panic("DO NOT WANT!")
		}
	}
	return mis
}

func (mis mergedIntervals[T]) Contains(value T) bool {
	for _, subInt := range mis {
		if subInt.Contains(value) {
			return true
		}
	}
	return false
}

func (mis mergedIntervals[T]) Overlaps(other Interval[T]) bool {
	for _, subInt := range mis {
		if subInt.Overlaps(other) {
			return true
		}
	}
	return false
}

func (mis mergedIntervals[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || !other.IsCompound() {
		return false
	}
	subInts := other.Intervals()
	if len(subInts) != len(mis) {
		return false
	}
	for si, subInt := range subInts {
		if !subInt.Equals(mis[si]) {
			return false
		}
	}
	return true

}

func (mis mergedIntervals[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return mis
	}
	if !other.IsCompound() {
		return other.Union(mis)
	}

	return NewMerged(other, mis)
}

func (mis mergedIntervals[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return NewEmpty[T]()
	}
	var res []Interval[T]
	for _, subInt := range mis {
		res = append(res, other.Intersection(subInt))
	}
	return RawMerged(res...)
}

func (mis mergedIntervals[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Overlaps(mis) {
		return mis
	}

	var res []Interval[T]
	if !other.IsCompound() {
		for _, subInt := range mis {
			res = append(res, subInt.Difference(other).Intervals()...)
		}
		return RawMerged(res...)
	}
	otherIntervals := other.Intervals()
	for _, subInt := range mis {
		for _, otherSubInt := range otherIntervals {
			subInt = subInt.Difference(otherSubInt)
		}
		res = append(res, subInt)
	}
	return RawMerged(res...)
}

func (mis mergedIntervals[T]) Resize(newSize T, growMode GrowFlags) Interval[T] {
	panic("#TODO: not implemented")
}
func (mis mergedIntervals[T]) Scale(scale float64, growMode GrowFlags) Interval[T] {
	panic("#TODO: not implemented")
}
func (mis mergedIntervals[T]) Translate(offset T, back bool) Interval[T] {
	panic("#TODO: not implemented")
}
