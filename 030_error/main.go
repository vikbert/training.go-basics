package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("go.mod")
	fmt.Println(string(bytes), err)
}
