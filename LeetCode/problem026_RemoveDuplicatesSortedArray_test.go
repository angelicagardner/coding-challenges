package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
		output   int
	}{
		{"Short array", []int{1, 1, 2}, []int{1, 2}, 2},
		{"Longer array", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		{"Empty array", []int{}, []int{}, 0},
		{"Only duplicates", []int{1, 1, 1, 1}, []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveDuplicates(tt.input)
			if result != tt.output {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.input, result, tt.output)
			}
		})
	}
}
