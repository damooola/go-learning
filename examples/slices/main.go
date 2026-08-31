// Program slices demonstrates Go's slice type: a flexible-length
// view over an underlying array. This is what you'll actually use
// day to day, unlike the fixed-length arrays from the previous topic.
package main

import "fmt"

func main() {
	// A slice literal looks like an array literal but without a
	// length — Go manages the length for you.
	topics := []string{"variables", "functions", "loops"}
	fmt.Println("topics:", topics, "len:", len(topics))

	// append adds an element and returns the (possibly new) slice.
	// You must reassign it — append does not modify in place.
	topics = append(topics, "structs")
	fmt.Println("after append:", topics, "len:", len(topics))

	// The zero value of a slice is nil, not an empty slice, but
	// len() and append() both treat nil safely as length 0.
	var scores []int
	fmt.Println("\nnil slice:", scores, "len:", len(scores), "is nil:", scores == nil)
	scores = append(scores, 95)
	fmt.Println("after append to nil:", scores)

	// make(sliceType, length) creates a slice pre-filled with zero
	// values, useful when you know the size ahead of time.
	buffer := make([]int, 3)
	fmt.Println("\nmake([]int, 3):", buffer)

	// Slicing syntax s[low:high] takes a view from index low up to
	// (not including) high.
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("\nnumbers:", numbers)
	fmt.Println("numbers[1:3]:", numbers[1:3])
	fmt.Println("numbers[:2]:", numbers[:2])
	fmt.Println("numbers[3:]:", numbers[3:])

	// Dart comparison: Go's slice is the closest match to Dart's
	// growable List — both feel like a normal dynamic array day to
	// day. The gotcha: a slice expression like numbers[1:3] shares
	// the SAME underlying array as numbers, it doesn't copy. Editing
	// one can edit the other, until an append grows past capacity
	// and Go allocates a new backing array behind the scenes.
	view := numbers[1:3]
	view[0] = 999
	fmt.Println("\nnumbers after editing view:", numbers, "(numbers[1] changed too)")

	// Use copy() when you want an independent slice instead of a
	// shared view.
	independent := make([]int, len(numbers))
	copy(independent, numbers)
	independent[0] = -1
	fmt.Println("original numbers:", numbers, "(unchanged)")
	fmt.Println("independent copy:", independent, "(changed)")

	// Backend connection: functions that return rows from a database
	// query, items from a request body, or results from a search
	// almost always return a slice, e.g. []User or []Order.
}
