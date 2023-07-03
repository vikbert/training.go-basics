package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch:")

	switch weekday := time.Now().Weekday(); weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("is weekend")
	default:
		fmt.Println("is not weekend")
	}
}
