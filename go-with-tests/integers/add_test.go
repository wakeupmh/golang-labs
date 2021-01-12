package integers 

import "testing"

func TestAdd (t *testing.T) {
	t.Run("should sum args", func(t *testing.T) {
		sum := Add(2, 2)
		expect := 4

		if sum != expect {
			t.Errorf("expect to be: '%d' but got: '%d'", expect, sum)
		}
	})
}