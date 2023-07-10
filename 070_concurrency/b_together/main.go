package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	name string
}

func (runner *Runner) run() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	duration := r.Intn(10)

	fmt.Printf("%s: start running\n %ds\n", runner.name, duration)
	time.Sleep(time.Duration(duration) * time.Second)
}

func startRace(runner Runner, awarding chan string, wg *sync.WaitGroup) {
	runner.run()
	wg.Done()
	name := <-awarding
	fmt.Printf("Awarding: %s %s\n", name, time.Now())
}

func main() {
	wg := sync.WaitGroup{}
	runners := []Runner{
		Runner{"Tom"},
		{"Jerry"},
		{"Donald"},
	}

	awarding := make(chan string)

	fmt.Println("Start running ...\n")
	for _, runner := range runners {
		wg.Add(1)
		go startRace(runner, awarding, &wg)
	}

	wg.Wait()
	for _, runner := range runners {
		awarding <- runner.name
	}
	time.Sleep(time.Duration(10) * time.Second)

	fmt.Println("\nAll runners completed the course!")
}
