package sum

func Sum(numbers [5]int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

