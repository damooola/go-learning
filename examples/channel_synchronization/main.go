// Program channel_synchronization shows using a channel purely to
// wait for a goroutine to finish, an alternative to the
// sync.WaitGroup from goroutines/main.go.
package main

import (
	"fmt"
	"time"
)

// worker sends on "done" once it finishes, purely as a signal — the
// value itself (true) doesn't matter, only the fact that a value
// arrived at all.
func worker(id int, done chan<- bool) {
	fmt.Printf("worker %d: starting\n", id)
	time.Sleep(time.Millisecond * 100) // pretend to do some work
	fmt.Printf("worker %d: finished\n", id)
	done <- true
}

func main() {
	// done chan<- bool in worker's signature means "a channel this
	// function only ever SENDS on" — a direction restriction that
	// documents intent and stops worker from accidentally reading it.
	done := make(chan bool, 1)
	go worker(1, done)

	// This blocks until worker actually sends, so main can't
	// continue (and the program can't exit) before the goroutine
	// finishes its work. Without this line, main could exit before
	// worker ever gets to print "finished", same risk as the
	// unmanaged goroutine back in goroutines/main.go.
	<-done
	fmt.Println("main: worker signaled it's done")

	// This is really the same problem sync.WaitGroup solves, just
	// solved with a channel instead. A channel is the natural
	// choice when you also want to receive an actual RESULT (not
	// just a "finished" signal) or only have one goroutine to wait
	// on; WaitGroup fits better when you're waiting on many
	// goroutines and don't need any data back from them, as in
	// goroutines/main.go's loop over names.

	// Dart comparison: this is the same shape as "await someFuture"
	// in Dart — pausing until a concurrently-running operation
	// signals completion. The difference is what's doing the
	// signaling: Dart's Future/await is built into the language for
	// exactly this purpose, while Go builds the same behavior out of
	// its general-purpose channel primitive.

	// Backend connection: this pattern (or its more common cousin,
	// receiving an actual value back on the channel instead of just
	// true) is how a handler kicks off background work and then
	// waits for the result before responding to the request.
}
