package main

import "fmt"

func main() {
	x := 1
	y := 2

	//Go has only postfix

	x++
	y--

	fmt.Println(x == y)
}
