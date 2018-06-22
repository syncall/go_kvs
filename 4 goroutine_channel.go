// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func alpahbet(done chan bool) {
	for char := 'a'; char < 'a'+26; char++ {
		fmt.Printf("%c ", char)
		time.Sleep(1*time.Millisecond)
	}
	done <- true
}

func numbers(done chan bool) {
	for number := 1; number < 27; number++ {
		fmt.Printf("%d ", number)
		time.Sleep(1*time.Millisecond)
	}
	done <- true
}

func main() {
	// Set number of logical processors
	done := make (chan bool)

	go alpahbet(done)

	go numbers(done)

	<-done
	<-done
	fmt.Printf("\nDone")
}
