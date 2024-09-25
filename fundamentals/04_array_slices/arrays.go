package main

import "fmt"

func main() {
	fmt.Println("testing in array and slices")
	arr := [5]int{1, 2, 3, 4, 5}
	products := SumNumbers(arr[:])
	fmt.Println(products)
}

// func Sum(numbers [5]int) int {
// 	return 0
// }

func SumNumbers(numbers []int) int {
	summation := 0

	for i := 0; i < len(numbers); i++ {
		summation += numbers[i]
	}
	return summation
}

// function adding slices
func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

// requirement to summall
func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)

	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
