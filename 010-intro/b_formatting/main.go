package main

import (
	"fmt"
)

func main() {
	var number int64 = 40

	fmt.Println("Formatting:")
	fmt.Println("Number: ", number)
	fmt.Printf("Number Prefix: %04d\n", number)
	fmt.Printf("Float points: %.4f\n", 3.1)
	fmt.Println("Number Binary: ", fmt.Sprintf("%b", number))
	fmt.Println("Number Hex: ", fmt.Sprintf("%X", number))
	fmt.Println("Number Type: ", fmt.Sprintf("%T", number))
}
