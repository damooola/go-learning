// Program enums demonstrates Go's approach to enums. Go has no enum
// keyword — the idiom is a named type plus a block of constants.
package main

import "fmt"

// A named type based on int gives you a distinct type to restrict
// values to, even though under the hood it's just an int.
type OrderStatus int

// iota auto-increments starting at 0 for each line inside a const
// block, so this assigns Pending=0, Paid=1, Shipped=2, Cancelled=3
// without typing the numbers yourself.
const (
	Pending OrderStatus = iota
	Paid
	Shipped
	Cancelled
)

// Implementing String() makes OrderStatus satisfy the fmt.Stringer
// interface (the same implicit-interface idea from interfaces/main.go).
// fmt automatically calls it wherever OrderStatus is printed.
func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Paid:
		return "Paid"
	case Shipped:
		return "Shipped"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

func main() {
	status := Paid
	fmt.Println("status:", status)
	fmt.Printf("status: %v, underlying int: %d\n", status, status)

	// Because OrderStatus is its own type, passing a plain int where
	// one is expected is a compile error — this catches mistakes
	// that a bare "status int" field never would.
	var next OrderStatus = Shipped
	fmt.Println("next:", next)

	// A function can accept the named type instead of a loose int,
	// making invalid values much harder to pass in by accident.
	printReceipt(Cancelled)

	// iota also works for skipping values (use _ to skip a slot) or
	// for patterns like bit flags, though a simple sequence like
	// above covers most everyday cases.

	// Dart comparison: Dart has a real "enum" keyword
	// (enum OrderStatus { pending, paid, shipped, cancelled }) that
	// does this more directly, including built-in .name and
	// .values. Go's version is more manual — a type, a const block
	// with iota, and a hand-written String() method — but ends up
	// achieving the same goal: a small, closed set of named,
	// type-safe values.

	// Backend connection: this pattern shows up for anything with a
	// fixed set of states — order status, user role, request
	// method. The String() method matters for readable logs, and
	// you'll later see similar patterns for converting to/from JSON.
}

func printReceipt(status OrderStatus) {
	fmt.Println("receipt status:", status)
}
