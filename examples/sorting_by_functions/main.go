// Program sorting_by_functions covers sorting a slice by a custom
// rule — something sorting/main.go's sort.Strings/Ints/Float64s can't
// do, since they only know how to compare the type they're named for.
package main

import (
	"fmt"
	"sort"
)

// byLength implements sort.Interface: Len, Less, and Swap. Any type
// with these three methods can be handed to sort.Sort — this is the
// classic, pre-sort.Slice way to sort by a custom rule.
type byLength []string

func (s byLength) Len() int      { return len(s) }
func (s byLength) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

type person struct {
	name string
	age  int
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// byLength(fruits) doesn't copy the slice — it's the SAME
	// underlying array, just viewed through a named type that has
	// Len/Less/Swap attached. Sorting through that view sorts fruits
	// itself, same underlying-array sharing idea as slicing in
	// slices/main.go.
	sort.Sort(byLength(fruits))
	fmt.Println("fruits by length:", fruits)

	people := []person{
		{"Ada", 36},
		{"Grace", 85},
		{"Linus", 54},
	}

	// sort.Slice is the shorter modern alternative: instead of
	// defining a whole named type with 3 methods, pass a "less"
	// function directly. No Len/Swap to write — sort.Slice handles
	// those internally.
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("\npeople by age:")
	for _, p := range people {
		fmt.Printf("- %s (%d)\n", p.name, p.age)
	}

	// sort.Slice does not guarantee STABILITY: elements that compare
	// equal can still end up reordered relative to each other.
	// sort.SliceStable preserves their original relative order when
	// two elements compare equal — matters when a later sort should
	// only break ties, not scramble them.
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println("\npeople by name (stable):")
	for _, p := range people {
		fmt.Printf("- %s (%d)\n", p.name, p.age)
	}

	// Dart comparison: Dart's list.sort((a, b) => a.age - b.age)
	// takes a comparator the same way sort.Slice takes a "less"
	// function — this is the closer everyday match. sort.Interface's
	// Len/Less/Swap trio has no real Dart equivalent; it predates
	// sort.Slice (which needs generics-adjacent reflection tricks
	// under the hood) as Go's only way to customize sort behavior.

	// Backend connection: sorting a slice of your own structs — users
	// by signup date, transactions by amount, results by relevance
	// score — is one of the most common uses of sort.Slice in
	// backend code. Reach for sort.Interface only when you want a
	// named, reusable sort order rather than a one-off closure.
}
