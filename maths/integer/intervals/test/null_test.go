package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/relvox/iridescence_go/maths"
	"github.com/relvox/iridescence_go/maths/integer/intervals"
)

func NullInterval_Tester[T intervals.Number](t *testing.T) {
	zero := *new(T)
	one := T(1)
	req := require.New(t)

	// Values for testing, including edge cases and potential problem points
	values := []T{
		zero,
		one,
		-one,
		2 * one,
		2 * -one,
		2*one + one,
		2*-one - one,
		// Additional values for edge cases
		maths.MinValue[T](),
		maths.MaxValue[T](),
	}

	t.Run(fmt.Sprintf("NullInterval[%T]", zero), func(t *testing.T) {
		t.Run("Sanity", func(t *testing.T) {
			testInterval := intervals.NewEmpty[T]()
			t.Run("Min", func(t *testing.T) {
				req.Equal(zero, testInterval.Min(), "Min of NullInterval should be zero")
			})
			t.Run("Max", func(t *testing.T) {
				req.Equal(zero, testInterval.Max(), "Max of NullInterval should be zero")
			})

			t.Run("IsEmpty", func(t *testing.T) {
				req.True(testInterval.IsEmpty(), "NullInterval should be empty")
			})
			t.Run("IsEmpty and IsSingleton", func(t *testing.T) {
				req.False(testInterval.IsSingleton(), "NullInterval should not be a singleton")
			})

			t.Run("Enumerate", func(t *testing.T) {
				for _, step := range values {
					t.Run(fmt.Sprint("Step=", step), func(t *testing.T) {
						req.Nil(testInterval.Enumerate(step), "Enumerate(%d) should return nil for NullInterval", step)
					})
				}
			})

			// t.Run("Intervals", func(t *testing.T) {
			// 	req.Nil(testInterval.FLOOP(), "Intervals should return nil for NullInterval")
			// })

			t.Run("Contains", func(t *testing.T) {
				for _, val := range values {
					t.Run(fmt.Sprint(val), func(t *testing.T) {
						req.False(testInterval.Contains(val), "NullInterval should not contain any value")
					})
				}
			})
		})

		t.Run("Vs_Empty", func(t *testing.T) {
			testInterval := intervals.NewEmpty[T]()
			otherEmpty := intervals.NewEmpty[T]()

			t.Run("Overlaps", func(t *testing.T) {
				req.False(testInterval.Overlaps(testInterval), "NullInterval should not overlap with self")
				req.False(testInterval.Overlaps(otherEmpty), "NullInterval should not overlap with other empty interval")
			})

			t.Run("Equals", func(t *testing.T) {
				req.True(testInterval.Equals(testInterval), "NullInterval should equal itself")
				req.True(testInterval.Equals(otherEmpty), "NullInterval should equal another empty interval")
			})

			t.Run("Union", func(t *testing.T) {
				req.True(testInterval.Union(testInterval).IsEmpty(), "Union with self should be empty")
				req.False(testInterval.Union(testInterval).IsSingleton(), "Union with self should not be singleton")
				// req.Nil(testInterval.Union(testInterval).FLOOP(), "Union with self should not have intervals")

				req.True(testInterval.Union(otherEmpty).IsEmpty(), "Union with other Null should be empty")
				req.False(testInterval.Union(otherEmpty).IsSingleton(), "Union with other Null should not be singleton")
				// req.Nil(testInterval.Union(otherEmpty).FLOOP(), "Union with other Null should not have intervals")
			})

			t.Run("Intersection", func(t *testing.T) {
				req.True(testInterval.Intersection(testInterval).IsEmpty(), "Intersection with self should be empty")
				req.False(testInterval.Intersection(testInterval).IsSingleton(), "Intersection with self should not be singleton")
				// req.Nil(testInterval.Intersection(testInterval).FLOOP(), "Intersection with self should not have intervals")

				req.True(testInterval.Intersection(otherEmpty).IsEmpty(), "Intersection with other Null should be empty")
				req.False(testInterval.Intersection(otherEmpty).IsSingleton(), "Intersection with other Null should not be singleton")
				// req.Nil(testInterval.Intersection(otherEmpty).FLOOP(), "Intersection with other Null should not have intervals")
			})

			t.Run("Difference", func(t *testing.T) {
				req.True(testInterval.Difference(testInterval).IsEmpty(), "Difference with self should be empty")
				req.False(testInterval.Difference(testInterval).IsSingleton(), "Difference with self should not be singleton")
				// req.Nil(testInterval.Difference(testInterval).FLOOP(), "Difference with self should not have intervals")

				req.True(testInterval.Difference(otherEmpty).IsEmpty(), "Difference with other Null should be empty")
				req.False(testInterval.Difference(otherEmpty).IsSingleton(), "Difference with other Null should not be singleton")
				// req.Nil(testInterval.Difference(otherEmpty).FLOOP(), "Difference with other Null should not have intervals")
			})
		})

		t.Run("Vs_Single", func(t *testing.T) {
			testInterval := intervals.NewEmpty[T]()
			single := intervals.NewSingleton[T](zero) // Create a singleton interval

			t.Run("Overlaps", func(t *testing.T) {
				req.False(testInterval.Overlaps(single), "NullInterval should not overlap with a singleton")
			})

			t.Run("Equals", func(t *testing.T) {
				req.False(testInterval.Equals(single), "NullInterval should not equal a singleton")
			})

			t.Run("Union", func(t *testing.T) {
				union := testInterval.Union(single)
				req.Equal(single, union, "Union with singleton should result in the singleton")
			})

			t.Run("Intersection", func(t *testing.T) {
				intersection := testInterval.Intersection(single)
				req.Equal(testInterval, intersection, "Intersection with singleton should result in the null interval")
			})

			t.Run("Difference", func(t *testing.T) {
				difference := testInterval.Difference(single)
				req.Equal(testInterval, difference, "Difference with singleton should result in the null interval")
			})
		})

		t.Run("Vs_Merged", func(t *testing.T) {
			// testInterval := intervals.NewEmpty[T]()
			// merged := intervals.NewMerged(intervals.NewSingleton[T](one), intervals.NewSingleton[T](2*one)) // Create a merged interval

			// t.Run("Overlaps", func(t *testing.T) {
			//     req.False(testInterval.Overlaps(merged), "NullInterval should not overlap with a merged interval")
			// })

			// t.Run("Equals", func(t *testing.T) {
			//     req.False(testInterval.Equals(merged), "NullInterval should not equal a merged interval")
			// })

			// t.Run("Union", func(t *testing.T) {
			//     union := testInterval.Union(merged)
			//     req.Equal(merged, union, "Union with merged interval should result in the merged interval")
			// })

			// t.Run("Intersection", func(t *testing.T) {
			//     intersection := testInterval.Intersection(merged)
			//     req.Equal(testInterval, intersection, "Intersection with merged interval should result in the null interval")
			// })

			// t.Run("Difference", func(t *testing.T) {
			//     difference := testInterval.Difference(merged)
			//     req.Equal(testInterval, difference, "Difference with merged interval should result in the null interval")
			// })
		})
	})
}
