package leetcode

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 1, 0},
	}

	for _, tt := range tests {
		result := searchInsert(tt.nums, tt.target)
		if result != tt.expected {
			t.Errorf("searchInsert(%v, %d) = %d; expected %d", tt.nums, tt.target, result, tt.expected)
		}
	}
}
