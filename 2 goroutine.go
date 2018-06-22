// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"runtime"
	"time"
)

func alpahbet() {
	for char := 'a'; char < 'a'+26; char++ {
		fmt.Printf("%c ", char)
		time.Sleep(time.Millisecond)
	}
}

func numbers() {
	for number := 1; number < 27; number++ {
		fmt.Printf("%d ", number)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// Set number of logical processors
	max_proc := 1
	runtime.GOMAXPROCS(max_proc)

	go alpahbet()

	go numbers()

	time.Sleep(2* time.Second)
	fmt.Printf("\nDone with %v Threads", max_proc)
}
