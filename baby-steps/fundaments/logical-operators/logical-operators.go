package main

import "fmt"

func shop(work1, work2 bool) (bool, bool, bool) {
	buy50TV := work1 && work2
	buy32TV := work1 != work2 // exclusive or
	buyIceCream := work1 || work2

	return buy50TV, buy32TV, buyIceCream
}

func main() {
	tv50, tv32, iceCream := shop(true, true)
	fmt.Println("50 TV: %t, 32 TV: %t, ice cream: %t",
		tv50, tv32, iceCream)
}
