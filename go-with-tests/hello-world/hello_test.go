package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, result, expect string) {
		t.Helper()

		if(result != expect) {
			t.Errorf("result: '%s', expect: '%s'", result, expect)
		}
	}

	t.Run("should Say hello to Marcos", func(t *testing.T) { 
		result := Hello("Marcos", "")
		expect := "Hello, Marcos"

		verifyMessage(t, result, expect)
	})

	t.Run("should return world as a default response", func(t *testing.T) { 
		result := Hello("", "")
		expect := "Hello, World"

		verifyMessage(t, result, expect)
	})

	t.Run("should return Hello in spanish", func(t *testing.T) { 
		result := Hello("Marcos", "spanish")
		expect := "Hola, Marcos"

		verifyMessage(t, result, expect)
	})
	
	t.Run("should return Hello in french", func(t *testing.T) { 
		result := Hello("Marcos", "french")
		expect := "Bonjour, Marcos"

		verifyMessage(t, result, expect)
	})
}