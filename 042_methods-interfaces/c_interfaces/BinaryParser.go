package main

import (
	"fmt"
	"strconv"
)

type BinaryParser struct{}

func (bf BinaryParser) Parse(s string) int {
	value, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic("Failed to convert string to binary")
	}

	fmt.Println("string value", value)
	return int(value)
}
