package main

import "fmt"

func main() {
	i := 1
	var p *int = nil
	p = &i // get address of variable i
	*p++   // desreference (getting value)
	i++

	fmt.Println(p, *p, i, &i)
}
