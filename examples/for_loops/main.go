// Program for_loops demonstrates Go's for loop, the only looping
// keyword in the language (no while/do-while — for covers all of it).
package main

import "fmt"

func main() {
	// Classic three-part for: init; condition; post.
	fmt.Println("Classic for:")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Drop the init and post parts and it behaves like a while loop.
	fmt.Println("\nWhile-style for:")
	n := 3
	for n > 0 {
		fmt.Println(n)
		n--
	}

	// Drop the condition too and it loops forever, so you need
	// break to escape it.
	fmt.Println("\nInfinite for with break:")
	count := 0
	for {
		if count == 3 {
			break
		}
		fmt.Println(count)
		count++
	}

	// continue skips to the next iteration without exiting the loop.
	fmt.Println("\nfor with continue (skip even numbers):")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// for range walks a slice, giving you the index and the value.
	fmt.Println("\nfor range over a slice:")
	topics := []string{"variables", "functions", "loops"}
	for index, topic := range topics {
		fmt.Printf("%d: %s\n", index, topic)
	}

	// Use _ to ignore the index when you only need the value.
	fmt.Println("\nfor range, values only:")
	for _, topic := range topics {
		fmt.Println(topic)
	}

	// for range over a map gives you key and value, in no
	// guaranteed order (Go randomizes map iteration order on purpose).
	fmt.Println("\nfor range over a map:")
	scores := map[string]int{"Ada": 95, "Grace": 98}
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
}
