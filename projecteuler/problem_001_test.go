package projecteuler

import "testing"

func TestSolution(t *testing.T) {
	t.Run("limit 10", func(t *testing.T) {
		limit := 10
		expected := 23
		got := Solution(limit)

		if got != expected {
			t.Errorf("Solution(%d) = %d; want %d", limit, got, expected)
		}
	})

	t.Run("limit 1000", func(t *testing.T) {
		limit := 1000
		expected := 233168
		got := Solution(limit)

		if got != expected {
			t.Errorf("Solution(%d) = %d; want %d", limit, got, expected)
		}
	})
}
