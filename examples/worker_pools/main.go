// Program worker_pools combines several earlier lessons — channel
// directions, range over channels, and close — into the pattern they
// were all building toward: several goroutines pulling jobs off one
// shared channel and pushing results onto another.
package main

import (
	"fmt"
	"time"
)

// worker is launched multiple times with the SAME jobs and results
// channels. jobs <-chan int / results chan<- int (channel_direction/
// main.go) documents that a worker only ever reads jobs and only
// ever writes results — never the other way around.
func worker(id int, jobs <-chan int, results chan<- int) {
	// range over jobs (range_over_channels/main.go) keeps every
	// worker pulling the next available job until jobs is closed and
	// drained — whichever worker happens to be free picks up the
	// next one, so work is spread across all of them automatically.
	for j := range jobs {
		fmt.Printf("worker %d: started job %d\n", id, j)
		time.Sleep(50 * time.Millisecond) // pretend this job takes some time
		fmt.Printf("worker %d: finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers before sending any jobs. They immediately
	// block inside their range loops, waiting for jobs to have
	// something to receive — starting them early costs nothing.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Closing jobs is what eventually lets every worker's range loop
	// end (closing_channels/main.go's broadcast: all 3 workers notice
	// the close once they've drained the jobs meant for them).
	close(jobs)

	// Collect exactly numJobs results. We know the exact count up
	// front, so a plain counted loop works — no need to close
	// results or range over it here.
	total := 0
	for a := 1; a <= numJobs; a++ {
		total += <-results
	}
	fmt.Println("\nsum of all results:", total)

	// Dart comparison: Dart has no built-in worker pool primitive —
	// the closest match is manually spinning up several Isolates and
	// giving them SendPort/ReceivePort pairs, which is considerably
	// more setup than "launch a few goroutines against one channel."
	// Go's goroutines being cheap (thousands are normal, see
	// goroutines/main.go) is what makes this pattern so lightweight.

	// Backend connection: worker pools are the standard way to bound
	// concurrency — e.g. only ever running 5 downstream API calls or
	// database queries at once, no matter how many requests are
	// queued, by giving every request to the same small set of
	// worker goroutines instead of spawning one goroutine per request.
}
