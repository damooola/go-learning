// Program functions demonstrates how functions are declared and
// used in Go.
package main

import "fmt"

// Basic function: name, parameters with their types, then the
// return type after the parameter list.
func add(a int, b int) int {
	return a + b
}

// Consecutive parameters of the same type can share one type
// annotation: "a, b int" instead of "a int, b int".
func multiply(a, b int) int {
	return a * b
}

// Go functions can return more than one value. This has no direct
// Dart equivalent (Dart would use a class, a record, or a tuple
// package) — it's a core Go idiom.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

// Named return values declare the return variables up front and
// let a bare "return" send back whatever they currently hold.
// Handy for short functions, but can hurt readability in longer ones.
func minMax(numbers []int) (min, max int) {
	min, max = numbers[0], numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

// Variadic parameters (the ...) accept zero or more arguments,
// which arrive inside the function as a slice.
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Println("add(2, 3):", add(2, 3))
	fmt.Println("multiply(2, 3):", multiply(2, 3))

	// The idiomatic way to call a function returning (value, error)
	// is to check the error immediately. You'll see this pattern
	// everywhere in real Go backend code.
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("divide(10, 2):", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("divide(10, 0):", result)
	}

	low, high := minMax([]int{5, 2, 9, 1, 7})
	fmt.Println("\nminMax:", low, high)

	fmt.Println("\nsum():", sum())
	fmt.Println("sum(1, 2, 3):", sum(1, 2, 3))
	// A slice can be spread into a variadic parameter with "...".
	values := []int{4, 5, 6}
	fmt.Println("sum(values...):", sum(values...))

	// Functions are values in Go, same as in Dart — you can store
	// one in a variable, pass it as an argument, or write one
	// inline as an anonymous function (a "func literal").
	square := func(n int) int {
		return n * n
	}
	fmt.Println("\nsquare(5):", square(5))

	// Anonymous functions close over variables from the surrounding
	// scope, same as Dart closures.
	counter := makeCounter()
	fmt.Println("\ncounter():", counter())
	fmt.Println("counter():", counter())
	fmt.Println("counter():", counter())
}

// makeCounter returns a closure that remembers "count" between
// calls, because the returned function still references it.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
