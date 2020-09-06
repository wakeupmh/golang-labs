package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	note := 6.9
	finalNote := int(note)
	fmt.Println(finalNote)

	// beware......
	fmt.Println("Test " + string(97)) //asc table

	//int to string conversion
	fmt.Println("Test" + strconv.Itoa(123))

	//string to int conversion
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("false")
	if !b {
		fmt.Println(b)
	}
}
