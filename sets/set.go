package sets

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

// Unit is an empty struct
type Unit = struct{}

// U is the unit
var U = Unit{}

// Set of T is a map from T to Unit
type Set[T comparable] map[T]Unit

// NewSet creates a new set with initial elements
func NewSet[T comparable](elements ...T) Set[T] {
	result := make(Set[T])
	for _, e := range elements {
		result[e] = U
	}
	return result
}

// Add an element to the set
func (s Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = U
	}
}

func (s Set[T]) SetUnion(other Set[T]) Set[T] {
	result := make(Set[T], max(len(s), len(other)))
	for k := range s {
		result[k] = U
	}
	for k := range other {
		result[k] = U
	}
	return result
}

func (s Set[T]) Union(elements ...T) Set[T] {
	result := make(Set[T], max(len(s), len(elements)))
	for k := range s {
		result[k] = U
	}
	for _, e := range elements {
		result[e] = U
	}
	return result
}

func (s Set[T]) SetIntersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	for k := range s {
		if _, ok := other[k]; ok {
			result[k] = U
		}
	}
	return result
}

func (s Set[T]) Intersection(elements ...T) Set[T] {
	result := NewSet[T]()
	for _, k := range elements {
		if _, ok := s[k]; ok {
			result[k] = U
		}
	}
	return result
}

func (s Set[T]) SetDifference(other Set[T]) Set[T] {
	result := maps.Clone(s)
	for k := range other {
		delete(result, k)
	}
	return result
}

func (s Set[T]) Difference(elements ...T) Set[T] {
	result := maps.Clone(s)
	for _, k := range elements {
		delete(result, k)
	}
	return result
}

func (s Set[T]) ThreeWay(other Set[T]) [3]Set[T] {
	blr := [3]Set[T]{NewSet[T](), NewSet[T](), NewSet[T]()}
	for k := range s {
		if other.Contains(k) {
			blr[0].Add(k)
		} else {
			blr[1].Add(k)
		}
	}

	for ok := range other {
		if s.Contains(ok) {
			continue
		}
		blr[2].Add(ok)
	}
	return blr
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
