package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Unit is an empty struct
type Unit struct{}

// U is the unit
var U = Unit{}

// Set of T is a map from T to Unit
type Set[T comparable] map[T]Unit

func (s Set[T]) Add(elem T) {
	s[elem] = U
}

func NewSet[T comparable](elements ...T) Set[T] {
	result := make(Set[T])
	for _, e := range elements {
		result[e] = U
	}
	return result
}

func (s Set[T]) Union(elements ...T) Set[T] {
	var result Set[T] = make(Set[T])
	for k := range s {
		result[k] = U
	}
	for _, e := range elements {
		result[e] = U
	}
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if _, ok := other[k]; !ok {
			result[k] = U
		}
	}
	return result
}

func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

func (s Set[T]) String() string {
	var result []string
	for k := range s {
		result = append(result, fmt.Sprint(k))
	}
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}
func (s Set[T]) MarshalJSON() ([]byte, error) {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}

	return json.Marshal(slice)
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*s = make(Set[T])
	for _, elem := range slice {
		s.Add(elem)
	}

	return nil
}

func (s Set[T]) MarshalYAML() (interface{}, error) {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}
	return slice, nil
}

func (s *Set[T]) UnmarshalYAML(value *yaml.Node) error {
	var slice []T
	for _, v := range value.Content {
		var t T
		err := v.Decode(&t)
		if err != nil {
			return err
		}
		slice = append(slice, t)
	}

	*s = make(Set[T])
	for _, elem := range slice {
		s.Add(elem)
	}

	return nil
}
