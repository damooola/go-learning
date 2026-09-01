// Program goroutines demonstrates Go's lightweight concurrency
// primitive: the goroutine.
package main

import (
	"fmt"
	"sync"
)

func sayHello(name string) {
	fmt.Println("hello,", name)
}

func main() {
	// A normal call runs and finishes before the next line executes.
	sayHello("synchronous")

	// The "go" keyword launches a function in a new goroutine and
	// returns immediately — it does NOT wait for sayHello to finish.
	go sayHello("goroutine")

	// Because main() doesn't wait for goroutines by default, the
	// program can (and often does) exit before "hello, goroutine"
	// ever gets a chance to print. This line proves it: main keeps
	// going immediately, without pausing for the line above.
	fmt.Println("main kept going without waiting")

	// sync.WaitGroup is the standard way to wait for a group of
	// goroutines to finish. Add(1) registers one goroutine to wait
	// for, Done() signals it finished, and Wait() blocks until every
	// registered goroutine has called Done().
	var wg sync.WaitGroup
	names := []string{"Ada", "Grace", "Linus"}

	fmt.Println("\nlaunching goroutines with a WaitGroup:")
	for _, name := range names {
		wg.Add(1)
		// Pass name as a parameter (rather than reading the loop
		// variable from the closure) so each goroutine captures its
		// own copy — otherwise every goroutine could see whichever
		// name the loop had reached by the time it actually ran.
		go func(name string) {
			defer wg.Done()
			sayHello(name)
		}(name)
	}
	wg.Wait()
	fmt.Println("all goroutines finished")

	// Dart comparison: Dart's async/await runs on a single-threaded
	// event loop by default — it gives you CONCURRENCY (interleaving
	// tasks while waiting on I/O) but not true PARALLELISM across
	// CPU cores unless you spin up an Isolate, which is heavyweight
	// and has no shared memory with the main isolate. A goroutine is
	// much cheaper to create (thousands are normal) and the Go
	// runtime schedules goroutines across real OS threads, so they
	// can run truly in parallel on a multi-core machine, while
	// sharing memory directly (which is why coordination tools like
	// sync.WaitGroup and, next topic, channels, matter).

	// Backend connection: Go's net/http server automatically handles
	// every incoming request in its own goroutine, which is a big
	// part of why Go backends handle high concurrent load well
	// without you writing any special code for it. You'll reach for
	// goroutines yourself for things like firing off several
	// downstream calls at once instead of one after another.
}
