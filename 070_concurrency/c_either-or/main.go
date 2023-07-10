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
			fmt.Println("received", msg1)
		case msg2 := <-iChan:
			fmt.Println("received", msg2)
		}
	}

}

func writeString(ch chan string) {
	for {
		time.Sleep(time.Millisecond * 800)
		msg := "hello world"
		ch <- msg
		fmt.Println("New message: ", msg)
	}
}

func writeInteger(ch chan int) {
	for {
		time.Sleep(time.Millisecond * 1100)
		ch <- 999
		fmt.Println("New number: ", 999)
	}
}
