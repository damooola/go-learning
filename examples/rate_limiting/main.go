// Program rate_limiting reuses tickers/main.go's time.Ticker for a
// different purpose: throttling how fast work is allowed to
// proceed, instead of scheduling a repeating action.
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// One tick every 200ms caps how often the loop below can proceed
	// — regardless of how fast requests themselves arrive, at most
	// one gets handled per interval.
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop() // same cleanup habit as tickers/main.go

	fmt.Println("basic rate limiting:")
	for req := range requests {
		<-limiter.C // blocks until the next tick — this IS the throttle
		fmt.Println("request", req, "at", time.Now().Format("15:04:05.000"))
	}

	// Real traffic often arrives in bursts rather than evenly
	// spaced. A buffered channel, pre-filled with a few timestamps,
	// lets a short burst through immediately while still enforcing
	// the same average rate over time.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// This goroutine keeps burstyLimiter topped up at the ticker's
	// rate. Once the pre-filled 3 slots are used, requests fall back
	// to waiting on each new tick, same as the basic version above.
	burstyTicker := time.NewTicker(200 * time.Millisecond)
	defer burstyTicker.Stop()
	go func() {
		for t := range burstyTicker.C { // range_over_channels: runs until Stop
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	fmt.Println("\nbursty rate limiting:")
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, "at", time.Now().Format("15:04:05.000"))
	}
	// Expect the first 3 bursty requests to print almost immediately
	// (using the pre-filled slots), then the last 2 to wait ~200ms
	// apart each, once the pre-fill is used up.

	// Dart comparison: Dart has no built-in rate limiter either —
	// you'd hand-roll the same idea with a Timer.periodic refilling
	// a token count, or reach for a package. Nothing here is
	// Go-specific in concept, just built from primitives you already
	// have (Ticker + a buffered channel as a token bucket).

	// Backend connection: this is the classic TOKEN BUCKET pattern —
	// used to cap outbound requests to a third-party API that
	// enforces its own rate limit, or to protect your own service
	// from being overwhelmed by a sudden spike of incoming requests.
}
