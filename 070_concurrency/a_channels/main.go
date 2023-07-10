package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	randomInt := r.Intn(100)

	ch := make(chan int)
	go read(ch)
	ch <- randomInt
	time.Sleep(time.Millisecond)
}

func read(ch chan int) {
	for {
		name := <-ch
		fmt.Println("Number read from channel: ", name)
	}
}
