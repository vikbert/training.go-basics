package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch:")

	isAtWeekend, weekday := isWeekend()
	if isAtWeekend {
		fmt.Printf("Today (%s) is weekend\n", weekday.String())
	} else {
		fmt.Printf("Today (%s) is NOT weekend\n", weekday.String())
	}
}

func isWeekend() (bool, time.Weekday) {
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Saturday, time.Sunday:
		return true, weekday
	default:
		return false, weekday
	}
}
