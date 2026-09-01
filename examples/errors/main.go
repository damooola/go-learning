// Program errors goes deeper into Go's error handling, building on
// the (value, error) return from divide() in functions/main.go.
package main

import (
	"errors"
	"fmt"
)

// error is a built-in interface — just one method. Anything with an
// Error() string method IS an error, the same implicit-satisfaction
// idea from interfaces/main.go.
//
//	type error interface {
//	    Error() string
//	}

// errors.New makes a simple error from a message. Package-level
// "sentinel" errors like this are meant to be compared against
// later with errors.Is, similar to well-known constants.
var ErrNotFound = errors.New("item not found")

func findItem(id int) (string, error) {
	if id != 1 {
		return "", ErrNotFound
	}
	return "widget", nil
}

// A custom error type carries structured data, not just a string —
// useful when the caller needs to react to specific details.
type ValidationError struct {
	Field string
	Issue string
}

// Implementing Error() string is all it takes to satisfy the error
// interface — no special keyword or declared relationship needed.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Issue)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Issue: "cannot be negative"}
	}
	return nil
}

// fmt.Errorf with %w wraps another error, keeping the original
// error reachable while adding context about where it happened.
func loadItem(id int) (string, error) {
	item, err := findItem(id)
	if err != nil {
		return "", fmt.Errorf("loadItem: %w", err)
	}
	return item, nil
}

func main() {
	// The idiomatic pattern: call, then immediately check err.
	item, err := findItem(1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("found:", item)
	}

	_, err = findItem(2)
	if err != nil {
		fmt.Println("error:", err)
	}

	// errors.Is checks whether an error IS (or wraps) a specific
	// sentinel error — safer than "err == ErrNotFound" once wrapping
	// is involved, because == would fail on a wrapped error.
	_, err = loadItem(2)
	fmt.Println("\nwrapped error:", err)
	fmt.Println("errors.Is(err, ErrNotFound):", errors.Is(err, ErrNotFound))

	// errors.As checks whether an error IS (or wraps) a specific
	// error TYPE, and if so, copies it into the target so you can
	// read its fields.
	err = validateAge(-5)
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("\nvalidation failed on field:", validationErr.Field)
	}

	// Dart comparison: Dart signals failure with exceptions
	// (throw/try/catch), which unwind the call stack automatically
	// until something catches them. Go has no exceptions for
	// expected failures — a function that can fail simply returns
	// an error value as normal data, and the CALLER decides whether
	// to handle it, propagate it, or wrap it with more context. Go
	// does have panic/recover for truly exceptional situations (a
	// programmer bug, not an expected failure), but reaching for it
	// like a Dart try/catch is considered bad Go style.

	// Backend connection: this return-and-check pattern is
	// everywhere in a Go backend — a handler calls a service, which
	// calls a repository, and errors get wrapped with fmt.Errorf at
	// each layer so the final log message shows the whole chain of
	// what was being attempted when something failed.
}
