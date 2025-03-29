package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"sadbutsad", "but", 3},
		{"sadbutsad", "notfound", -1},
		{"abc", "", 0},
		{"", "", 0},
		{"", "a", -1},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}

	for _, tc := range tests {
		result := strStr(tc.haystack, tc.needle)
		if result != tc.expected {
			t.Errorf("strStr(%q, %q) = %d; expected %d", tc.haystack, tc.needle, result, tc.expected)
		}
	}
}
