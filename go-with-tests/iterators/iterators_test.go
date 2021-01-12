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

func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}