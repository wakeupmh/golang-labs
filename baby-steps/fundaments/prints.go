package main

import "fmt"

func main() {
	fmt.Print("Same ")
	fmt.Print("line.")

	fmt.Println("\nAnother")
	fmt.Println("line.")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("The x's valor is " + xs)
	fmt.Println("The x's valor is", x)

	fmt.Printf("The x's valor is %.2f", x)

	a := 1
	b := 1.999999
	c := false
	d := "hey"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
