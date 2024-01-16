package intervals

import "github.com/relvox/iridescence_go/maths"

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
func (ci closedInterval[T]) Len() T { return ci[1] - ci[0] + 1 }

func (ci closedInterval[T]) IsEmpty() bool     { return false }
func (ci closedInterval[T]) IsSingleton() bool { return false }
func (ci closedInterval[T]) IsCompound() bool  { return false }

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

func (ci closedInterval[T]) Intervals() []Interval[T] { return []Interval[T]{ci} }

func (ci closedInterval[T]) Contains(value T) bool { return value >= ci[0] && value <= ci[1] }

// overlapsOne assumes:
//
//	other != nil && !other.IsEmpty() && !other.IsSingleton() && !other.IsCompound()
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

	for _, subInt := range other.Intervals() {
		if !ci.overlapsOne(subInt) {
			continue
		}
		return true
	}
	return false
}

func (ci closedInterval[T]) Equals(other Interval[T]) bool {
	if other == nil || other.IsEmpty() || other.IsSingleton() || other.IsCompound() {
		return false
	}

	return ci[0] == other.Min() && ci[1] == other.Max()
}

func (ci closedInterval[T]) Union(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() ||
		(other.IsSingleton() && ci.Contains(other.Min())) ||
		(!other.IsCompound() && ci.Contains(other.Min()) && ci.Contains(other.Max())) {
		return ci
	}

	var myMin, myMax T = ci[0], ci[1]
	subInts := other.Intervals()
	var res []Interval[T] = make([]Interval[T], 0, len(subInts))
	for si, subInt := range subInts {
		if subInt.Max() < myMin {
			res = append(res, subInt)
			continue
		}
		if subInt.Min() > myMax {
			res = append(res, NewClosed(myMin, myMax))
			res = append(res, subInts[si:]...)
			return RawMerged(res...)
		}
		myMin, myMax = min(myMin, subInt.Min()), max(myMax, subInt.Max())
	}
	res = append(res, NewClosed(myMin, myMax))
	return RawMerged(res...)
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
	if !other.IsCompound() {
		if !ci.overlapsOne(other) {
			return NewEmpty[T]()
		}
		return NewClosed(max(ci[0], other.Min()), min(ci[1], other.Max()))
	}

	return other.Intersection(ci)
}

func (ci closedInterval[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !ci.Overlaps(other) {
		return ci
	}

	if other.IsSingleton() {
		if ci[0] == other.Max() || ci[1] == other.Min() {
			return NewInterval(ci[0], ci[0] != other.Min(), ci[1], ci[1] != other.Max())
		}
		return RawMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}

	if !other.IsCompound() {
		if !ci.overlapsOne(other) {
			return ci
		}

		if other.Contains(ci[0]) {
			if other.Max() >= ci[1] {
				return NewEmpty[T]()
			}
			return NewClosed(other.Max()+1, ci[1])
		}

		if other.Contains(ci[1]) {
			if other.Min() <= ci[0] {
				return NewEmpty[T]()
			}
			return NewClosed(ci[0], other.Min()-1)
		}

		return RawMerged(NewClosed(ci[0], other.Min()-1), NewClosed(other.Max()+1, ci[1]))
	}

	var res Interval[T] = ci
	for _, subInt := range other.Intervals() {
		res = res.Difference(subInt)
	}
	return res
}

func (ci closedInterval[T]) Resize(newSize T, growMode GrowFlags) Interval[T] {
	currentSize := ci.Len()
	minT, maxT := maths.MinValue[T](), maths.MaxValue[T]()
	growth := newSize - currentSize

	left, right := Growths[T](growMode, minT, ci.Min(), growth, ci.Max(), maxT)
	return NewClosed(left, right)
}

func (ci closedInterval[T]) Scale(scale float64, growMode GrowFlags) Interval[T] {
	if scale <= 0 {
		return NewEmpty[T]()
	}

	newSize := T(float64(ci.Len()) * scale)
	return ci.Resize(newSize, growMode)
}

func (ci closedInterval[T]) Translate(offset T, back bool) Interval[T] {
	minT, maxT := maths.MinValue[T](), maths.MaxValue[T]()
	var newMin, newMax T

	if back {
		newMin = ci[0] - offset
		newMax = ci[1] - offset
		if newMin < minT {
			difference := minT - newMin
			newMin = minT
			if newMax-difference < minT {
				newMax = minT
			} else {
				newMax -= difference
			}
		}
	} else {
		newMin = ci[0] + offset
		newMax = ci[1] + offset
		if newMax > maxT {
			difference := newMax - maxT
			newMax = maxT
			if newMin+difference > maxT {
				newMin = maxT
			} else {
				newMin += difference
			}
		}
	}

	return NewClosed(newMin, newMax)
}
