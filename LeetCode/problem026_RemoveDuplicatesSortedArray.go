package leetcode

// Time complexity:  O(n)
// Space complexity: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueElements := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[uniqueElements] = nums[i]
			uniqueElements++
		}
	}

	return uniqueElements
}
