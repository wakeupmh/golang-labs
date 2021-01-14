package sum

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(numsToSum ...[]int) (result []int) {
	for _, numbers := range numsToSum {
		result = append(result, Sum(numbers))
	}

	return
}

