package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions:")

	a, b, c := 1, 1, 1
	fmt.Println("a, b, c: ", a, b, c)
	doubleVars(&a, &b, &c)
	fmt.Println("a, b, c: ", a, b, c)
}

func doubleVars(pointers ...*int) {
	for _, ptr := range pointers {
		*ptr *= 2
	}
}
