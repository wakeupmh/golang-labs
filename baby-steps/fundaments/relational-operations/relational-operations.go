package main

import "fmt"

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	type People struct {
		Name string
	}

	p1 := People{"Dylan"}
	p2 := People{"Inaê"}
	fmt.Println("Pessoas:", p1 == p2) // values are compared not memory referation

}
