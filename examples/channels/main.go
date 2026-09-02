// Program channels demonstrates how goroutines communicate safely,
// building on goroutines/main.go.
package main

import "fmt"

func main() {
	// make(chan T) creates a channel that carries values of type T.
	// Sending (ch <- value) and receiving (<-ch) on an UNBUFFERED
	// channel each block until the other side is ready — this is
	// what makes channels double as a synchronization tool, not just
	// a way to pass data.
	messages := make(chan string)

	go func() {
		// This send blocks until main is ready to receive below.
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println("received:", msg)

	// A buffered channel can hold N values without a receiver being
	// ready yet — sends only block once the buffer is full.
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	// A 3rd send here would block, since the buffer only holds 2.
	fmt.Println("\nbuffered:", <-buffered, <-buffered)

	// close(ch) signals "no more values are coming." Receiving from
	// a closed channel returns the zero value immediately instead of
	// blocking forever — the second return value (ok) tells you
	// whether that came from an actual send or just the close.
	jobs := make(chan int, 3)
	jobs <- 10
	jobs <- 20
	jobs <- 30
	close(jobs)

	fmt.Println("\nreading a closed channel manually:")
	for {
		value, ok := <-jobs
		if !ok {
			fmt.Println("channel closed, stopping")
			break
		}
		fmt.Println("got:", value)
	}

	// range over a channel does that loop for you automatically —
	// it reads until the channel is closed, then stops.
	moreJobs := make(chan int, 3)
	moreJobs <- 100
	moreJobs <- 200
	close(moreJobs)

	fmt.Println("\nrange over a channel:")
	for value := range moreJobs {
		fmt.Println("got:", value)
	}

	// select waits on multiple channels at once and runs whichever
	// case becomes ready first — like a switch, but for channel
	// operations instead of values.
	channelA := make(chan string)
	channelB := make(chan string)
	go func() { channelA <- "from A" }()
	go func() { channelB <- "from B" }()

	fmt.Println("\nselect over two channels:")
	for i := 0; i < 2; i++ {
		select {
		case a := <-channelA:
			fmt.Println("received", a)
		case b := <-channelB:
			fmt.Println("received", b)
		}
	}

	// Dart comparison: Dart has no built-in channel type. The
	// closest everyday equivalent is Stream/StreamController, but
	// those model an ASYNC EVENT PIPELINE (subscribe, get called
	// back over time) rather than blocking hand-offs between two
	// concurrently running pieces of code. Dart's Isolates do use a
	// channel-like SendPort/ReceivePort pair for cross-isolate
	// messages, which is the closer conceptual match, but isolates
	// themselves are a much heavier tool than a goroutine+channel.

	// Backend connection: channels are how goroutines coordinate
	// safely instead of sharing memory directly — common patterns
	// include a worker pool (many goroutines pulling jobs off one
	// channel) and pipelines (one goroutine's output channel feeds
	// the next stage's input channel).
}
