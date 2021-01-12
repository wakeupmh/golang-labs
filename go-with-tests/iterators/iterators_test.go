package iterators

import "testing"

func TestIterator(t *testing.T) {
	t.Run("should return letter 'a' 5 times", func(t *testing.T) {
		repetitions := Repeat("a", 5)
		expect := "aaaaa"

		if repetitions != expect {
			t.Errorf("expected to be %s, but got %s", expect, repetitions)
		}
	})
}