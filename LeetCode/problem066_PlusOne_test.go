package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{0}, []int{1}},
		{[]int{9}, []int{1, 0}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{4, 9, 9}, []int{5, 0, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1, 0, 0}, []int{1, 0, 1}},
		{nil, []int{1}},
		{[]int{}, []int{1}},
	}

	for _, tt := range tests {
		result := plusOne(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("plusOne(%v) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
