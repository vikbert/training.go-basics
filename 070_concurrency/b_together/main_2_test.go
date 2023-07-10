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

	signal := make(chan bool)
	wg := new(sync.WaitGroup)

	numberProcesses := 5
	wg.Add(numberProcesses)

	for i := 0; i < numberProcesses; i++ {
		go subProcessWithSleep(r.Intn(10), signal, wg)
	}

	wg.Wait()

	for i := 0; i < numberProcesses; i++ {
		signal <- true
	}

	time.Sleep(time.Second)
}

func subProcessWithSleep(seconds int, signal chan bool, wg *sync.WaitGroup) {
	startUp(seconds)
	wg.Done()
	<-signal
	fmt.Println("#", seconds, "- finished at:", time.Now())
}
