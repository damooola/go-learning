// Program sorting demonstrates Go's built-in sort package for the
// common case: sorting a slice of a type Go already knows how to
// compare (strings, ints, floats).
package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort.Strings/Ints/Float64s all sort IN PLACE — they don't
	// return a new slice, they mutate the one you pass in, similar
	// to append when it doesn't need to grow the underlying array.
	strs := []string{"cherry", "apple", "banana"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 9, 4}
	sort.Ints(ints)
	fmt.Println("ints:   ", ints)

	floats := []float64{3.2, 1.1, 2.9}
	sort.Float64s(floats)
	fmt.Println("floats: ", floats)

	// Each type also has an IsSorted check, useful for confirming a
	// slice is already sorted before doing something that requires it.
	fmt.Println("\nints already sorted:", sort.IntsAreSorted(ints))

	// sort.SearchInts does a binary search — MUCH faster than
	// scanning the whole slice — but it only works correctly on an
	// ALREADY-sorted slice; there's no way for it to check that for
	// you, so calling it on unsorted data gives a meaningless result.
	position := sort.SearchInts(ints, 9)
	fmt.Println("index of 9 in ints:", position)

	// Dart comparison: Dart's List has sort() as a method on the
	// list itself (strs.sort()), mutating in place just like Go's
	// sort.Strings. Go's version being a package-level function
	// instead of a method follows the same shape as len()/append() —
	// operations that work across many types live as functions, not
	// methods, since Go has no generic "any sortable list" method set
	// for its built-in types the way Dart's List<T> does.

	// Backend connection: sorting query results, leaderboard entries,
	// or log timestamps before display are common; sort.SearchInts'
	// binary search matters once a slice is large enough that a
	// linear scan would be noticeably slow. The next lesson covers
	// sorting slices of your OWN struct types, which these built-in
	// functions can't do.
}
