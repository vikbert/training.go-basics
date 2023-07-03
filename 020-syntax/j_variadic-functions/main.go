package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions:")

	fmt.Println("1 2 3")
	fmt.Println("sumWithDuplication: ", sumWithDuplication(1, 2, 3))
}

func sumWithDuplication(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		value = value * 2
		total += value
	}

	return total
}
