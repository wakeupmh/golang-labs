package main

import "fmt"

const prefixGrettings = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
        name = "World"
	}
	
	if language == "spanish" {
		return "Hola, " + name
	}

	return prefixGrettings + name
}

func main() {
	fmt.Println(Hello("Marcos", ""))
}