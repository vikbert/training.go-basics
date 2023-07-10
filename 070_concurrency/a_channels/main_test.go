package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	ch <- 9
	ch <- 8
	ch <- 7
	go sayHello(ch, wg)
	wg.Wait()
}

func sayHello(ch chan int, wg *sync.WaitGroup) {
	for {
		name := <-ch
		fmt.Println("name read from channel: ", name)
		wg.Done()
	}
}

func add(a, b int, wg *sync.WaitGroup) int {
	defer func() {
		wg.Done()
	}()
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("\nstart addition: %d + %d = %d\n", a, b, a+b)

	return a + b
}
