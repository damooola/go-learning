// Program panic demonstrates panic: Go's way to stop normal
// execution immediately when something is unrecoverable, rather than
// returning an error for the caller to handle.
package main

import (
	"fmt"
	"os"
)

func main() {
	// A deferred call still runs even though panic below prevents
	// execution from ever reaching the end of main normally — see
	// the defer lesson for why: defer's "runs no matter how the
	// function returns" guarantee holds during a panic too.
	defer fmt.Println("deferred cleanup still runs")

	_, err := os.Open("does/not/exist.txt")
	if err != nil {
		// panic accepts any value, not just a string — an error is
		// idiomatic here since it already describes what went wrong.
		// Once panic runs, no more lines in this function execute;
		// Go starts unwinding the stack, running deferred calls as
		// it goes.
		panic(err)
	}

	fmt.Println("this line never runs")
}

// Running this program prints "deferred cleanup still runs", then Go
// prints the panic value and a stack trace to stderr, then exits
// with a non-zero status — this is the EXPECTED result of running
// this file, not a bug in the lesson.

// Dart comparison: panic is closest to Dart's throw with no matching
// catch — the program terminates instead of continuing. The
// difference is Go's convention: reach for panic only for programmer
// errors and truly unrecoverable situations, and use a normal
// returned error (errors/main.go) for anything a caller might
// reasonably want to handle. Dart's exceptions are used far more
// liberally for ordinary error handling than Go's panic is meant to be.

// Backend connection: an unhandled panic in a request handler brings
// down the whole server process unless something recovers from it —
// which is exactly what the recover lesson covers next.
