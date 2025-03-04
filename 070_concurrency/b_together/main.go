package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	name  string
	score int
}

func (runner *Runner) run() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	duration := r.Intn(1000)

	fmt.Println("start: ", runner.name)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	runner.score = duration
}

func startRace(runner Runner, register chan Runner, wg *sync.WaitGroup) {
	runner.run()
	wg.Done()

	// register the name to the channel, but it is blocked, if no name is read
	register <- runner

	fmt.Printf("\n%s Awarding time: %s", time.Now().Format(time.StampNano), runner.name)
}

func main() {
	fmt.Println("\n100-Meter race ...\n")
	runners := []Runner{Runner{name: "Tom"}, {name: "Jerry"}, {name: "Donald"}}
	register := make(chan Runner)

	wg := sync.WaitGroup{}
	for _, runner := range runners {
		wg.Add(1)
		go startRace(runner, register, &wg)
	}

	wg.Wait()

	for index, _ := range runners {
		// read the registered name, and it frees the blocked process, see line-29
		runner := <-register
		fmt.Printf("\nPlatz %d (%d ms) %s\n", index+1, runner.score, runner.name)
	}

	time.Sleep(time.Duration(5) * time.Second)
}
