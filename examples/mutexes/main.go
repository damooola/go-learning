// Program mutexes demonstrates sync.Mutex: protecting shared state
// that's more than a single number, unlike atomic_counters/main.go's
// atomic.Uint64.
package main

import (
	"fmt"
	"sync"
)

// container holds shared state — a map, not just a number — so
// sync/atomic (which only has counters for individual numeric types)
// can't protect it. mu guards every access to counters.
type container struct {
	mu       sync.Mutex
	counters map[string]int
}

// inc locks mu before touching counters and unlocks it (via defer,
// same reasoning as defer wg.Done() in wait_groups/main.go) once
// done. Whichever goroutine holds the lock has exclusive access;
// every other goroutine calling inc blocks until Unlock happens.
func (c *container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := container{counters: map[string]int{"a": 0, "b": 0}}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	// Two goroutines hammer the SAME key "a" concurrently, one
	// hammers "b". Without the mutex inside inc, two goroutines could
	// both read counters["a"] before either writes back — the exact
	// same lost-update race that atomic_counters/main.go warned about,
	// just now on a map instead of a plain number.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait()

	fmt.Println("a:", c.counters["a"]) // always exactly 20000
	fmt.Println("b:", c.counters["b"]) // always exactly 10000

	// Dart comparison: Dart's single-threaded event loop means it
	// almost never needs an explicit lock — only one piece of code
	// runs at a time, so there's no window for two operations to
	// interleave on the same data. Go's goroutines run genuinely in
	// parallel and share memory directly, so a Mutex is how you
	// carve out "only one goroutine touches this at a time," the
	// same way Dart's isolation is automatic and Go's is opt-in.

	// Backend connection: reach for atomic (previous lesson) when
	// shared state is a single number; reach for Mutex whenever it's
	// anything more — a map, a struct with multiple fields, a slice
	// — which describes most real shared state in a backend service,
	// e.g. an in-memory cache guarded by a mutex.
}
