package intervals

type nullInterval[T Number] struct{}

func NewEmpty[T Number]() Interval[T] {
	return nullInterval[T]{}
}

func (ni nullInterval[T]) Min() T { return *new(T) }
func (ni nullInterval[T]) Max() T { return *new(T) }
func (ni nullInterval[T]) Len() T { return *new(T) }

func (ni nullInterval[T]) IsEmpty() bool     { return true }
func (ni nullInterval[T]) IsSingleton() bool { return false }
func (ni nullInterval[T]) IsCompound() bool  { return false }

func (ni nullInterval[T]) Enumerate(_ T) []T        { return nil }
func (ni nullInterval[T]) Intervals() []Interval[T] { return nil }

func (ni nullInterval[T]) Contains(_ T) bool             { return false }
func (ni nullInterval[T]) Overlaps(_ Interval[T]) bool   { return false }
func (ni nullInterval[T]) Equals(other Interval[T]) bool { return other == nil || other.IsEmpty() }

func (ni nullInterval[T]) Union(other Interval[T]) Interval[T]    { return other }
func (ni nullInterval[T]) Intersection(_ Interval[T]) Interval[T] { return ni }
func (ni nullInterval[T]) Difference(_ Interval[T]) Interval[T]   { return ni }

func (ni nullInterval[T]) Resize(_ T, _ GrowFlags) Interval[T]      { return ni }
func (ni nullInterval[T]) Scale(_ float64, _ GrowFlags) Interval[T] { return ni }
func (ni nullInterval[T]) Translate(_ T, _ bool) Interval[T]        { return ni }
