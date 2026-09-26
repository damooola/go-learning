// Program wait_groups looks at sync.WaitGroup in more depth than
// goroutines/main.go did, by using it to solve worker_pools/main.go's
// exercise 4: collecting results without knowing the job count in
// advance.
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker takes *sync.WaitGroup (a pointer), not sync.WaitGroup. A
// WaitGroup tracks internal state that must never be copied — go vet
// itself warns about this — so it's always passed around by address,
// same reasoning as the pointer receivers from methods/main.go.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(20 * time.Millisecond)
		results <- j * 2
	}
}

func main() {
	const numJobs = 7
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		// wg.Add(1) happens HERE, before the goroutine starts — not
		// inside worker itself. If Add ran inside the goroutine,
		// wg.Wait() below could run before any Add happened at all
		// (goroutines aren't guaranteed to start in order), and Wait
		// would return immediately having never actually waited.
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// worker_pools/main.go's main goroutine collected exactly numJobs
	// results because it already knew that count. Here, a separate
	// goroutine waits for every worker to call Done(), THEN closes
	// results — which lets main below use a plain range over results
	// (range_over_channels/main.go) instead of a counted loop, with
	// no need to know numJobs at all.
	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for r := range results {
		total += r
	}
	fmt.Println("sum of all results:", total)

	// Dart comparison: the nearest match is Future.wait([...]) —
	// both wait for a fixed group of concurrent operations to all
	// finish. The difference is Future.wait takes a known list of
	// Futures up front, while a WaitGroup's count can grow
	// dynamically via more Add calls even after some goroutines have
	// already registered and are running.

	// Backend connection: this fan-out/fan-in shape — launch several
	// workers, wait for all of them, then close the output channel —
	// is the standard way to run a batch of concurrent work (e.g.
	// several downstream calls) and collect every result exactly
	// once, without hardcoding how many there'll be.
}
