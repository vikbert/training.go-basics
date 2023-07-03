package main

import "fmt"

func main() {
	fmt.Println("If-Statements:")
	numbers := [3]int{9, 1, 3}
	min, max := getMinAndMax(numbers)
	fmt.Printf("%v: Min: %d, Max: %d", numbers, min, max)
}

func getMinAndMax(numbers [3]int) (min, max int) {
	min = numbers[0]
	max = numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return min, max
}
