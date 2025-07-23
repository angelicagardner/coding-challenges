package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b     string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"111", "111", "1110"},
		{"110010", "10111", "1001001"},
		{"1111111111111111", "1", "10000000000000000"},
	}

	for _, tt := range tests {
		result := addBinary(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("addBinary(%q, %q) = %q; want %q", tt.a, tt.b, result, tt.expected)
		}
	}
}
