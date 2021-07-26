package maps

import "testing"

func assertStrings(t testing.TB, got, expect string) {
	t.Helper()

	if got != expect {
		t.Errorf("got %q want %q", got, expect)
	}
}

func assertError(t testing.TB, got, expect error) {
	t.Helper()

	if got != expect {
		t.Errorf("got error %q want %q", got, expect)
	}
}

func TestSearches(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("should return a known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expect := "this is a test"

		assertStrings(t, got, expect)
	})

	t.Run("should return an error for a unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}
