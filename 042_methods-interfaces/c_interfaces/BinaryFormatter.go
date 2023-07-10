package main

import (
	"fmt"
	"strconv"
)

type BinaryFormatter struct{}

func (bf BinaryFormatter) Format(i int) string {
	formatInt := strconv.FormatInt(int64(i), 2)
	fmt.Println("binary value", formatInt)
	return formatInt
}
