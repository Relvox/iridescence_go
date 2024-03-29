package asserts_test

import (
	"testing"

	"github.com/relvox/iridescence_go/asserts"
)

func TestSameElements(t *testing.T) {
	tests := []struct {
		name           string
		expected       []int
		actual         []int
		expectedErrors []string
	}{
		{
			name:           "Same elements, same order",
			expected:       []int{1, 2, 3, 4, 5},
			actual:         []int{1, 2, 3, 4, 5},
			expectedErrors: []string{},
		},
		{
			name:           "Same elements, different order",
			expected:       []int{1, 2, 3, 4, 5},
			actual:         []int{5, 4, 3, 2, 1},
			expectedErrors: []string{},
		},
		{
			name:     "Not same sizes",
			expected: []int{1, 2, 3, 4, 4},
			actual:   []int{4, 3, 2, 1},
			expectedErrors: []string{
				"Not all expected items found in actual:\n\tmap[4:1]\n",
			},
		},
		{
			name:     "Different elements",
			expected: []int{1, 2, 3, 4, 5, 5},
			actual:   []int{1, 2, 3, 4, 5, 6},
			expectedErrors: []string{
				"Extra items found in actual that were not expected:\n\tmap[6:1]\n",
				"Not all expected items found in actual:\n\tmap[5:1]\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := asserts.NewMockT()
			asserts.SameElements(mockT, tt.expected, tt.actual)
			mockT.Assert(t, tt.expectedErrors...)
		})
	}
}

func TestSameKeyValues(t *testing.T) {
	tests := []struct {
		name           string
		expected       map[string]int
		actual         map[string]int
		expectedErrors []string
	}{
		{
			name:           "Same key-value pairs, same order",
			expected:       map[string]int{"a": 1, "b": 2, "c": 3},
			actual:         map[string]int{"a": 1, "b": 2, "c": 3},
			expectedErrors: []string{},
		},
		{
			name:           "Same key-value pairs, different order",
			expected:       map[string]int{"a": 1, "b": 2, "c": 3},
			actual:         map[string]int{"c": 3, "b": 2, "a": 1},
			expectedErrors: []string{},
		},
		{
			name:     "Different values",
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
			actual:   map[string]int{"a": 1, "b": 2, "c": 4},
			expectedErrors: []string{
				"Extra items found in actual that were not expected:\n\tmap[c:4]\n",
				"Not all expected items found in actual:\n\tmap[c:3]\n",
			},
		},
		{
			name:     "Missing keys",
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
			actual:   map[string]int{"a": 1, "b": 2},
			expectedErrors: []string{
				"Not all expected items found in actual:\n\tmap[c:3]\n",
			},
		},
		{
			name:     "Extra keys",
			expected: map[string]int{"a": 1, "b": 2},
			actual:   map[string]int{"a": 1, "b": 2, "c": 3},
			expectedErrors: []string{
				"Extra items found in actual that were not expected:\n\tmap[c:3]\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := asserts.NewMockT()
			asserts.SameKeyValues(mockT, tt.expected, tt.actual)
			mockT.Assert(t, tt.expectedErrors...)
		})
	}
}
