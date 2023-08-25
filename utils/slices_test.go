package utils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"

	"github.com/relvox/iridescence_go/utils"
)

func Test_Intersect(t *testing.T) {
	tests := []struct {
		list1    []string
		list2    []string
		expected []string
	}{
		{
			list1:    []string{},
			list2:    []string{},
			expected: []string{},
		},
		{
			list1:    []string{"a", "b"},
			list2:    []string{"b", "c"},
			expected: []string{"b"},
		},
		{
			list1:    []string{"a", "b", "c"},
			list2:    []string{"c", "b", "d"},
			expected: []string{"c", "b"},
		},
	}

	for ti, tt := range tests {
		t.Run(fmt.Sprint(ti), func(t *testing.T) {
			actual := utils.Intersect(tt.list1, tt.list2)
			assert.Equal(t, len(tt.expected), len(actual))
			assert.True(t, slices.Equal(tt.expected, actual))
		})
	}
}

func TestCrop(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		index  int
		output []int
	}{
		{
			name:   "Remove first element",
			input:  []int{1, 2, 3, 4, 5},
			index:  0,
			output: []int{2, 3, 4, 5},
		},
		{
			name:   "Remove middle element",
			input:  []int{1, 2, 3, 4, 5},
			index:  2,
			output: []int{1, 2, 4, 5},
		},
		{
			name:   "Remove last element",
			input:  []int{1, 2, 3, 4, 5},
			index:  4,
			output: []int{1, 2, 3, 4},
		},
		{
			name:   "Index out of range (negative)",
			input:  []int{1, 2, 3, 4, 5},
			index:  -1,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Index out of range (too large)",
			input:  []int{1, 2, 3, 4, 5},
			index:  10,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Empty slice",
			input:  []int{},
			index:  0,
			output: []int{},
		},
		{
			name:   "Single element slice, valid index",
			input:  []int{1},
			index:  0,
			output: []int{},
		},
		{
			name:   "Single element slice, invalid index",
			input:  []int{1},
			index:  2,
			output: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.Crop(tt.input, tt.index)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("expected %v, got %v", tt.output, got)
			}
		})
	}
}
