// Program range_over_channels is a more thorough look at "for v :=
// range ch", building on the brief version already used at the end
// of channels/main.go.
package main

import "fmt"

// queueWork is the producer: it sends every value, then closes the
// channel so range on the consuming side knows when to stop.
func queueWork(items []string, queue chan<- string) {
	for _, item := range items {
		queue <- item
	}
	close(queue)
}

func main() {
	queue := make(chan string, 3)
	go queueWork([]string{"wash", "dry", "fold"}, queue)

	// range over a channel receives one value per iteration and
	// keeps going until the channel is BOTH closed and drained — at
	// that point the loop ends on its own, no ok check or break
	// needed (contrast the manual for/ok/break loop in
	// channels/main.go, which is what range replaces).
	fmt.Println("processing queue:")
	for task := range queue {
		fmt.Println("-", task)
	}
	fmt.Println("queue drained, channel closed")

	// range over a channel never gives you an index (unlike range
	// over a slice) — there's only ever one loop variable, the
	// received value itself.
	numbers := make(chan int, 5)
	for n := 1; n <= 5; n++ {
		numbers <- n * n
	}
	close(numbers)

	total := 0
	for n := range numbers {
		total += n
	}
	fmt.Println("\nsum of squares 1..5:", total)

	// If queueWork never called close(queue), this range would hang
	// forever after "fold" — it has no other way to know the
	// producer is finished. Forgetting to close a channel a range
	// loop depends on is one of the most common goroutine leaks in
	// Go.

	// Dart comparison: this is close to "await for (var task in
	// stream)" over a Dart Stream — both consume values one at a
	// time until the source signals it's done (Stream's "done"
	// event, Go's channel close). The difference is Go's version can
	// share the exact same channel as an ordinary send/receive
	// target elsewhere in the program, while Dart's Stream and its
	// controller are more separated.

	// Backend connection: range over a channel is the natural way to
	// drain a queue of work items or process a stream of results
	// from another goroutine — you'll see it inside the worker
	// pool's worker function in the next lesson.
}
