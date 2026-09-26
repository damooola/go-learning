// Program closing_channels goes deeper on close(ch): it isn't just
// "stop this one range loop" — it's a BROADCAST that every current
// and future receiver on that channel observes.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// --- Part 1: close reaches every waiting receiver at once ---
	// A sent value only reaches ONE receiver — whichever goroutine
	// happens to be ready to receive it. Closing is different: every
	// goroutine currently blocked on a receive from "ready" wakes up
	// at the same moment, each getting ready's zero value (false)
	// with ok=false. One close, three workers released together.
	ready := make(chan bool)
	var wg sync.WaitGroup

	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-ready // blocks here until ready is closed
			fmt.Printf("worker %d: starting\n", id)
		}(id)
	}

	fmt.Println("closing 'ready' releases all three workers at once:")
	close(ready)
	wg.Wait()

	// --- Part 2: producer closes jobs; a separate goroutine notices ---
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// The comma-ok receive from channels/main.go, now
			// checked inside a goroutine that's running concurrently
			// with the sends below — close can happen mid-loop.
			j, more := <-jobs
			if !more {
				fmt.Println("\nconsumer: jobs is closed, no more work")
				done <- true
				return
			}
			fmt.Println("consumer: received job", j)
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs) // safe to close even while the consumer goroutine is mid-loop
	<-done      // wait for the consumer to notice the close and finish

	// Receiving again from a channel that's already closed AND
	// drained is safe to repeat as many times as you like — it never
	// panics, it just keeps returning the zero value with ok=false.
	_, ok := <-jobs
	fmt.Println("receiving from jobs again, ok:", ok)

	// Dart comparison: closing a channel is closer to calling close()
	// on a StreamController than anything Future-based — both notify
	// every current subscriber that no more events are coming. The
	// difference is Dart delivers that as a stream "done" callback,
	// while Go represents it as an ordinary channel receive that
	// happens to come back with ok=false.

	// Backend connection: the broadcast property is what makes close
	// useful as a stop/cancel signal for a whole worker pool — closing
	// one shared channel tells every worker goroutine to stop at once,
	// which is simpler than sending each of them an individual value.
}
