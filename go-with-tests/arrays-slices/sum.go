package sum

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

