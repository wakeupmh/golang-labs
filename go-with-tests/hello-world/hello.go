package main

import "fmt"

const prefixGrettings = "Hello, "
const prefixSpanishGreetings = "Hola, "
const prefixFrenchGreetings = "Bonjour, "

func prefixGreeting(language string) (prefix string) {
	switch language {
	case "french":
		prefix = prefixFrenchGreetings
	case "spanish":
		prefix = prefixSpanishGreetings
	default:
		prefix = prefixGrettings
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
        name = "World"
	}
	
	return prefixGreeting(language) + name
}

func main() {
	fmt.Println(Hello("Marcos", ""))
}