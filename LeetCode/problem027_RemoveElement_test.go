package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		val            int
		expectedLength int
		expectedNums   []int
	}{
		{"Example 1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"Example 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3}},
		{"All elements to remove", []int{7, 7, 7, 7}, 7, 0, []int{}},
		{"No elements to remove", []int{1, 2, 3, 4}, 5, 4, []int{1, 2, 3, 4}},
		{"Empty array", []int{}, 1, 0, []int{}},
		{"Single element not to remove", []int{1}, 2, 1, []int{1}},
		{"Single element to remove", []int{1}, 1, 0, []int{}},
		{"Multiple occurrences scattered", []int{4, 5, 6, 4, 7, 4, 8}, 4, 4, []int{5, 6, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalNums := make([]int, len(tt.nums))
			copy(originalNums, tt.nums)

			length := RemoveElement(tt.nums, tt.val)

			if length != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, length)
			}
		})
	}
}
