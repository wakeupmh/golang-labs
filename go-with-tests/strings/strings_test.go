package strings

import "testing"

func TestAlphabeticOrder(t *testing.T) {
	t.Run("should a < b in alphabetic order", func(t *testing.T) {
		bigger := Compare("a", "b")
		expected := -1

		if bigger != expected {
			t.Errorf("expect '%d' but got '%d'", expected, bigger)
		}
	})

	t.Run("should a > b in alphabetic order", func(t *testing.T) {
		bigger := Compare("b", "a")
		expected := 1

		if bigger != expected {
			t.Errorf("expect '%d' but got '%d'", expected, bigger)
		}
	})

	t.Run("should a == b in alphabetic order", func(t *testing.T) {
		bigger := Compare("b", "b")
		expected := 0

		if bigger != expected {
			t.Errorf("expect '%d' but got '%d'", expected, bigger)
		}
	})
	
}