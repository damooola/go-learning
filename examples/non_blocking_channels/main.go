// Program non_blocking_channels demonstrates using select with a
// default case to attempt a channel operation without blocking.
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// A plain <-messages would block here forever, since nobody is
	// sending on it. Adding a default case changes that: select
	// tries the messages case, finds it's not ready, and immediately
	// falls through to default instead of waiting.
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// The same non-blocking idea applies to sends. messages has no
	// buffer and nobody is receiving, so a plain "messages <- value"
	// would also block forever — default lets us try instead of wait.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multiple non-blocking checks can share one default: select
	// still only runs ONE case per attempt, but here neither
	// messages nor signals is ready, so default runs regardless of
	// how many other cases are listed.
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}

	// Dart comparison: Dart's Stream is push/callback based — you
	// subscribe and get notified later, but there's no built-in "is a
	// value available right now, otherwise move on" check. select
	// with default is a synchronous poll with no Dart equivalent;
	// approximating it in Dart would mean maintaining your own flag
	// or buffer alongside the stream.

	// Backend connection: this differs from the timeouts lesson,
	// which waits up to a bounded duration before giving up. A
	// non-blocking select waits zero time — it's for checking
	// something (e.g. a shutdown signal) once per loop iteration
	// without ever stalling that iteration on the check.
}
