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
		{
			name:           "Example 1",
			nums:           []int{3, 2, 2, 3},
			val:            3,
			expectedLength: 2,
			expectedNums:   []int{2, 2},
		},
		{
			name:           "Example 2",
			nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:            2,
			expectedLength: 5,
			expectedNums:   []int{0, 1, 4, 0, 3},
		},
		{
			name:           "All elements to remove",
			nums:           []int{7, 7, 7, 7},
			val:            7,
			expectedLength: 0,
			expectedNums:   []int{},
		},
		{
			name:           "No elements to remove",
			nums:           []int{1, 2, 3, 4},
			val:            5,
			expectedLength: 4,
			expectedNums:   []int{1, 2, 3, 4},
		},
		{
			name:           "Empty array",
			nums:           []int{},
			val:            1,
			expectedLength: 0,
			expectedNums:   []int{},
		},
		{
			name:           "Single element not to remove",
			nums:           []int{1},
			val:            2,
			expectedLength: 1,
			expectedNums:   []int{1},
		},
		{
			name:           "Single element to remove",
			nums:           []int{1},
			val:            1,
			expectedLength: 0,
			expectedNums:   []int{},
		},
		{
			name:           "Multiple occurrences scattered",
			nums:           []int{4, 5, 6, 4, 7, 4, 8},
			val:            4,
			expectedLength: 4,
			expectedNums:   []int{5, 6, 7, 8},
		},
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
