// Program switch demonstrates Go's switch statement.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// A basic switch compares one value against each case.
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		// Multiple values can share the same case.
		fmt.Println(day, "is part of the weekend")
	case "Monday":
		fmt.Println("Monday starts the work week")
	default:
		fmt.Println(day, "is a weekday")
	}

	// Dart comparison: Go stops after the matching case automatically.
	// There is no need to add break at the end of every case.

	// A switch without a value checks Boolean conditions from top to bottom.
	// The first true case runs.
	score := 84
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// Backend connection: HTTP handlers frequently work with status codes.
	status := http.StatusNotFound
	switch status {
	case http.StatusOK:
		fmt.Println("Request succeeded")
	case http.StatusBadRequest:
		fmt.Println("The client sent an invalid request")
	case http.StatusNotFound:
		fmt.Println("The requested resource was not found")
	default:
		fmt.Println("Received HTTP status:", status)
	}
}
