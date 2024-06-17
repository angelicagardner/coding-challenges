package leetcode

import (
	"reflect"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	var tests = []struct {
		name   string
		input  string
		output int
	}{
		{"case1", "III", 3},
		{"case2", "LVIII", 58},
		{"case3", "MCMXCIV", 1994},
		{"case4", "CMXCVIII", 998},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInteger(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("romanToInteger(%v) = %v; want %v", tt.input, result, tt.output)
			}
		})
	}
}
