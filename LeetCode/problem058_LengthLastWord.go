package leetcode

// Time complexity: O(n)
// Space complexity: O(1)
func lengthOfLastWord(s string) int {
	// skip trailing spaces
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// count non-space chars
	length := 0
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}
