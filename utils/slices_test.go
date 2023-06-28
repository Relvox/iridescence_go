package utils_test

import (
	"fmt"
	"testing"

	"github.com/relvox/iridescence/utils"
	"github.com/stretchr/testify/assert"
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
			assert.True(t, utils.Equal(tt.expected, actual))
		})
	}
}
