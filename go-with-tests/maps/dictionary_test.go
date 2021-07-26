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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestSearch(t *testing.T) {
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

func TestAdd(t *testing.T) {
	t.Run("should add an item correctly", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add("test", "this is a test")

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("should return an existing word error", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
