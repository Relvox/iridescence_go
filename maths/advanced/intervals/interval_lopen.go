package intervals

import (
	"golang.org/x/exp/constraints"

	"github.com/relvox/iridescence_go/maths"
)

type LOpenInterval[T constraints.Ordered] [2]T

func NewLOpen[T constraints.Ordered](min, max T) Interval[T] {
	if max <= min {
		return NewNull[T]()
	}
	return LOpenInterval[T]{min, max}
}

func (loi LOpenInterval[T]) Min() T                  { return loi[0] }
func (loi LOpenInterval[T]) Max() T                  { return loi[1] }
func (loi LOpenInterval[T]) IncludeLR() (bool, bool) { return false, true }

func (loi LOpenInterval[T]) IsEmpty() bool     { return false }
func (loi LOpenInterval[T]) IsSingleton() bool { return false }

func (loi LOpenInterval[T]) Enumerate(step T) []T {
	if step == *new(T) {
		step = maths.One[T]()
	}
	return enumerateFromToStep(loi[0]+step, loi[1]+step, step)
}

func (loi LOpenInterval[T]) Intervals() []Interval[T] { return nil }

func (loi LOpenInterval[T]) Contains(value T) bool { return value > loi[0] && value <= loi[1] }

func (loi LOpenInterval[T]) overlapsOne(other Interval[T]) bool {
	inclMin, _ := other.IncludeLR()
	if loi.Min() >= other.Max() {
		return false
	}
	if loi.Max() < other.Min() || (loi.Max() == other.Min() && !inclMin) {
		return false
	}
	return true
}

func (loi LOpenInterval[T]) Overlaps(other Interval[T]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return loi.Contains(other.Min())
	}
	subIntervals := other.Intervals()
	if subIntervals == nil {
		return loi.overlapsOne(other)
	}
	for _, subItv := range subIntervals {
		if !loi.overlapsOne(subItv) {
			continue
		}
		return true
	}
	return false
}

func (loi LOpenInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.Intervals() != nil {
		return false
	}

	if inclMin, inclMax := other.IncludeLR(); !inclMin || !inclMax {
		return false
	}
	return loi[0] == other.Min() && loi[1] == other.Max()
}

func (loi LOpenInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && loi.Contains(other.Min())) {
		return loi
	}

	return NewMerged(other, loi)
}

func (loi LOpenInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return other
	}
	if other.IsSingleton() {
		if loi.Contains(other.Max()) {
			return other
		}
		return NewNull[T]()
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !loi.overlapsOne(other) {
			return NewNull[T]()
		}
		newMin := max(loi[0], other.Min())
		newMax := min(loi[1], other.Max())
		return NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax))
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !loi.overlapsOne(subItv) {
			continue
		}
		newMin := max(loi[0], other.Min())
		newMax := min(loi[1], other.Max())
		res = append(res, NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax)))
	}
	return NewMerged(res...)
}

func (loi LOpenInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return loi
	}

	if other.IsSingleton() {
		if !loi.Contains(other.Max()) {
			return loi
		}
		if loi[0] == other.Min() {
			return LOpenInterval[T](loi)
		}
		if loi[1] == other.Max() {
			return ROpenInterval[T](loi)
		}
		return NewMerged(NewROpen(loi.Min(), other.Max()), NewLOpen(other.Min(), loi.Max()))
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !loi.overlapsOne(other) {
			return loi
		}

		if other.Contains(loi.Min()) {
			if !loi.Contains(other.Max()) {
				return NewNull[T]()
			}
			if other.Contains(loi.Max()) {
				return NewNull[T]()
			}
			return NewSingleton(loi.Max())
		}

		minMin := max(loi[0], other.Min())
		minMax := min(loi[1], other.Max())
		// maxMin := max(other.Min(), loi[0])
		maxMax := min(other.Max(), loi[1])

		if minMin == minMax {
			return NewNull[T]()
		}

		if minMin == loi[0] && maxMax == loi[1] {
			// other is fully contained within loi
			// Create two open intervals on either side of other
			return NewMerged(NewROpen(loi.Min(), other.Max()), NewLOpen(other.Min(), loi.Max()))
		}

		return NewClosed(minMin, maxMax) // Remaining case: a closed interval within loi
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !loi.overlapsOne(subItv) {
			continue
		}
		// Calculate the difference between loi and each sub-interval
		diff := loi.Difference(subItv)
		res = append(res, diff)
	}
	return NewMerged(res...)
}
