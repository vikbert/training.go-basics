package main

import "fmt"

type mathFunc func(int, int) int

func main() {
	fmt.Println("1st class Functions:")

	var add mathFunc = func(a int, b int) int {
		return a + b
	}

	var sub mathFunc = func(a int, b int) int {
		return a - b
	}

	var mul = func(a int, b int, add mathFunc) int {
		result := 0
		for i := 0; i < b; i++ {
			result = add(result, a)
		}
		return result
	}

	fmt.Println("1 + 1 = ", add(1, 1))
	fmt.Println("1 - 1 = ", sub(1, 1))
	fmt.Println("1 * 0 = ", mul(1, 0, add))
}
