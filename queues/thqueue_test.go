package queues_test

import (
	"testing"

	"github.com/relvox/iridescence_go/queues"
	"github.com/stretchr/testify/assert"
)

var arr = []int{6, 5, 4, 3, 2}

func Test(t *testing.T) {
	q := queues.NewTHQueue[int](10)
	ix := 0
	ie := 0
	for i := 0; i < 1000; i++ {
		q.Enq(arr[ix])
		ix = (ix + 1) % len(arr)

		q.Enq(arr[ix])
		ix = (ix + 1) % len(arr)

		a := q.Deq()
		if arr[ie] != a {
			assert.Equal(t, arr[ie], a, "%d: %v", i, q)
			t.FailNow()
		}
		ie = (ie + 1) % len(arr)
	}
}

func BenchmarkEnqueue(b *testing.B) {
	q := queues.NewTHQueue[int](10)
	for i := 0; i < b.N; i++ {
		q.Enq(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := queues.NewTHQueue[int](10)
	for i := 0; i < b.N; i++ {
		q.Enq(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Deq()
	}
}

func BenchmarkMixed(b *testing.B) {
	q := queues.NewTHQueue[int](10)
	for i := 0; i < b.N; i++ {
		q.Enq(i)
		if i%2 == 0 {
			q.Deq()
		}
	}
}

func BenchmarkMediumData(b *testing.B) {
	q := queues.NewTHQueue[int](10)
	for i := 0; i < 1_000; i++ {
		q.Enq(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enq(i)
		q.Deq()
	}
}

func BenchmarkLargeData(b *testing.B) {
	q := queues.NewTHQueue[int](10)
	for i := 0; i < 1_000_000; i++ {
		q.Enq(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enq(i)
		q.Deq()
	}
}
