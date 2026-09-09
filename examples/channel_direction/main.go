// Program channel_direction demonstrates send-only and receive-only channels.
package main

import "fmt"

// produce receives a send-only view of the channel. It can send and close the
// channel, but trying to receive with <-output would be a compile-time error.
func produce(output chan<- int) {
	for number := 1; number <= 4; number++ {
		output <- number
	}
	close(output)
}

// multiply receives an input-only channel and an output-only channel.
// This directional signature clearly describes how data flows through it.
func multiply(input <-chan int, output chan<- int, factor int) {
	for number := range input {
		output <- number * factor
	}
	close(output)
}

// printResults can receive values, but it cannot send or close the channel.
func printResults(input <-chan int) {
	for result := range input {
		fmt.Println("result:", result)
	}
}

func main() {
	// make creates bidirectional channels. Passing them into the functions below
	// safely narrows what each function is allowed to do.
	numbers := make(chan int)
	results := make(chan int)

	go produce(numbers)
	go multiply(numbers, results, 2)

	// Data flows in one direction through this small pipeline:
	// produce -> numbers -> multiply -> results -> printResults
	printResults(results)

	// Dart comparison: SendPort and ReceivePort also separate sending from
	// receiving, but Go's chan<- and <-chan are restricted views of the same
	// channel and are checked by the compiler.
}
