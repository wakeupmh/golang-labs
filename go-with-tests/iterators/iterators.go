package iterators

func Repeat(letter string, times int) (repetitions string) {
	for i := 0; i < times; i++ {
		repetitions += letter
	}

	return
}