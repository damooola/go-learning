// Program methods demonstrates functions attached to a type,
// building on structs/main.go.
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// A method is a function with a "receiver" between func and the
// method name — (r Rectangle) here. This is a VALUE receiver: r is
// a copy of the Rectangle the method was called on, so Area cannot
// change the original. That's fine here since Area only reads.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale uses a POINTER receiver (*Rectangle), so r refers to the
// original Rectangle, not a copy. Use a pointer receiver whenever
// a method needs to modify the receiver.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

func (p *Person) HaveBirthday() {
	p.Age++
}

// A method can be attached to any named type, not just a struct —
// here Celsius is just a float64 with a name of its own.
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("rect:", rect)
	fmt.Println("rect.Area():", rect.Area())

	// Calling a pointer-receiver method on an addressable value
	// (like the local variable rect) works without writing &rect
	// yourself — Go inserts the & automatically.
	rect.Scale(2)
	fmt.Println("rect after Scale(2):", rect)

	ada := Person{Name: "Ada", Age: 30}
	fmt.Println("\n" + ada.Greet())
	ada.HaveBirthday()
	fmt.Println("ada after HaveBirthday():", ada)

	bodyTemp := Celsius(37)
	fmt.Printf("\n%.0f°C is %.1f°F\n", float64(bodyTemp), bodyTemp.ToFahrenheit())

	// Dart comparison: Go has no class keyword, so instead of
	// bundling fields and methods inside one class body, Go declares
	// the struct's fields separately from its methods (methods can
	// even live in a different file). The one real decision Dart
	// doesn't force on you: value receiver (works on a copy, like
	// passing the struct normally) vs pointer receiver (works on the
	// original, like passing a reference). Rule of thumb: use a
	// pointer receiver if the method mutates the receiver, or if the
	// struct is large enough that copying it would be wasteful.

	// Backend connection: this is exactly the shape of methods
	// you'll see in a real backend, e.g. (u *User) Save() or
	// (o Order) Total() — pointer receivers for anything that
	// changes state or hits a database, value receivers for small
	// read-only calculations.
}
