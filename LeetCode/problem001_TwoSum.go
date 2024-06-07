/*
Problem 1. Two Sum
*/

package leetcode

// Time complexity:  O(n^2)
// Space complexity: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Time complexity:  O(n)
// Space complexity: O(n)
func TwoSumHashMap(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		complementIndex, found := hashMap[complement]
		if found {
			return []int{complementIndex, i}
		}
		hashMap[num] = i
	}
	return nil
}
