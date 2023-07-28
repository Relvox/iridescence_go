package utils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	assutil "github.com/relvox/iridescence_go/assert"
	"github.com/relvox/iridescence_go/prototyping"
	"github.com/relvox/iridescence_go/utils"
)

func gen_map[K comparable, V any](size int, genKey func() K, genVal func() V) map[K]V {
	result := make(map[K]V)
	for i := 0; i < size; i++ {
		result[genKey()] = genVal()
	}
	return result
}

func split_map[K comparable, V any](toSplit map[K]V) (map[K]V, map[K]V) {
	resLeft, resRight := make(map[K]V), make(map[K]V)
	left := true
	for k, v := range toSplit {
		if left {
			resLeft[k] = v
		} else {
			resRight[k] = v
		}
		left = !left
	}
	return resLeft, resRight
}

func gen_entity_and_map(magnitude int) (*prototyping.Entity, map[string]any) {
	mapProps := make(map[string]any, magnitude)
	mapStats := make(map[string]any, magnitude)
	tags := make([]any, magnitude)
	resultEnt := prototyping.NewEntity(magnitude, fmt.Sprint(magnitude))
	for i := 0; i < magnitude; i++ {
		resultEnt = resultEnt.
			WithProperty(fmt.Sprint("p", i), fmt.Sprint(i)).
			WithStat(fmt.Sprint("s", i), i).
			WithTag(fmt.Sprint("t", i))
		mapProps[fmt.Sprint("p", i)] = fmt.Sprint(i)
		mapStats[fmt.Sprint("s", i)] = float64(i)
		tags[i] = fmt.Sprint("t", i)
	}
	resultMap := map[string]any{
		"Id":         float64(magnitude),
		"Name":       fmt.Sprint(magnitude),
		"Properties": mapProps,
		"Stats":      mapStats,
		"Tags":       tags,
	}
	return resultEnt, resultMap
}

func Test_MergeMaps(t *testing.T) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }

	for k := 10; k <= 1000001; k *= 10 {
		t.Run(fmt.Sprintf("%d keys", k), func(t *testing.T) {
			expected := gen_map(k, intGen, intGen)
			m1, m2 := split_map(expected)
			actual := utils.MergeMaps(m1, m2)
			assert.InDeltaMapValues(t, expected, actual, 0, "maps should be identical")
		})
	}
}

func Benchmark_MergeMaps(b *testing.B) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }

	for k := 10; k <= 1000001; k *= 10 {
		b.Run(fmt.Sprintf("%d keys", k), func(b *testing.B) {
			b.StopTimer()
			expected := gen_map(k, intGen, intGen)
			m1, m2 := split_map(expected)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				utils.MergeMaps(m1, m2)
			}
		})
	}
}

func Test_Keys(t *testing.T) {
	for k := 10; k <= 1000001; k *= 10 {
		t.Run(fmt.Sprintf("%d keys", k), func(t *testing.T) {
			__Z := 0
			intGen := func() int { __Z++; return __Z }
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = i*2 + 1
			}
			actual := utils.Keys(originalMap)
			assutil.SameElements(t, expected, actual)
		})
	}
}

func Benchmark_Keys(b *testing.B) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }
	for k := 10; k <= 1000001; k *= 10 {
		b.Run(fmt.Sprintf("%d keys", k), func(b *testing.B) {
			b.StopTimer()
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = i*2 + 1
			}
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				utils.Keys(originalMap)
			}
		})
	}
}

func Test_Values(t *testing.T) {
	for k := 10; k <= 1000001; k *= 10 {
		t.Run(fmt.Sprintf("%d keys", k), func(t *testing.T) {
			__Z := 0
			intGen := func() int { __Z++; return __Z }
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = (i + 1) * 2
			}
			actual := utils.Values(originalMap)
			assutil.SameElements(t, expected, actual)
		})
	}
}

func Benchmark_Values(b *testing.B) {
	__Z := 0
	intGen := func() int { __Z++; return __Z }
	for k := 10; k <= 1000001; k *= 10 {
		b.Run(fmt.Sprintf("%d keys", k), func(b *testing.B) {
			b.StopTimer()
			originalMap := gen_map(k, intGen, intGen)
			expected := make([]int, k)
			for i := 0; i < k; i++ {
				expected[i] = (i + 1) * 2
			}
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				utils.Values(originalMap)
			}
		})
	}
}

