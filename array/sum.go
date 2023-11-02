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

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var tailSum []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailSum = append(tailSum, 0)
		} else {
			tail := numbers[1:]
			tailSum = append(tailSum, Sum(tail))
		}
	}
	return tailSum
}
