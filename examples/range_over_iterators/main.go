// Program range_over_iterators demonstrates ranging over a function
// (Go 1.23+), the mechanism behind range_builtin_types' earlier note
// that "range also works over custom iterator functions."
package main

import (
	"fmt"
	"iter"
)

// iter.Seq[V] is just an alias for func(yield func(V) bool) — a
// function that receives a "yield" callback and calls it once per
// value it wants to produce. Any function with this shape can be
// ranged over directly with "for v := range someIterator".
func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			// yield returns false if the caller stopped ranging
			// early (e.g. hit a break) — stop producing values too.
			if !yield(i) {
				return
			}
		}
	}
}

// Backward walks a slice back to front, without the caller having
// to write the reverse loop themselves every time.
func Backward[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(items) - 1; i >= 0; i-- {
			if !yield(items[i]) {
				return
			}
		}
	}
}

// iter.Seq2[K, V] is the two-value version — yield takes two
// arguments, which is what lets "for k, v := range ..." work.
func Enumerate[T any](items []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range items {
			if !yield(i, item) {
				return
			}
		}
	}
}

func main() {
	fmt.Println("Count(5):")
	for i := range Count(5) {
		fmt.Println(i)
	}

	fmt.Println("\nCount(10) with an early break:")
	for i := range Count(10) {
		if i == 3 {
			break // this makes the iterator's yield return false
		}
		fmt.Println(i)
	}

	words := []string{"a", "b", "c"}
	fmt.Println("\nBackward(words):")
	for w := range Backward(words) {
		fmt.Println(w)
	}

	fmt.Println("\nEnumerate(words):")
	for i, w := range Enumerate(words) {
		fmt.Printf("%d: %s\n", i, w)
	}

	// Dart comparison: Dart's Iterable/Iterator classes solve the
	// same problem — producing a sequence lazily, one value at a
	// time, usable in a for-in loop — but as an interface with
	// moveNext()/current you implement on a class. Go instead uses a
	// plain function with a callback (yield); "range" recognizes
	// that function shape and drives it for you. Conceptually
	// equivalent, syntactically quite different.

	// Backend connection: custom iterators are handy for streaming
	// results you don't want to materialize all at once — e.g.
	// paging through database rows or a large file line by line —
	// without loading everything into a slice up front.
}
