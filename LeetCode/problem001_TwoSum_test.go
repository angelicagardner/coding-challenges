package leetcode

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name   string
	nums   []int
	target int
	output []int
}{
	{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"case2", []int{2, 1, 5, 3}, 4, []int{1, 3}},
	{"case3", []int{3, 2, 4}, 6, []int{1, 2}},
	{"case4", []int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSumBruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSumBruteForce(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("TwoSumBruteForce(%v, %v) = %v; want %v", tt.nums, tt.target, result, tt.output)
			}
		})
	}
}

func TestTwoSumHashMap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSumHashMap(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("TwoSumHashMap(%v, %v) = %v; want %v", tt.nums, tt.target, result, tt.output)
			}
		})
	}
}
