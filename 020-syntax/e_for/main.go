package main

import "fmt"

func main() {
	fmt.Println("For-Loops:")

	numbers := [2]int{10, 90}
	var sum int
	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("sum of %v is %d\n", numbers, sum)
}
