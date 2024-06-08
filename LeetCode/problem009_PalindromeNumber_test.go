package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{1221, true},
		{12321, true},
		{1234567899, false},
		{1, true},
		{-1, false},
		{1000021, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d", test.input), func(t *testing.T) {
			result := isPalindrome(test.input)
			if result != test.expected {
				t.Errorf("For input %d, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
