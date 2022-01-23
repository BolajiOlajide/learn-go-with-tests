package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) (tails []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			tails = append(tails, 0)
		} else {
			tails = append(tails, Sum(slice[1:]))
		}
	}

	return tails
}
