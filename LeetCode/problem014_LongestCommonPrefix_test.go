package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		name   string
		input  []string
		output string
	}{
		{
			name:   "Example 1",
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			name:   "Example 2",
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			name:   "All strings identical",
			input:  []string{"test", "test", "test"},
			output: "test",
		},
		{
			name:   "Single string in array",
			input:  []string{"single"},
			output: "single",
		},
		{
			name:   "Empty array",
			input:  []string{},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			if result != tt.output {
				t.Errorf("longestCommonPrefix(%v) = %q; expected %q", tt.input, result, tt.output)
			}
		})
	}
}
