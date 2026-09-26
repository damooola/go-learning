// Program tickers demonstrates time.Ticker: a REPEATING event on a
// fixed interval, building directly on timers/main.go's one-shot
// time.Timer.
package main

import (
	"fmt"
	"time"
)

func main() {
	// A Timer's channel fires exactly once. A Ticker's channel fires
	// again and again, once per interval, until you call Stop().
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool) // the goroutine-signal pattern from channel_synchronization/main.go

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t.Format("15:04:05.000"))
			}
		}
	}()

	// Let the ticker fire a handful of times before shutting it down.
	time.Sleep(550 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

	// Forgetting ticker.Stop() leaks the ticker's underlying
	// goroutine and timer for as long as the program runs — unlike a
	// Timer, which naturally stops generating events after its one
	// fire, a Ticker keeps firing forever if nothing stops it.

	// Dart comparison: time.NewTicker is the same idea as Dart's
	// Timer.periodic(duration, callback) — both fire repeatedly on
	// an interval and both need an explicit cancel (ticker.Stop() vs.
	// timer.cancel()) or they keep running indefinitely.

	// Backend connection: tickers drive periodic background work —
	// polling a queue every few seconds, refreshing a cache,
	// emitting metrics on a schedule. The rate limiting lesson next
	// reuses this same "channel that fires on an interval" idea for
	// a different purpose: throttling instead of scheduling.
}
