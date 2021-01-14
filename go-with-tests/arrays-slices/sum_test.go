package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should sum a collection with five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expect := 15

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given: %d", expect, sum, numbers)
		}
	})

	t.Run("should sum a collection with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expect := 6

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given: %d", expect, sum, numbers)
		}
	})

}