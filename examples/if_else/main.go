// Program if_else demonstrates Go's if/else statements.
package main

import "fmt"

func main() {
	age := 20

	// Same shape as Dart: no parentheses needed around the condition,
	// but the { } braces are required even for one-line bodies.
	if age >= 18 {
		fmt.Println("You're an adult.")
	} else {
		fmt.Println("You're a minor.")
	}

	// else if chains work like Dart too.
	score := 72
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Go lets you run a short statement before the condition, scoped
	// to the if/else blocks only. Common for checking a function's
	// second "ok" return value (e.g. map lookups) in one line.
	scores := map[string]int{"Ada": 95, "Grace": 98}
	if value, ok := scores["Ada"]; ok {
		fmt.Println("Ada's score:", value)
	} else {
		fmt.Println("Ada not found")
	}
	// value and ok only exist inside the if/else above — this would
	// be a compile error: fmt.Println(value)

	// Go has no ternary operator (cond ? a : b like Dart) — if/else
	// assigning to a variable is the idiomatic replacement.
	var status string
	if age >= 18 {
		status = "adult"
	} else {
		status = "minor"
	}
	fmt.Println("status:", status)
}
