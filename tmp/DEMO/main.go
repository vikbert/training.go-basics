package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("demo.txt")

	if err != nil {
		fmt.Println("can not open file", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	writeString, err := file.WriteString("new words")
	if err != nil {
		return
	}

	fmt.Println(writeString)
}
