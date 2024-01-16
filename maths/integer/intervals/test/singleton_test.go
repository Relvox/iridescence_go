package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/relvox/iridescence_go/maths/integer/intervals"
)

func SingletonInterval_Tester[T intervals.Number](t *testing.T) {
	zero := *new(T)
	one := T(1)
	req := require.New(t)
	values := []T{zero, one, -one, 2 * one, 2 * -one, 2*one + one, 2*-one - one}
	testValues := []T{values[0], values[1], values[6]}
	t.Run(fmt.Sprintf("SingletonInterval[%T]", zero), func(t *testing.T) {
		for _, v := range testValues {
			t.Run(fmt.Sprint(v), func(t *testing.T) {
				t.Run("Sanity", func(t *testing.T) {
					testInterval := intervals.NewSingleton(v)
					t.Run("Min", func(t *testing.T) {
						req.Equal(v, testInterval.Min(), "Min of Singleton(%d) should be %d", v, v)
					})
					t.Run("Max", func(t *testing.T) {
						req.Equal(v, testInterval.Max(), "Max of Singleton(%d) should be %d", v, v)
					})

					t.Run("IsEmpty", func(t *testing.T) {
						req.False(testInterval.IsEmpty(), "Singleton(%d) should not be empty", v)
					})
					t.Run("IsEmpty and IsSingleton", func(t *testing.T) {
						req.True(testInterval.IsSingleton(), "Singleton(%d) should be a singleton", v)
					})

					t.Run("Enumerate", func(t *testing.T) {
						for _, step := range values {
							t.Run(fmt.Sprint("Step=", step), func(t *testing.T) {
								req.ElementsMatch([]T{v}, testInterval.Enumerate(step), "Enumerate(%d) should return [%d]", step, v)
							})
						}
					})

					// t.Run("Intervals", func(t *testing.T) {
					// 	req.Nil(testInterval.FLOOP(), "Intervals should return nil for Singleton")
					// })

					t.Run("Contains", func(t *testing.T) {
						for _, val := range values {
							t.Run(fmt.Sprint(val), func(t *testing.T) {
								if v == val {
									req.True(testInterval.Contains(val), "Singleton(%d) should contain %d", v, val)
								} else {
									req.False(testInterval.Contains(val), "Singleton(%d) should not contain %d", v, val)
								}
							})
						}
					})
				})

				t.Run("Vs_Empty", func(t *testing.T) {
					testInterval := intervals.NewSingleton(v)
					otherEmpty := intervals.NewEmpty[T]()

					t.Run("Overlaps", func(t *testing.T) {
						req.True(testInterval.Overlaps(testInterval), "Singleton should overlap with self")
						req.False(testInterval.Overlaps(otherEmpty), "Singleton should not overlap with empty interval")
					})

					t.Run("Equals", func(t *testing.T) {
						req.True(testInterval.Equals(testInterval), "Singleton should equal itself")
						req.False(testInterval.Equals(otherEmpty), "Singleton should not equal empty interval")
					})

					t.Run("Union", func(t *testing.T) {
						req.False(testInterval.Union(testInterval).IsEmpty(), "Union with self should not be empty")
						req.True(testInterval.Union(testInterval).IsSingleton(), "Union with self should be singleton")
						// req.Nil(testInterval.Union(testInterval).FLOOP(), "Union with self should not have intervals")

						req.False(testInterval.Union(otherEmpty).IsEmpty(), "Union with Null should not be empty")
						req.True(testInterval.Union(otherEmpty).IsSingleton(), "Union with Null should be singleton")
						// req.Nil(testInterval.Union(otherEmpty).FLOOP(), "Union with Null should not have intervals")
					})

					t.Run("Intersection", func(t *testing.T) {
						req.False(testInterval.Intersection(testInterval).IsEmpty(), "Intersection with self should not be empty")
						req.True(testInterval.Intersection(testInterval).IsSingleton(), "Intersection with self should e singleton")
						// req.Nil(testInterval.Intersection(testInterval).FLOOP(), "Intersection with self should not have intervals")

						req.True(testInterval.Intersection(otherEmpty).IsEmpty(), "Intersection with Null should be empty")
						req.False(testInterval.Intersection(otherEmpty).IsSingleton(), "Intersection with Null should not be singleton")
						// req.Nil(testInterval.Intersection(otherEmpty).FLOOP(), "Intersection with Null should not have intervals")
					})

					t.Run("Difference", func(t *testing.T) {
						req.True(testInterval.Difference(testInterval).IsEmpty(), "Difference with self should be empty")
						req.False(testInterval.Difference(testInterval).IsSingleton(), "Difference with self should not be singleton")
						// req.Nil(testInterval.Difference(testInterval).FLOOP(), "Difference with self should not have intervals")

						req.False(testInterval.Difference(otherEmpty).IsEmpty(), "Difference with Null should not be empty")
						req.True(testInterval.Difference(otherEmpty).IsSingleton(), "Difference with Null should be singleton")
						// req.Nil(testInterval.Difference(otherEmpty).FLOOP(), "Difference with Null should not have intervals")
					})
				})

				t.Run("Vs_Single", func(t *testing.T) {

				})

				t.Run("Vs_Merged", func(t *testing.T) {

				})
			})
		}
	})
}
