// Program maps demonstrates Go's map type: unordered key/value pairs,
// similar to Dart's Map.
package main

import "fmt"

func main() {
	// A map literal: map[KeyType]ValueType{...}
	scores := map[string]int{"Ada": 95, "Grace": 98}
	fmt.Println("scores:", scores, "len:", len(scores))

	// Add or update a key with normal assignment.
	scores["Linus"] = 88
	scores["Ada"] = 97 // overwrites the existing value
	fmt.Println("after updates:", scores)

	// Reading a missing key does NOT panic — it returns the value
	// type's zero value, which can be misleading (0 could mean
	// "no such student" or "a real score of 0").
	fmt.Println("\nmissing key:", scores["Nobody"])

	// The "comma ok" idiom tells you whether the key actually
	// existed, which is the safe way to check.
	value, ok := scores["Nobody"]
	fmt.Println("value:", value, "ok:", ok)

	value, ok = scores["Ada"]
	fmt.Println("value:", value, "ok:", ok)

	// delete removes a key. Deleting a key that doesn't exist is
	// a harmless no-op, not an error.
	delete(scores, "Linus")
	fmt.Println("\nafter delete:", scores)

	// The zero value of a map is nil. You can read from a nil map
	// safely (it behaves like empty), but writing to one panics.
	var nilMap map[string]int
	fmt.Println("\nnil map read:", nilMap["anything"], "(safe)")
	// nilMap["key"] = 1 // this line would panic: assignment to entry in nil map

	// make(map[K]V) creates an empty, writable map.
	writable := make(map[string]int)
	writable["first"] = 1
	fmt.Println("writable map:", writable)

	// for range over a map gives key and value, but in randomized
	// order — Go deliberately does not guarantee iteration order.
	fmt.Println("\nfor range over scores:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Dart comparison: Go's map is the same idea as Dart's Map<K, V>,
	// but two differences matter: reading a missing Dart key returns
	// null (and Dart would complain at compile time if the value
	// type isn't nullable), while Go silently returns the zero value
	// — always use "comma ok" when a missing key is meaningfully
	// different from a zero value. Also, Dart's Map preserves
	// insertion order; Go's does not.

	// Backend connection: maps show up constantly for lookups by ID
	// (e.g. a cache of users keyed by user ID) and when decoding
	// loosely-structured JSON into map[string]interface{}.
}
