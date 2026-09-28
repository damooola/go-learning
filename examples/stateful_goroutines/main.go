// Program stateful_goroutines is an alternative to mutexes/main.go's
// approach: instead of locking shared state so many goroutines can
// touch it safely, keep the state inside exactly ONE goroutine, and
// have every other goroutine reach it only through channels. This is
// Go's "share memory by communicating" proverb, in contrast to
// mutexes' "communicate by sharing memory" (a shared map, guarded by
// a lock).
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// readOp/writeOp are requests sent TO the state-owning goroutine.
// resp is a channel included INSIDE the request itself — the only
// way the state goroutine can reply directly to the one goroutine
// that asked, rather than broadcasting to whoever happens to be
// listening.
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)
	done := make(chan bool)

	// The ONLY goroutine that ever touches "state" directly. Every
	// other goroutine reaches it exclusively by sending on reads/
	// writes — never a shared pointer, never a lock. select (from
	// the select lesson) lets this one goroutine serve whichever
	// kind of request arrives next.
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			case <-done:
				return
			}
		}
	}()

	var wg sync.WaitGroup

	// 20 goroutines, each owning a distinct key: write a random
	// value, then immediately read the same key back. Because every
	// read and write is funneled through the single state goroutine
	// above, and each key belongs to only one writer here, every
	// read is guaranteed to see exactly what that goroutine just
	// wrote — no race is possible, without a single Lock() call.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			val := rand.Intn(100)

			write := writeOp{key: key, val: val, resp: make(chan bool)}
			writes <- write
			<-write.resp

			read := readOp{key: key, resp: make(chan int)}
			reads <- read
			got := <-read.resp

			if got != val {
				fmt.Printf("key %d: wrote %d, read back %d (mismatch!)\n", key, val, got)
			}
		}(i)
	}

	wg.Wait()
	done <- true
	fmt.Println("all 20 reads matched their writes — state never left its owning goroutine")

	// Dart comparison: this resembles routing every state change
	// through a single StreamController/BLoC-style event handler
	// rather than letting multiple pieces of code mutate a shared
	// object directly — one owner, everyone else sends messages to
	// it. The difference is Dart usually reaches for that as an
	// architectural choice; here it's solving a genuine data-race
	// problem that only exists because goroutines run in parallel.

	// Backend connection: this pattern shows up as a dedicated
	// "manager" goroutine owning something like an in-memory
	// connection pool or cache, where every other goroutine sends
	// requests to it instead of touching the underlying data
	// structure themselves — useful when the access pattern is more
	// complex than a Mutex comfortably expresses.
}
