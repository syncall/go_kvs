// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
	"runtime"
)

func alpahbet() {
	for char := 'a'; char < 'a'+26; char++ {
		fmt.Printf("%c ", char)
	}
}

func numbers() {
	for number := 1; number < 27; number++ {
		fmt.Printf("%d ", number)
	}
}

func main() {
	// Set number of logical processors
	max_proc := 1
	runtime.GOMAXPROCS(max_proc)

	alpahbet()

	numbers()

	time.Sleep(1* time.Second)
	fmt.Printf("\nDone with %v Threads", max_proc)
}
