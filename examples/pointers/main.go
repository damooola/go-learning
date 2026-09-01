// Program pointers demonstrates Go's pointers: variables that hold
// the memory address of another variable.
package main

import "fmt"

// increment takes a plain int. Go passes arguments BY VALUE, so n
// here is a copy — changing it has no effect on the caller's variable.
func increment(n int) {
	n++
}

// incrementPointer takes a *int (a pointer to an int). Dereferencing
// it with *pointer reaches through to the original variable and
// modifies it directly.
func incrementPointer(pointer *int) {
	*pointer++
}

func main() {
	age := 30

	// & gives you the address of a variable — "a pointer to age".
	pointerToAge := &age
	fmt.Println("age:", age)
	fmt.Println("pointerToAge:", pointerToAge, "(a memory address)")

	// * in front of a pointer dereferences it — "the value pointer
	// points to".
	fmt.Println("*pointerToAge:", *pointerToAge)

	// Calling increment does NOT change age, because age was copied.
	increment(age)
	fmt.Println("\nafter increment(age):", age, "(unchanged)")

	// Calling incrementPointer DOES change age, because we handed it
	// age's address instead of a copy of its value.
	incrementPointer(&age)
	fmt.Println("after incrementPointer(&age):", age, "(changed)")

	// Dereferencing through the pointer can also assign a new value.
	*pointerToAge = 100
	fmt.Println("after *pointerToAge = 100:", age)

	// The zero value of a pointer is nil — it points to nothing yet.
	var missing *int
	fmt.Println("\nnil pointer:", missing)
	// fmt.Println(*missing) // this would panic: nil pointer dereference

	// Dart comparison: in Dart, every object variable is already a
	// reference under the hood, so passing one into a function lets
	// that function mutate the same object — you never see & or *.
	// Go is different: EVERYTHING is passed by value by default
	// (including whole structs, which get fully copied), and & / *
	// are how you explicitly opt in to reference-like behavior when
	// you want a function to modify the caller's data.

	// Backend connection: you'll see pointers constantly with
	// structs, e.g. func UpdateUser(u *User) so the function can
	// modify the actual User instead of a throwaway copy. Methods
	// (next major topic) are usually defined on pointer receivers
	// for the same reason.
}
