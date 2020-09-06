package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 4

	fmt.Println("Sum => ", a+b)
	fmt.Println("Sub => ", a-b)
	fmt.Println("Division => ", a/b)
	fmt.Println("Multiply => ", a*b)
	fmt.Println("MOD => ", a%b)

	// bitwise
	fmt.Println("AND => ", a&b)
	fmt.Println("OR => ", a|b)
	fmt.Println("XOR => ", a^b)

	c := 3.0
	d := 2.0

	// another operations using math.
	fmt.Println("Greater => ", math.Max(float64(a), float64(b)))
	fmt.Println("Smaller => ", math.Min(c, d))
	fmt.Println("Power => ", math.Pow(c, d))

}
