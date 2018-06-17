package main

import (
	"fmt"
	"time"
	"strconv"
)

func worker(input chan string) {
	for {
		time.Sleep(100*time.Millisecond)
		msg := <-input
		fmt.Println("Processing: ", msg)
	}
}

func main() {
	messages := make(chan string, 2)

	go worker(messages)

	for i:=1; i<5; i++ {
		messages <- strconv.Itoa(i)
		fmt.Println("Scheduling ", i)
	}

	time.Sleep(1*time.Second)
}