func Test_MapToStruct(t *testing.T) {
	t.Run("small map", func(t *testing.T) {
		m := map[string]interface{}{
			"Id":   1,
			"Name": "Alice",
			"Properties": map[string]string{
				"prop1": "value1",
				"prop2": "value2",
			},
			"Stats": map[string]int{
				"stat1": 10,
				"stat2": 20,
			},
			"Tags": []string{
				"tag1",
				"tag2",
			},
		}

		e, err := utils.MapToStruct[prototyping.Entity](m)

		assert.NoError(t, err)
		assert.Equal(t, 1, e.Id)
		assert.Equal(t, "Alice", e.Name)
		assert.True(t, reflect.DeepEqual(map[string]string{"prop1": "value1", "prop2": "value2"}, e.Properties))
		assert.True(t, reflect.DeepEqual(map[string]int{"stat1": 10, "stat2": 20}, e.Stats))
		assert.True(t, reflect.DeepEqual(utils.NewSet("tag1", "tag2"), e.Tags))
	})

	for k := 100; k < 1000001; k *= 10 {
		t.Run(fmt.Sprintf("big map %d", k), func(t *testing.T) {
			expected, m := gen_entity_and_map(k)
			actual, err := utils.MapToStruct[prototyping.Entity](m)

			assert.NoError(t, err)
			assert.Equal(t, expected.Id, actual.Id)
			assert.Equal(t, expected.Name, actual.Name)
			assert.True(t, reflect.DeepEqual(expected.Properties, actual.Properties))
			assert.True(t, reflect.DeepEqual(expected.Stats, actual.Stats))
			assert.True(t, reflect.DeepEqual(expected.Tags, actual.Tags))
		})
	}
}

func Benchmark_MapToStruct(b *testing.B) {
	b.Run("small map", func(b *testing.B) {
		m := map[string]interface{}{
			"Id":   1,
			"Name": "Alice",
			"Properties": map[string]string{
				"prop1": "value1",
				"prop2": "value2",
			},
			"Stats": map[string]int{
				"stat1": 10,
				"stat2": 20,
			},
			"Tags": []string{
				"tag1",
				"tag2",
			},
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			utils.MapToStruct[prototyping.Entity](m)
		}
	})
	for k := 100; k < 1000001; k *= 10 {
		b.Run(fmt.Sprintf("big map %d", k), func(b *testing.B) {
			b.StopTimer()
			_, m := gen_entity_and_map(k)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				utils.MapToStruct[prototyping.Entity](m)
			}
		})
	}
}

func Test_StructToMap(t *testing.T) {
	t.Run("small entity", func(t *testing.T) {
		e := prototyping.NewEntity(1, "Alice").
			WithProperty("prop1", "value1").
			WithProperty("prop2", "value2").
			WithStat("stat1", 10).
			WithStat("stat2", 20).
			WithTag("tag1").
			WithTag("tag2")

		actual, err := utils.StructToMap(e)
		assert.NoError(t, err)
		assert.Equal(t, 1.0, actual["Id"])
		assert.Equal(t, "Alice", actual["Name"])
		assert.True(t, reflect.DeepEqual(map[string]any{"prop1": "value1", "prop2": "value2"}, actual["Properties"]))
		assert.True(t, reflect.DeepEqual(map[string]any{"stat1": 10.0, "stat2": 20.0}, actual["Stats"]))
		assutil.SameElements(t, []any{"tag1", "tag2"}, actual["Tags"].([]any))
	})
	for k := 100; k < 1000001; k *= 10 {
		t.Run(fmt.Sprintf("big entity %d", k), func(t *testing.T) {
			e, expected := gen_entity_and_map(k)
			actual, err := utils.StructToMap(e)
			assert.NoError(t, err)
			assert.Equal(t, expected["Id"], actual["Id"])
			assert.Equal(t, expected["Name"], actual["Name"])
			assert.True(t, reflect.DeepEqual(expected["Properties"], actual["Properties"]))
			assert.True(t, reflect.DeepEqual(expected["Stats"], actual["Stats"]))
			assutil.SameElements(t, expected["Tags"].([]any), actual["Tags"].([]any))
		})
	}
}

func Benchmark_StructToMap(b *testing.B) {
	b.Run("small entity", func(b *testing.B) {
		e := prototyping.NewEntity(1, "Alice").
			WithProperty("prop1", "value1").
			WithProperty("prop2", "value2").
			WithStat("stat1", 10).
			WithStat("stat2", 20).
			WithTag("tag1").
			WithTag("tag2")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			utils.StructToMap(e)
		}
	})
	for k := 100; k < 1000001; k *= 10 {
		b.Run(fmt.Sprintf("big map %d", k), func(b *testing.B) {
			b.StopTimer()
			e, _ := gen_entity_and_map(k)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				utils.StructToMap(e)
			}
		})
	}
}
