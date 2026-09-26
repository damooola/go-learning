// Program timers demonstrates time.Timer: a one-shot event that
// fires once after a delay, and can be stopped before it does.
package main

import (
	"fmt"
	"time"
)

func main() {
	// NewTimer starts counting immediately and returns a *Timer.
	// timer1.C is a receive-only channel that gets exactly one value
	// (the current time) once the duration has elapsed.
	timer1 := time.NewTimer(150 * time.Millisecond)

	// This blocks until timer1 fires — there's nothing racing it
	// with select here, just a plain receive.
	<-timer1.C
	fmt.Println("timer1 fired")

	// A timer can be stopped before it fires. Stop returns true if
	// it successfully prevented the firing, false if the timer had
	// already fired (or already been stopped).
	timer2 := time.NewTimer(200 * time.Millisecond)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired") // this should never print
	}()
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("timer2 stopped before it fired")
	}

	// Give the goroutine above a moment to prove it never runs.
	// (A real program would use a WaitGroup or channel instead of
	// sleeping — see goroutines/main.go and channel_synchronization/
	// main.go. A short sleep keeps this example focused on timers.)
	time.Sleep(300 * time.Millisecond)

	// time.After (used already in timeouts/main.go) is really just
	// time.NewTimer(d).C in disguise — a convenience for when you
	// only need the channel and don't need Stop().
	fmt.Println("\n(reminder: time.After(d) is shorthand for" +
		" time.NewTimer(d).C when you don't need to Stop it)")

	// Dart comparison: time.NewTimer plays the same role as Dart's
	// Timer(duration, callback) — both schedule a one-shot event, and
	// both can be cancelled before firing (Dart's timer.cancel() vs.
	// Go's timer.Stop()). The difference is what "firing" looks
	// like: Dart runs your callback directly, while Go delivers a
	// value on a channel that your own code then receives from.

	// Backend connection: timers back the timeouts lesson's
	// time.After, and show up directly whenever a handler needs to
	// delay work or cancel a scheduled action if it's no longer
	// needed, e.g. debouncing repeated requests.
}
