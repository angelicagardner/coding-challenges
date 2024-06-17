/*
13. Roman to Integer
*/

package leetcode

// Time complexity:  O(n)
// Space complexity: O(1)
func romanToInteger(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	total := 0

	for i := 0; i < n; i++ {
		if i+1 < n && roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else {
			total += roman[s[i]]
		}
	}

	return total
}
