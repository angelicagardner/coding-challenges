package leetcode

// Time complexity: O(n)
// Space complexity: O(n)
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	out := make([]int, n+1)
	out[0] = 1
	return out
}
