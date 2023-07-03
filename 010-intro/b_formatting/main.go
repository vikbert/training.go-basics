package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int64 = 40
	fmt.Println("Formatting:")
	fmt.Println("Number: ", number)
	fmt.Println("Number Prefix: ", fmt.Sprintf("%04d", number))
	fmt.Println("Number Binary: ", strconv.FormatInt(number, 2))
	fmt.Println("Number Hex: ", fmt.Sprintf("%X", number))
	fmt.Println("Number Type: ", fmt.Sprintf("%T", number))
}
