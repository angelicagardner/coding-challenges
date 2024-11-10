/*
Problem 14. Longest Common Prefix
*/

package leetcode

// Time complexity:  O(n*m)
// Space complexity: O(1)
func longestCommonPrefix(strs []string) string {
	// Edge case: if the input array is empty, return an empty string.
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {
		j := 0
		for j < len(prefix) && j < len(s) && prefix[j] == s[j] {
			j++
		}
		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
