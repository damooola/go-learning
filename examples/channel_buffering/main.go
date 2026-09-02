// Program channel_buffering looks specifically at channel capacity,
// building on the brief buffered-channel mention in channels/main.go.
package main

import "fmt"

func main() {
	// The second argument to make is the buffer's capacity: how many
	// values it can hold before a send has to block.
	messages := make(chan string, 2)

	// Because the buffer holds 2, both of these sends complete
	// immediately without needing a goroutine on the other end —
	// there's nobody receiving yet, and that's fine.
	messages <- "buffered"
	messages <- "channel"

	// cap() reports the buffer's total capacity; len() reports how
	// many values are currently sitting in it, waiting to be read.
	fmt.Println("len:", len(messages), "cap:", cap(messages))

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("len after draining:", len(messages))

	// A 3rd send would block if nothing were reading, since the
	// buffer only holds 2 — proven here by starting the send in a
	// goroutine (so it CAN block) while main drains it concurrently.
	numbers := make(chan int, 2)
	numbers <- 1
	numbers <- 2
	fmt.Println("\nnumbers is full: len:", len(numbers), "cap:", cap(numbers))

	go func() {
		// This blocks until main frees a slot by receiving below.
		numbers <- 3
		fmt.Println("goroutine: send of 3 finally went through")
	}()

	fmt.Println("main:", <-numbers) // frees one slot, unblocking the goroutine
	fmt.Println("main:", <-numbers)
	fmt.Println("main:", <-numbers)

	// Dart comparison: think of a buffered channel like a
	// fixed-size Queue with a capacity limit built in — pushing onto
	// it is instant while there's room, and only blocks (pausing
	// that goroutine) once it's full, unlike Dart's Queue/List which
	// just keeps growing. An unbuffered channel (capacity 0, the
	// default from channels/main.go) is the extreme case: zero room,
	// so every single send waits for a receiver.

	// Backend connection: a buffered channel is the usual way to
	// build a bounded job queue — e.g. make(chan Job, 100) lets up
	// to 100 pending jobs queue up before producers are forced to
	// slow down and wait for workers to catch up, which is a simple
	// built-in form of backpressure.
}
