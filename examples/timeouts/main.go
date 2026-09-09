// Program timeouts demonstrates using select and time.After to limit waiting.
package main

import (
	"fmt"
	"time"
)

// doWork simulates an operation such as a database or external API call.
func doWork(delay time.Duration, output chan<- string) {
	time.Sleep(delay)
	output <- "work completed"
}

func main() {
	// This operation takes longer than we are willing to wait.
	// The buffer lets the worker send once and exit even after main times out.
	slowResult := make(chan string, 1)
	go doWork(200*time.Millisecond, slowResult)

	select {
	case result := <-slowResult:
		fmt.Println("slow operation:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("slow operation: timed out")
	}

	// This operation finishes before its timeout.
	fastResult := make(chan string, 1)
	go doWork(50*time.Millisecond, fastResult)

	select {
	case result := <-fastResult:
		fmt.Println("fast operation:", result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("fast operation: timed out")
	}

	// Dart comparison: this plays a similar role to Future.timeout.
	// A Go timeout stops this code from waiting, but it does not automatically
	// cancel doWork. We will use context for cancellation in a later lesson.
}
