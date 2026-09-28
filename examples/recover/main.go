// Program recover demonstrates stopping a panic (panic/main.go) from
// crashing the whole program, using defer (defer/main.go) — the
// three concurrency-adjacent lessons this ties together.
package main

import "fmt"

// mayPanic simulates something going wrong badly enough to justify a
// panic rather than a returned error.
func mayPanic() {
	panic("a problem")
}

// safelyDo wraps mayPanic in a deferred recover. recover only does
// anything when called DIRECTLY inside a deferred function — calling
// it anywhere else always returns nil, even during an active panic.
func safelyDo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	mayPanic()
	fmt.Println("this line never runs — panic skips straight to the deferred func")
}

func main() {
	safelyDo()
	// Without recover inside safelyDo, panic/main.go's example
	// already showed what happens next: the whole program crashes,
	// and this line would never run either. Here, recover stopped
	// the panic from propagating past safelyDo, so main carries on
	// completely normally.
	fmt.Println("main continues normally after safelyDo()")

	// Important limit: recover only works within the SAME goroutine
	// as the panic. A panic inside a goroutine launched with "go f()"
	// that's never recovered INSIDE that goroutine crashes the whole
	// program, no matter what recover calls exist elsewhere in the
	// program — there's no cross-goroutine try/catch.

	// Dart comparison: recover is closest to a catch block — panic
	// throws, recover catches, both stop the error from propagating
	// further up the call stack. The difference is structural: Dart's
	// try/catch wraps the risky code directly, while Go's recover
	// must sit inside a deferred function, one level removed from the
	// code that actually panics.

	// Backend connection: an HTTP server commonly wraps every request
	// handler in a recover so ONE handler panicking (e.g. an
	// unexpected nil from malformed input) returns a 500 to that one
	// request instead of crashing the entire server process and
	// taking down every other in-flight request with it.
}
