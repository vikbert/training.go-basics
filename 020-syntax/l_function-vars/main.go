package main

import "fmt"

type operatorFunc func(int, int) int

func main() {
	fmt.Println("1st class Functions:")

	var sum operatorFunc = func(a int, b int) int { return a + b }

	var minus operatorFunc = func(a int, b int) int { return a - b }

	var multiply = func(a int, b int, add operatorFunc) int {
		result := 0
		for i := 0; i < b; i++ {
			result = add(result, a)
		}
		return result
	}

	fmt.Println("1 + 1 = ", sum(1, 1))
	fmt.Println("1 - 1 = ", minus(1, 1))
	fmt.Println("2 * 3 = ", multiply(2, 3, sum))
}
