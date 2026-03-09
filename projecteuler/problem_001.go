package projecteuler

// Time complexity:     O(1)
// Space complexity:    O(1)
func Solution(limit int) int {
	n := limit - 1

	sumMultiples := func(k int) int {
		p := n / k
		return k * (p * (p + 1)) / 2
	}

	return sumMultiples(3) + sumMultiples(5) - sumMultiples(15)
}
