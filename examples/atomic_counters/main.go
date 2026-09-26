// Program atomic_counters demonstrates sync/atomic: a way to safely
// share a single counter across many goroutines WITHOUT a mutex.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// atomic.Uint64 (not a plain uint64) is the whole point here: its
	// Add and Load methods are guaranteed safe to call from many
	// goroutines at once. A plain "ops++" from 50 goroutines at once
	// would be a DATA RACE — two goroutines could both read the same
	// old value before either writes back, silently losing updates.
	var ops atomic.Uint64
	var wg sync.WaitGroup

	// 50 goroutines, each incrementing the SAME shared counter 1000
	// times — 50,000 increments total, all racing against each other.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
	// Always prints exactly 50000. Try replacing atomic.Uint64 with a
	// plain uint64 and ops++ (see the commented block below) and this
	// number becomes unpredictable — usually a bit less than 50000,
	// and it changes between runs, because some increments get lost.

	// var unsafeOps uint64
	// var wg2 sync.WaitGroup
	// for i := 0; i < 50; i++ {
	// 	wg2.Add(1)
	// 	go func() {
	// 		defer wg2.Done()
	// 		for c := 0; c < 1000; c++ {
	// 			unsafeOps++ // NOT safe: a plain read-modify-write, racing
	// 		}
	// 	}()
	// }
	// wg2.Wait()
	// fmt.Println("unsafeOps:", unsafeOps) // unreliable, often < 50000

	// Dart comparison: Dart rarely needs this at all. Its
	// single-threaded event loop means only one piece of code ever
	// runs at a time (barring Isolates, which don't share memory
	// anyway), so there's no equivalent of two goroutines racing to
	// update the same variable. Go's goroutines genuinely run in
	// parallel on multiple CPU cores and share memory directly,
	// which is exactly the situation atomic operations protect
	// against.

	// Backend connection: atomic counters are the lightweight option
	// for shared state that's JUST a number — request counts, active
	// connection counts, simple metrics. For anything more than a
	// single number (multiple related fields, a map, a struct), a
	// sync.Mutex is the more general tool, and is worth learning
	// right after this.
}
