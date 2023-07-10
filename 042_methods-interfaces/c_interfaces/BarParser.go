package main

import (
	"fmt"
	"strconv"
)

type BarParser struct{}

func (bp BarParser) Parse(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic("Failed to convert string to int")
	}

	fmt.Println("数值", value)
	return int(value)
}
