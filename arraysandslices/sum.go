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
		for idx, number := range slice {
			if idx != 0 {
				tails = append(tails, number)
			}
		}
	}

	return tails
}
