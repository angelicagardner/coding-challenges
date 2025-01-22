package leetcode

// Time complexity:  O(n)
// Space complexity: O(1)
func RemoveElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
