package main

import (
	"fmt"
	"time"
)

func main() {

	sChan := make(chan string)
	go writeString(sChan)

	iChan := make(chan int)
	go writeInteger(iChan)

	for {
		select {
		case msg1 := <-sChan:
			fmt.Printf("received: %s\n\n", msg1)
		case msg2 := <-iChan:
			fmt.Printf("received: %d\n\n", msg2)
		}
	}

}

func writeString(ch chan string) {
	for {
		time.Sleep(time.Millisecond * 800)
		msg := "HELLO WORLD"
		ch <- msg
		fmt.Println("Writing message: ", msg)
	}
}

func writeInteger(ch chan int) {
	for {
		time.Sleep(time.Millisecond * 4100)
		ch <- 1100
		fmt.Println("Writing number: ", 1100)
	}
}
