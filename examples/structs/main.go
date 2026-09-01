// Program structs goes deeper into Go structs, building on the
// simple struct from data_types/main.go.
package main

import "fmt"

// A struct is usually declared at package level, so every function
// in the package can use it — unlike data_types/main.go, which
// declared Person inside main() just for that one demo.
type Person struct {
	Name string
	Age  int
}

// A function that builds and returns a struct is a common
// replacement for what other languages do with a constructor. Go
// has no constructor keyword.
func newPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

// newPersonPointer returns a *Person instead. Returning a pointer
// avoids copying the whole struct and lets callers mutate the
// original — the same tradeoff covered in the pointers topic.
func newPersonPointer(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// haveBirthday takes a *Person so it can modify the caller's data.
// If this took a plain Person, it would only modify a throwaway copy.
func haveBirthday(p *Person) {
	p.Age++
}

func main() {
	// Field names make it clear which value goes where, and let you
	// skip fields (they get their type's zero value).
	ada := Person{Name: "Ada", Age: 30}
	fmt.Println("ada:", ada)

	// Fields without names must be given in declaration order —
	// fragile if the struct's fields ever get reordered, so prefer
	// named fields.
	grace := Person{"Grace", 34}
	fmt.Println("grace:", grace)

	// The zero value of a struct fills every field with its own
	// zero value.
	var nobody Person
	fmt.Println("zero-value Person:", nobody)

	// Access fields with dot notation.
	fmt.Println("\nada.Name:", ada.Name)
	ada.Age = 31
	fmt.Println("ada after birthday:", ada)

	fmt.Println("\nnewPerson(...):", newPerson("Linus", 55))

	// Dot notation works the same way on a pointer to a struct — Go
	// automatically dereferences it for you, no special arrow
	// syntax needed (unlike C's ->).
	gracePointer := newPersonPointer("Grace", 34)
	fmt.Println("gracePointer.Name:", gracePointer.Name)

	// This works because haveBirthday receives a pointer, so it
	// modifies the original struct rather than a copy.
	haveBirthday(gracePointer)
	fmt.Println("gracePointer after birthday:", *gracePointer)

	// Structs are value types, same as arrays: assigning one struct
	// variable to another copies every field.
	adaCopy := ada
	adaCopy.Name = "Ada Copy"
	fmt.Println("\nada:", ada, "(unchanged)")
	fmt.Println("adaCopy:", adaCopy, "(changed)")

	// Dart comparison: a Go struct is closest to a small Dart class
	// with only fields and no behavior. Two differences matter:
	// Go has no constructors (a plain function returning the struct
	// is the idiom) and no inheritance — Go favors composition
	// (embedding one struct inside another) instead, covered later.
	// Also remember: unlike a Dart object, a Go struct is copied on
	// assignment unless you're holding a pointer to it.

	// Backend connection: structs are how you model everything in a
	// backend — a User, an Order, a request body, a database row.
	// Methods (functions attached to a struct, using the *Person
	// syntax from haveBirthday) are the very next topic.
}
