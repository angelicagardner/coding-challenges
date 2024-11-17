package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		name        string
		inputString string
		expected    bool
	}{
		{"Case 1", "()", true},
		{"Case 2", "()[]{}", true},
		{"Case 3", "(]", false},
		{"Case 4", "{[]}", true},
		{"Correct and incorrect example", "{[}", false},
		{"Empty string", "", true},
		{"Only opening parentheses", "(", false},
		{"Only closing parentheses", ")", false},
		{"Longer correct example", "({[]})[]{}", true},
		{"Longer incorrect example", "({[]})[]{}{", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.inputString)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v; expected %v", tt.inputString, result, tt.expected)
			}
		})
	}
}
