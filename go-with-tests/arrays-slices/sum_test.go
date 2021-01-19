package sum

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	verifySums := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, result)
		}
	}

	t.Run("should sum a collection with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expect := 6

		if sum != expect {
			t.Errorf("expected '%d' but got '%d', given: %d", expect, sum, numbers)
		}
	})

	t.Run("should sum all collection elements", func(t *testing.T) {
		sum := SumAll([]int{1,2}, []int{0,9})
		expect := []int{3, 9}

		verifySums(t, sum, expect)

	})

	t.Run("should sum the rest", func(t *testing.T) {
		sum := SumRest([]int{1,2}, []int{0,9})
		expect := []int{2, 9}

		verifySums(t, sum, expect)
	})
}