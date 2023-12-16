package queues

type THQueue[T any] struct {
	Q    []T
	H, T int
}

func NewTHQueue[T any]() *THQueue[T] {
	return &THQueue[T]{Q: make([]T, 10)}
}

func (q *THQueue[T]) Enq(item T) {
	if len(q.Q) <= q.T {
		if q.H > len(q.Q)/3 {
			copy(q.Q, q.Q[q.H:])
			q.T, q.H = q.T-q.H, 0
		} else {
			newQ := make([]T, len(q.Q)*2)
			copy(newQ, q.Q)
			q.Q = newQ
		}
	}
	q.Q[q.T] = item
	q.T++
}

func (q *THQueue[T]) Deq() T {
	res := q.Q[q.H]
	q.H++
	if q.H == q.T {
		q.Recycle()
	}
	return res
}

func (q *THQueue[T]) IsEmpty() bool { return q.T == q.H }
func (q *THQueue[T]) Recycle()      { q.H, q.T = 0, 0 }
