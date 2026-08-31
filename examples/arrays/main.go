// Program arrays demonstrates Go's array type: a fixed-length
// sequence of values of the same type.
package main

import "fmt"

func main() {
	// The length is part of the type: [3]int and [5]int are
	// different types. You cannot append to an array or resize it.
	var numbers [3]int
	fmt.Println("zero-valued array:", numbers)

	// Set values by index, same as Dart's List.
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println("filled array:", numbers)

	// Declare and fill in one line with an array literal.
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes, "length:", len(primes))

	// [...] lets Go count the elements for you instead of writing
	// the number yourself.
	colors := [...]string{"red", "green", "blue"}
	fmt.Println("colors:", colors, "length:", len(colors))

	// for range walks an array just like it walks a slice.
	fmt.Println("\nlisting colors:")
	for index, color := range colors {
		fmt.Printf("%d: %s\n", index, color)
	}

	// Dart comparison: Dart's List is always growable and is a
	// reference type — copying a List variable copies the reference,
	// so both variables see the same underlying data. A Go array is
	// a value type: copying it (or passing it to a function) copies
	// every element into a brand new array.
	original := [3]int{1, 2, 3}
	copyOfOriginal := original
	copyOfOriginal[0] = 999
	fmt.Println("\noriginal:", original, "(unchanged)")
	fmt.Println("copy:", copyOfOriginal, "(changed)")

	// Backend connection: plain arrays are rare in real Go backend
	// code because the fixed length is inconvenient — you usually
	// don't know how many items you'll have (rows from a database,
	// items in a request body, etc.). Slices, which are covered
	// next, solve that and are what you'll actually use day to day.
}
