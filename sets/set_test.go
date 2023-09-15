package sets_test

import (
	"testing"

	"github.com/relvox/iridescence_go/assert"
)

const (
	mil  = 1_000_000
	semi = 500_000
)

func TestUnion(t *testing.T) {
	set1 := generateSet(mil, 0)
	set2 := generateSet(mil, semi)
	elements2 := generateElements(mil, semi)
	expected := generateSet(mil+semi, 0)
	t.Run("Set", func(t *testing.T) {
		actual := set1.SetUnion(set2)
		assert.SameKeyValues(t, expected, actual)
	})
	t.Run("Elements", func(t *testing.T) {
		actual := set1.Union(elements2...)
		assert.SameKeyValues(t, expected, actual)
	})
}

func TestIntersection(t *testing.T) {
	set1 := generateSet(mil, 0)
	set2 := generateSet(mil, semi)
	elements2 := generateElements(mil, semi)
	expected := generateSet(semi, semi)
	t.Run("Set", func(t *testing.T) {
		actual := set1.SetIntersection(set2)
		assert.SameKeyValues(t, expected, actual)
	})
	t.Run("Elements", func(t *testing.T) {
		actual := set1.Intersection(elements2...)
		assert.SameKeyValues(t, expected, actual)
	})
}

func TestDifference(t *testing.T) {
	set1 := generateSet(mil, 0)
	set2 := generateSet(mil, semi)
	elements2 := generateElements(mil, semi)
	expected := generateSet(semi, 0)
	t.Run("Set", func(t *testing.T) {
		actual := set1.SetDifference(set2)
		assert.SameKeyValues(t, expected, actual)
	})
	t.Run("Elements", func(t *testing.T) {
		actual := set1.Difference(elements2...)
		assert.SameKeyValues(t, expected, actual)
	})
}
