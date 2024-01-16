package test

import (
	"fmt"
	"testing"

	"github.com/relvox/iridescence_go/maths/integer/intervals"
	"github.com/stretchr/testify/assert"
)

type TestInterval = intervals.Interval[uint16]

var (
	Merge = intervals.NewMerged[uint16]
	Raw   = intervals.RawMerged[uint16]
	New   = intervals.NewClosed[uint16]
	Point = intervals.NewSingleton[uint16]
)

func Test_Difference(t *testing.T) {
	tests := []struct {
		input      TestInterval
		difference TestInterval
		expected   TestInterval
	}{
		{
			input:      Merge(New(69, 109)),
			difference: Point(58),
			expected:   Merge(New(69, 109)),
		},
		{
			input:      Merge(New(10, 67)),
			difference: Point(58),
			expected:   Merge(New(10, 57), New(59, 67)),
		},
		{
			input:      Merge(Merge(New(10, 67))),
			difference: Point(58),
			expected:   Merge(New(10, 57), New(59, 67)),
		},
		{
			input:      Merge(New(0, 10), New(20, 30)),
			difference: Point(5),
			expected:   Merge(New(0, 4), New(6, 10), New(20, 30)),
		},
	}
	for ti, tt := range tests {
		t.Run(fmt.Sprintf("%d", ti), func(t *testing.T) {
			actual := tt.input.Difference(tt.difference)
			if !assert.Equal(t, tt.expected, actual) {
				t.FailNow()
			}
		})
	}
}
