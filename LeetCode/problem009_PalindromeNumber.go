/*
Problem 9. Palindrome Number
*/

package leetcode

import (
	"math"
)

// Time complexity:  O(log10​(n))
// Space complexity: O(1)
func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	// Reversing the number
	for x != 0 {
		// Pop the last digit from x
		pop := x % 10
		x /= 10

		// Check for overflow (as reversing might cause overflow)
		if reversed > (math.MaxInt32-pop)/10 {
			return false
		}
		// Push the digit to the reversed number
		reversed = reversed*10 + pop
	}

	// Check if the original number is the same as the reversed number
	return original == reversed
}
