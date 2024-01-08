package intervals

type closedInterval[T Number] [2]T

func NewClosed[T Number](min, max T) Interval[T] {
	if max < min {
		return NewEmpty[T]()
	}
	if max == min {
		return NewSingleton(min)
	}

	return closedInterval[T]{min, max}
}

func NewInterval[T Number](min T, includeMin bool, max T, includeMax bool) Interval[T] {
	if max == min && includeMin && includeMax {
		return NewSingleton(min)
	}
	if max <= min {
		return NewEmpty[T]()
	}
	switch {
	case includeMin && includeMax:
		return closedInterval[T]{min, max}
	case includeMax:
		return closedInterval[T]{min + 1, max}
	case includeMin:
		return closedInterval[T]{min, max - 1}
	default:
		return closedInterval[T]{min + 1, max - 1}
	}
}

func (ci closedInterval[T]) Min() T { return ci[0] }
func (ci closedInterval[T]) Max() T { return ci[1] }

func (ci closedInterval[T]) IsEmpty() bool     { return false }
func (ci closedInterval[T]) IsSingleton() bool { return false }

func (ci closedInterval[T]) Enumerate(step T) []T {
	if step == 0 {
		step = 1
	}

	var res []T
	for i := ci[0]; i <= ci[1]; i += step {
		res = append(res, i)
	}
	return res
}

func (ci closedInterval[T]) Intervals() []Interval[T] { return nil }

func (ci closedInterval[T]) Contains(value T) bool { return value >= ci[0] && value <= ci[1] }

// overlapsOne assumes:
//
//	other != nil && !other.IsEmpty() && !other.IsSingleton()
func (ci closedInterval[T]) overlapsOne(other Interval[T]) bool {
	return ci[0] <= other.Max() && ci[1] >= other.Min()
}

func (ci closedInterval[T]) Overlaps(other Interval[T]) bool {
	if other == nil || other.IsEmpty() {
		return false
	}
	if other.IsSingleton() {
		return ci.Contains(other.Min())
	}

	if subIntervals := other.Intervals(); subIntervals == nil {
		return ci.overlapsOne(other)
	}

	return other.Overlaps(ci)
}

func (ci closedInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.Intervals() != nil {
		return false
	}

	return ci[0] == other.Min() && ci[1] == other.Max()
}

func (ci closedInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && ci.Contains(other.Min())) ||
		(other.Intervals() == nil && ci.Contains(other.Min()) && ci.Contains(other.Max())) {
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
		return NewEmpty[T]()
	}

	if subIntervals := other.Intervals(); subIntervals == nil {
		if !ci.overlapsOne(other) {
			return NewEmpty[T]()
		}
		return NewClosed(max(ci[0], other.Min()), min(ci[1], other.Max()))
	}

	return other.Intersection(ci)
}

func (ci closedInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Overlaps(ci) {
		return ci
	}

	if other.IsSingleton() {
		if ci[0] == other.Max() || ci[1] == other.Min() {
			return NewInterval(ci[0], ci[0] != other.Min(), ci[1], ci[1] != other.Max())
		}
		return NewMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}

	subIntervals := other.Intervals()
	if subIntervals == nil {
		if !ci.overlapsOne(other) {
			return ci
		}

		if other.Contains(ci[0]) {
			if !ci.Contains(other.Max()) || other.Contains(ci[1]) {
				return NewEmpty[T]()
			}
			return NewClosed(other.Max()+1, ci[1])
		}

		if other.Contains(ci[1]) {
			if !ci.Contains(other.Min()) || other.Contains(ci[0]) {
				return NewEmpty[T]()
			}
			return NewClosed(ci[0], other.Min()-1)
		}
		return NewMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}
	var res Interval[T] = ci
	for _, subInt := range subIntervals {
		res = res.Difference(subInt)
	}
	return res
}
