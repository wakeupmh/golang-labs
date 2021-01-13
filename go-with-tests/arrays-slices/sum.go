package sum

func Sum(numbers [5]int) (result int) {
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}

	return
}

