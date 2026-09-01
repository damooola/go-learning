// Program interfaces demonstrates Go's interfaces, building on the
// methods from methods/main.go.
package main

import (
	"fmt"
	"math"
)

// An interface lists a set of methods. Any type that has all of
// those methods automatically satisfies the interface — there is no
// "implements" keyword to write anywhere.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// describe accepts anything that satisfies Shape — it doesn't care
// whether it's actually a Rectangle, a Circle, or some other type
// written later that also happens to have Area() and Perimeter().
func describe(s Shape) {
	fmt.Printf("area: %.2f, perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	circle := Circle{Radius: 2}

	describe(rect)
	describe(circle)

	// This is why interfaces are useful: one slice can hold several
	// different concrete types, as long as they all satisfy Shape.
	shapes := []Shape{rect, circle}
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	fmt.Printf("\ntotal area: %.2f\n", total)

	// A type switch recovers the concrete type when you need
	// type-specific behavior instead of just the shared interface.
	fmt.Println("\ntype switch:")
	for _, shape := range shapes {
		switch value := shape.(type) {
		case Rectangle:
			fmt.Printf("a %vx%v rectangle\n", value.Width, value.Height)
		case Circle:
			fmt.Printf("a circle with radius %v\n", value.Radius)
		}
	}

	// any (an alias for the empty interface{}) has NO methods, so
	// every type satisfies it — it can hold literally anything, at
	// the cost of losing all type safety until you assert it back.
	var anything any
	anything = 42
	fmt.Println("\nanything:", anything)
	anything = "now a string"
	fmt.Println("anything:", anything)

	// Dart comparison: Dart requires "class Circle implements Shape"
	// — the relationship is explicit and checked where the class is
	// declared. Go's interfaces are implicit (structural typing,
	// sometimes called "duck typing"): a type satisfies an interface
	// just by having the right methods, with no declared connection
	// between them at all. This means you can define an interface
	// for someone else's existing type without touching it.

	// Backend connection: interfaces are how Go does dependency
	// injection and testing. A UserRepository interface might have
	// one real implementation backed by a database and one fake
	// implementation used only in tests — code that depends on the
	// interface doesn't need to know or care which one it got.
}
