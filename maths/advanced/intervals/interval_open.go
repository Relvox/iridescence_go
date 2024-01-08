package intervals

import (
	"golang.org/x/exp/constraints"

	"github.com/relvox/iridescence_go/maths"
)

type OpenInterval[T constraints.Ordered] [2]T

func NewOpen[T constraints.Ordered](min, max T) Interval[T] {
	if max <= min {
		return NewNull[T]()
	}
	return OpenInterval[T]{min, max}
}
func (oi OpenInterval[T]) Min() T                  { return oi[0] }
func (oi OpenInterval[T]) Max() T                  { return oi[1] }
func (oi OpenInterval[T]) IncludeLR() (bool, bool) { return false, false }

func (oi OpenInterval[T]) IsEmpty() bool     { return false }
func (oi OpenInterval[T]) IsSingleton() bool { return false }

func (oi OpenInterval[T]) Enumerate(step T) []T {
	if step == *new(T) {
		step = maths.One[T]()
	}
	return enumerateFromToStep(oi[0]+step, oi[1], step)
}

func (oi OpenInterval[T]) Intervals() []Interval[T] { return nil }

func (oi OpenInterval[T]) Contains(value T) bool { return value > oi[0] && value < oi[1] }

func (oi OpenInterval[T]) overlapsOne(other Interval[T]) bool {
	// inclMin, inclMax := other.IncludeLR()
	if oi.Min() >= other.Max() {
		return false
	}
	if oi.Max() <= other.Min() {
		return false
	}
	return true
}

func (oi OpenInterval[T]) Overlaps(other Interval[T]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return oi.Contains(other.Min())
	}
	subIntervals := other.Intervals()
	if subIntervals == nil {
		return oi.overlapsOne(other)
	}
	for _, subItv := range subIntervals {
		if !oi.overlapsOne(subItv) {
			continue
		}
		return true
	}
	return false
}

func (oi OpenInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.Intervals() != nil {
		return false
	}

	if inclMin, inclMax := other.IncludeLR(); !inclMin || !inclMax {
		return false
	}
	return oi[0] == other.Min() && oi[1] == other.Max()
}

func (oi OpenInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && oi.Contains(other.Min())) {
		return oi
	}

	return NewMerged(other, oi)
}

func (oi OpenInterval[T]) Intersection(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return other
	}
	if other.IsSingleton() {
		if oi.Contains(other.Max()) {
			return other
		}
		return NewNull[T]()
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !oi.overlapsOne(other) {
			return NewNull[T]()
		}
		newMin := max(oi[0], other.Min())
		newMax := min(oi[1], other.Max())
		return NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax))
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !oi.overlapsOne(subItv) {
			continue
		}
		newMin := max(oi[0], other.Min())
		newMax := min(oi[1], other.Max())
		res = append(res, NewInterval(newMin, other.Contains(newMin), newMax, other.Contains(newMax)))
	}
	return NewMerged(res...)
}

func (oi OpenInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() {
		return oi
	}

	if other.IsSingleton() {
		if !oi.Contains(other.Max()) {
			return oi
		}
		if oi[0] == other.Min() {
			return LOpenInterval[T](oi)
		}
		if oi[1] == other.Max() {
			return ROpenInterval[T](oi)
		}
		return NewMerged(NewROpen(oi.Min(), other.Max()), NewLOpen(other.Min(), oi.Max()))
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !oi.overlapsOne(other) {
			return oi
		}

		if other.Contains(oi.Min()) {
			if !oi.Contains(other.Max()) {
				return NewNull[T]()
			}
			if other.Contains(oi.Max()) {
				return NewNull[T]()
			}
			return NewSingleton(oi.Max())
		}

		minMin := max(oi[0], other.Min())
		minMax := min(oi[1], other.Max())
		// maxMin := max(other.Min(), oi[0])
		maxMax := min(other.Max(), oi[1])

		if minMin == minMax {
			return NewNull[T]()
		}

		if minMin == oi[0] && maxMax == oi[1] {
			// other is fully contained within oi
			// Create two open intervals on either side of other
			return NewMerged(NewROpen(oi.Min(), other.Max()), NewLOpen(other.Min(), oi.Max()))
		}

		return NewClosed(minMin, maxMax) // Remaining case: a closed interval within oi
	}

	var res []Interval[T]
	for _, subItv := range subIntervals {
		if !oi.overlapsOne(subItv) {
			continue
		}
		// Calculate the difference between oi and each sub-interval
		diff := oi.Difference(subItv)
		res = append(res, diff)
	}
	return NewMerged(res...)
}
