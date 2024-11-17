/*
Problem 20. Valid Parentheses
*/

package leetcode

// Time complexity:  O(n)
// Space complexity: O(n)
func isValid(s string) bool {
	stack := []rune{}

	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != bracketMap[char] {
				return false
			}
		}
	}

	return len(stack) == 0
}
