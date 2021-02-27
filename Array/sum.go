package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, value := range numbersToSum {
		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(value[1:]))
		}

	}
	return sums
}
