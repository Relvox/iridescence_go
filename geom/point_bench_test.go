package geom_test

import (
	"testing"

	"github.com/relvox/iridescence_go/geom"
)

func Benchmark_Offset(b *testing.B) {
	pt := geom.Point2{}
	// log.Println(b.N)
	for i := 0; i < b.N; i++ {
		pt = pt.Offset(i%10+1, i%10+1)
	}
}
func Benchmark_OffsetPt(b *testing.B) {
	pt := geom.Point2{}
	offsets := []geom.Point2{}
	for i := 0; i < b.N; i++ {
		offsets = append(offsets, geom.Point2{i%10 + 1, i%10 + 1})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pt = pt.OffsetPt(offsets[i])
	}
}
