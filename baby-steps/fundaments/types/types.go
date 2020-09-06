package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal integer is", reflect.TypeOf(32000))

	//uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("Byte is", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Max integer value is", i1)

	var i2 rune = 'a' // represents a mapped from unicode's table (int32)
	fmt.Println("Rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.999
	fmt.Println("The x's type is ", reflect.TypeOf(x))
	fmt.Println("The literals type is ", reflect.TypeOf(49.99)) //float64 by default

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Hello man"
	fmt.Println(s1)
	fmt.Println("The lenght of string is ", len(s1))

	s2 := `My
	name
	is
	nothing`
	fmt.Println(s2)
	fmt.Println("The lenght of string2 is ", len(s2))
}
