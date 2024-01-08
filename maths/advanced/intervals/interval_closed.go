package intervals

import (
	"golang.org/x/exp/constraints"

	"github.com/relvox/iridescence_go/maths"
)

type closedInterval[T constraints.Ordered] [2]T

func NewClosed[T constraints.Ordered](min, max T) Interval[T] {
	if max < min {
		return NewNull[T]()
	}
	if max == min {
		return NewSingleton(min)
	}

	return closedInterval[T]{min, max}
}

func (ci closedInterval[T]) Min() T                  { return ci[0] }
func (ci closedInterval[T]) Max() T                  { return ci[1] }
func (ci closedInterval[T]) IncludeLR() (bool, bool) { return true, true }

func (ci closedInterval[T]) IsEmpty() bool     { return false }
func (ci closedInterval[T]) IsSingleton() bool { return false }

func (ci closedInterval[T]) Enumerate(step T) []T {
	if step == *new(T) {
		step = maths.One[T]()
	}
	return enumerateFromToStep(ci[0], ci[1]+step, step)
}

func (ci closedInterval[T]) Intervals() []Interval[T] { return nil }

func (ci closedInterval[T]) Contains(value T) bool { return value >= ci[0] && value <= ci[1] }

func (ci closedInterval[T]) overlapsOne(other Interval[T]) bool {
	inclMin, inclMax := other.IncludeLR()
	if ci.Min() > other.Max() || (ci.Min() == other.Max() && !inclMax) {
		return false
	}
	if ci.Max() < other.Min() || (ci.Max() == other.Min() && !inclMin) {
		return false
	}
	return true
}

func (ci closedInterval[T]) Overlaps(other Interval[T]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return ci.Contains(other.Min())
	}
	subIntervals := other.Intervals()
	if subIntervals == nil {
		return ci.overlapsOne(other)
	}
	for _, subItv := range subIntervals {
		if !ci.overlapsOne(subItv) {
			continue
		}
		return true
	}
	return false
}

func (ci closedInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.Intervals() != nil {
		return false
	}

	if inclMin, inclMax := other.IncludeLR(); !inclMin || !inclMax {
		return false
	}
	return ci[0] == other.Min() && ci[1] == other.Max()
}

func (ci closedInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && ci.Contains(other.Min())) {
		return ci
	}

	return NewMerged(other, ci)
}

func (ci closedInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return other
	}
	if other.IsSingleton() {
		if ci.Contains(other.Max()) {
			return other
		}
		return NewNull[T]()
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !ci.overlapsOne(other) {
			return NewNull[T]()
		}
		newMin := max(ci[0], other.Min())
		newMax := min(ci[1], other.Max())
		return NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax))
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !ci.overlapsOne(subItv) {
			continue
		}
		newMin := max(ci[0], other.Min())
		newMax := min(ci[1], other.Max())
		res = append(res, NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax)))
	}
	return NewMerged(res...)
}

func (ci closedInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return ci
	}

	if other.IsSingleton() {
		if !ci.Contains(other.Max()) {
			return ci
		}
		if ci[0] == other.Min() {
			return LOpenInterval[T](ci)
		}
		if ci[1] == other.Max() {
			return ROpenInterval[T](ci)
		}
		return NewMerged(NewROpen(ci.Min(), other.Max()), NewLOpen(other.Min(), ci.Max()))
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !ci.overlapsOne(other) {
			return ci
		}

		if other.Contains(ci.Min()) {
			if !ci.Contains(other.Max()) {
				return NewNull[T]()
			}
			if other.Contains(ci.Max()) {
				return NewNull[T]()
			}
			return NewSingleton(ci.Max())
		}

		minMin := max(ci[0], other.Min())
		minMax := min(ci[1], other.Max())
		// maxMin := max(other.Min(), ci[0])
		maxMax := min(other.Max(), ci[1])

		if minMin == minMax {
			return NewNull[T]()
		}

		if minMin == ci[0] && maxMax == ci[1] {
			// other is fully contained within ci
			// Create two open intervals on either side of other
			return NewMerged(NewROpen(ci.Min(), other.Max()), NewLOpen(other.Min(), ci.Max()))
		}

		return NewClosed(minMin, maxMax) // Remaining case: a closed interval within ci
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !ci.overlapsOne(subItv) {
			continue
		}
		// Calculate the difference between ci and each sub-interval
		diff := ci.Difference(subItv)
		res = append(res, diff)
	}
	return NewMerged(res...)
}
