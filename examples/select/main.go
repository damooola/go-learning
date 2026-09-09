// Program select demonstrates waiting on multiple channel operations.
package main

import (
	"fmt"
	"time"
)

// sendAfter simulates work that finishes after a delay.
func sendAfter(output chan<- string, message string, delay time.Duration) {
	time.Sleep(delay)
	output <- message
}

func main() {
	fastResult := make(chan string)
	slowResult := make(chan string)

	go sendAfter(fastResult, "fast job finished", 100*time.Millisecond)
	go sendAfter(slowResult, "slow job finished", 250*time.Millisecond)

	// One select handles one ready case. Repeat it because we expect two results.
	for received := 0; received < 2; received++ {
		select {
		case message := <-fastResult:
			fmt.Println("fastResult:", message)
		case message := <-slowResult:
			fmt.Println("slowResult:", message)
		}
	}

	// Dart comparison: this resembles reacting to the first Future completed by
	// Future.any. Unlike Future.any, select works with channel operations and the
	// loop above lets us continue until both results have arrived.

	// Backend connection: a service might start independent API or database
	// calls concurrently, then use select to react as each result arrives.
}
