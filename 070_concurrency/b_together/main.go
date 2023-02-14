package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startUp(id int) {
	fmt.Printf("#%d starting up ...\n", id)
	seconds := rand.Intn(10) + 1
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("#%d ready!\n", id)
}

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// TODO

	// wait a second at the very end of main() to give goroutine enough time to print
	time.Sleep(time.Second)
}
