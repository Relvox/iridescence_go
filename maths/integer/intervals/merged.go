package intervals

type mergedIntervals[T Number] []Interval[T]

func NewMerged[T Number](intervals ...Interval[T]) Interval[T] {
	rawIntervals := make([][2]T, len(intervals))
	for i := 0; i < len(rawIntervals); i++ {
		rawIntervals[i][0], rawIntervals[i][1] = intervals[i].Min(), intervals[i].Max()
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

func (mis mergedIntervals[T]) IsEmpty() bool     { return false }
func (mis mergedIntervals[T]) IsSingleton() bool { return false }

func (mis mergedIntervals[T]) Enumerate(step T) []T {
	var res []T
	for _, subInt := range mis {
		res = append(res, subInt.Enumerate(step)...)
	}
	return res
}

func (mis mergedIntervals[T]) Intervals() []Interval[T] {
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
	if other == nil || other.IsEmpty() || other.IsSingleton() {
		return false
	}
	subInts := other.Intervals()
	if subInts == nil || len(subInts) != len(mis) {
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
	return NewMerged(res...)
}

func (mis mergedIntervals[T]) Difference(other Interval[T]) Interval[T] {
	if other == nil || other.IsEmpty() || !other.Overlaps(mis) {
		return mis
	}
	var res []Interval[T]
	otherIntervals := other.Intervals()
	if otherIntervals == nil {
		for _, subInt := range mis {
			res = append(res, subInt.Difference(other))
		}
	}

	for _, subInt := range mis {
		for _, otherSubInt := range otherIntervals {
			subInt = subInt.Difference(otherSubInt)
		}
		res = append(res, subInt)
	}
	return NewMerged(res...)
}
