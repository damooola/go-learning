// Program recursion demonstrates functions that call themselves.
package main

import "fmt"

// countdown prints n and then calls itself with a smaller number.
func countdown(n int) {
	if n <= 0 {
		// Base case: stop making recursive calls.
		fmt.Println("Go!")
		return
	}

	fmt.Println(n)
	countdown(n - 1)
}

// factorial calculates n * (n-1) * ... * 1.
// This lesson calls it only with non-negative numbers.
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// sumSlice handles one value and asks another call to handle the rest.
func sumSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + sumSlice(numbers[1:])
}

func main() {
	fmt.Println("Countdown:")
	countdown(3)

	fmt.Println("\nFactorial:")
	fmt.Println("5! =", factorial(5))

	fmt.Println("\nRecursive slice sum:")
	numbers := []int{10, 20, 30, 40}
	fmt.Printf("sum of %v = %d\n", numbers, sumSlice(numbers))

	// A loop is often clearer for simple repetition in Go.
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("same sum using a loop =", total)
}
