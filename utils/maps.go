package utils

import (
	"encoding/json"
)

func MergeMaps[K comparable, V any](target, source map[K]V) map[K]V {
	var result map[K]V = make(map[K]V)
	for k, v := range target {
		result[k] = v
	}
	for k, v := range source {
		result[k] = v
	}
	return result
}

func Values[K comparable, V any](self map[K]V) []V {
	var result []V = make([]V, len(self))
	id := 0
	for _, v := range self {
		result[id] = v
		id++
	}
	return result
}

func MapToStruct[T any](m map[string]any) (T, error) {
	var result T
	data, err := json.Marshal(m)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func StructToMap[T any](t T) (map[string]interface{}, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
