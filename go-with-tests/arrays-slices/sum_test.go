package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should sum five numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expect := 15

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given: %d", expect, sum, numbers)
		}
	})
}