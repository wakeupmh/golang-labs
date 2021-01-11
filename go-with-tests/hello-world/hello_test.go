package main

import "testing"

func TestHello(t *testing.T) {
	name := "Marcos"
	result := Hello(name)
	expect := "Hello, " + name 

	if(result != expect) {
		t.Errorf("result: '%s', expect: '%s'", result, expect)
	}
}