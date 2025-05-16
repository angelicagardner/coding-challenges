package leetcode

import (
	"testing"
)

// The LeetCode problem defines a “word” to be any continuous run of non-space characters,
// so here every non‐space character is considered as part of a word, incl. dots, hyphens, etc.
func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"", 0},
		{"     ", 0},
		{"word", 4},
		{"a", 1},
		{"a ", 1},
		{" a", 1},
		{"Hello, world!", 6},
		{"punctuation-test.", 17},
		{"...trailing.dots...", 19},
		{"你好 世界", len("世界")},
		{"emoji 🚀🚀 🚀", len("🚀")},
	}

	for _, tt := range tests {
		result := lengthOfLastWord(tt.input)
		if result != tt.expected {
			t.Errorf("lengthOfLastWord(%q) = %d; expected %d", tt.input, result, tt.expected)
		}
	}
}
