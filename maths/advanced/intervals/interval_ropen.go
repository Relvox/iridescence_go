package intervals

import (
	"golang.org/x/exp/constraints"

	"github.com/relvox/iridescence_go/maths"
)

type ROpenInterval[T constraints.Ordered] [2]T

func NewROpen[T constraints.Ordered](min, max T) Interval[T] {
	if max <= min {
		return NewNull[T]()
	}
	return ROpenInterval[T]{min, max}
}

func (roi ROpenInterval[T]) Min() T                  { return roi[0] }
func (roi ROpenInterval[T]) Max() T                  { return roi[1] }
func (roi ROpenInterval[T]) IncludeLR() (bool, bool) { return true, false }

func (roi ROpenInterval[T]) IsEmpty() bool     { return false }
func (roi ROpenInterval[T]) IsSingleton() bool { return false }

func (roi ROpenInterval[T]) Enumerate(step T) []T {
	if step == *new(T) {
		step = maths.One[T]()
	}
	return enumerateFromToStep(roi[0], roi[1], step)
}

func (roi ROpenInterval[T]) Intervals() []Interval[T] { return nil }

func (roi ROpenInterval[T]) Contains(value T) bool { return value >= roi[0] && value < roi[1] }

func (roi ROpenInterval[T]) overlapsOne(other Interval[T]) bool {
	_, inclMax := other.IncludeLR()
	if roi.Min() > other.Max() || (roi.Min() == other.Max() && !inclMax) {
		return false
	}
	if roi.Max() <= other.Min() {
		return false
	}
	return true
}

func (roi ROpenInterval[T]) Overlaps(other Interval[T]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return roi.Contains(other.Min())
	}
	subIntervals := other.Intervals()
	if subIntervals == nil {
		return roi.overlapsOne(other)
	}
	for _, subItv := range subIntervals {
		if !roi.overlapsOne(subItv) {
			continue
		}
		return true
	}
	return false
}

func (roi ROpenInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.Intervals() != nil {
		return false
	}

	if inclMin, inclMax := other.IncludeLR(); !inclMin || !inclMax {
		return false
	}
	return roi[0] == other.Min() && roi[1] == other.Max()
}

func (roi ROpenInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && roi.Contains(other.Min())) {
		return roi
	}

	return NewMerged(other, roi)
}

func (roi ROpenInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return other
	}
	if other.IsSingleton() {
		if roi.Contains(other.Max()) {
			return other
		}
		return NewNull[T]()
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !roi.overlapsOne(other) {
			return NewNull[T]()
		}
		newMin := max(roi[0], other.Min())
		newMax := min(roi[1], other.Max())
		return NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax))
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !roi.overlapsOne(subItv) {
			continue
		}
		newMin := max(roi[0], other.Min())
		newMax := min(roi[1], other.Max())
		res = append(res, NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax)))
	}
	return NewMerged(res...)
}

func (roi ROpenInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return roi
	}

	if other.IsSingleton() {
		if !roi.Contains(other.Max()) {
			return roi
		}
		if roi[0] == other.Min() {
			return LOpenInterval[T](roi)
		}
		if roi[1] == other.Max() {
			return ROpenInterval[T](roi)
		}
		return NewMerged(NewROpen(roi.Min(), other.Max()), NewLOpen(other.Min(), roi.Max()))
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !roi.overlapsOne(other) {
			return roi
		}

		if other.Contains(roi.Min()) {
			if !roi.Contains(other.Max()) {
				return NewNull[T]()
			}
			if other.Contains(roi.Max()) {
				return NewNull[T]()
			}
			return NewSingleton(roi.Max())
		}

		minMin := max(roi[0], other.Min())
		minMax := min(roi[1], other.Max())
		// maxMin := max(other.Min(), roi[0])
		maxMax := min(other.Max(), roi[1])

		if minMin == minMax {
			return NewNull[T]()
		}

		if minMin == roi[0] && maxMax == roi[1] {
			// other is fully contained within roi
			// Create two open intervals on either side of other
			return NewMerged(NewROpen(roi.Min(), other.Max()), NewLOpen(other.Min(), roi.Max()))
		}

		return NewClosed(minMin, maxMax) // Remaining case: a closed interval within roi
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !roi.overlapsOne(subItv) {
			continue
		}
		// Calculate the difference between roi and each sub-interval
		diff := roi.Difference(subItv)
		res = append(res, diff)
	}
	return NewMerged(res...)
}
