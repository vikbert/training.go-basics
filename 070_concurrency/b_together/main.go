package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startUp(id int) {
	fmt.Printf("#%d starting up ...\n", id)
	seconds := rand.Intn(10) + 1
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("#%d ready!\n", id)
}

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go myTask(ch, wg)
	ch <- r.Intn(10)
	ch <- r.Intn(10)
	ch <- r.Intn(10)

	wg.Wait()

	// wait a second at the very end of main() to give goroutine enough time to print
	time.Sleep(time.Second)
}

func myTask(ch chan int, wg *sync.WaitGroup) {
	for {
		number := <-ch
		startUp(number)
		fmt.Println("Current date and time:", time.Now())
		wg.Done()
	}
}
