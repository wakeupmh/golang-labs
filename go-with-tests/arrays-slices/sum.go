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

func SumRest(numsToSum ...[]int) (result []int) {
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			rest := numbers[1:]
			result = append(result, Sum(rest))
		}
	}

	return 
}
