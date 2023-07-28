package utils

import (
	"encoding/json"
	"sort"

	"golang.org/x/exp/constraints"
)

// MergeMaps creates a new map and copies target and then source into it
func MergeMaps[K comparable, V any](target, source map[K]V) map[K]V {
	var result map[K]V = make(map[K]V, len(source)+len(target))
	for k, v := range target {
		result[k] = v
	}
	for k, v := range source {
		result[k] = v
	}
	return result
}

// Keys gets a slice of all the keys in a map
func Keys[K comparable, V any](self map[K]V) []K {
	var result []K = make([]K, len(self))
	id := 0
	for k := range self {
		result[id] = k
		id++
	}
	return result
}

// SortedKeys gets a slice of all the keys in a map, sorted
func SortedKeys[K constraints.Ordered, V any](self map[K]V) []K {
	var result []K = make([]K, len(self))
	id := 0
	for k := range self {
		result[id] = k
		id++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// Values gets a slice of all the values in a map
func Values[K comparable, V any](self map[K]V) []V {
	var result []V = make([]V, len(self))
	id := 0
	for _, v := range self {
		result[id] = v
		id++
	}
	return result
}

// MapToStruct converts a map to a struct by converting through json
func MapToStruct[T any](m map[string]any) (T, error) {
	var result T
	data, err := json.Marshal(m)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// StructToMap converts a struct to a map by converting through json
func StructToMap[T any](t T) (map[string]any, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	var out map[string]any
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
