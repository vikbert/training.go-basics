package main

import (
	"fmt"
)

func main() {
	fmt.Println("Variables & constants:")

	var kb = 1 << 10
	fmt.Println("KB", kb)

	var mb = 1 << 20
	fmt.Println("MB", mb)

	var gb = 1 << 30
	fmt.Println("GB", gb)

	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
	)

	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
}
