package sum

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(numsToSum ...[]int) (result []int) {
	result = make([]int, len(numsToSum))
	
	for i, numbers := range numsToSum {
		result[i] = Sum(numbers)
	}

	return
}

