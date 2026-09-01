// Program closures demonstrates functions that capture surrounding variables.
package main

import "fmt"

// makeCounter returns a function with its own remembered count.
func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// makeGreeter returns a configured function. The returned function captures
// prefix, so callers only need to supply a name.
func makeGreeter(prefix string) func(string) string {
	return func(name string) string {
		return prefix + ", " + name + "!"
	}
}

// transformAll accepts a function as an argument and applies it to every item.
// The transform parameter has the type func(int) int.
func transformAll(numbers []int, transform func(int) int) []int {
	results := make([]int, 0, len(numbers))
	for _, number := range numbers {
		results = append(results, transform(number))
	}
	return results
}

func main() {
	// The closure remembers count between calls.
	counterA := makeCounter()
	fmt.Println("counter A:", counterA())
	fmt.Println("counter A:", counterA())
	fmt.Println("counter A:", counterA())

	// A second closure has separate captured state.
	counterB := makeCounter()
	fmt.Println("counter B:", counterB())
	fmt.Println("counter A:", counterA())

	// These two functions share the same code but capture different prefixes.
	friendlyGreeting := makeGreeter("Hello")
	formalGreeting := makeGreeter("Good afternoon")
	fmt.Println(friendlyGreeting("Ada"))
	fmt.Println(formalGreeting("Grace"))

	// An inline closure can capture a value from main.
	multiplier := 3
	tripled := transformAll([]int{1, 2, 3, 4}, func(number int) int {
		return number * multiplier
	})
	fmt.Println("tripled:", tripled)
}
