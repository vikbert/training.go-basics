package main

import (
	"fmt"
	"strconv"
)

type BarFormatter struct{}

func (bf BarFormatter) Format(i int) string {
	formatInt := strconv.Itoa(i)
	fmt.Println("内容", formatInt)
	return formatInt
}
