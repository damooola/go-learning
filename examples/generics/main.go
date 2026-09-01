// Program generics demonstrates writing code that works across
// multiple types without repeating it or falling back to any.
package main

import "fmt"

// [T any] declares a type parameter T, constrained to "any" type.
// Map applies transform to every element and returns a new slice —
// this is the Go equivalent of Dart's list.map(...).toList().
func Map[T, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}

// Filter keeps only the elements where predicate returns true.
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// comparable is a built-in constraint meaning "supports == and !=" —
// needed here because the function compares target against items.
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

// A custom constraint lists which underlying types are allowed. The
// ~ means "int itself, or any named type whose underlying type is
// int" (same idea for float64).
type Number interface {
	~int | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Generic types work the same way — Stack holds a slice of whatever
// type T it was created with.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// SlicesIndex takes TWO type parameters: E is the element type, and
// S is constrained to "any slice whose elements are of type E"
// (~[]E), rather than being fixed to []E directly. That extra
// parameter lets it accept not just []string but any named type
// whose underlying type is []E too (e.g. type Names []string) — a
// plain "s []E" parameter would only accept the literal []E type.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// A generic type isn't limited to wrapping a slice like Stack does
// — List builds a linked list, where each node (element[T]) points
// to the next one instead of living in one contiguous slice.
type element[T any] struct {
	next *element[T]
	val  T
}

type LinkedList[T any] struct {
	head, tail *element[T]
}

func (lst *LinkedList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *LinkedList[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	// T and U don't have to match — here int comes in, string goes out.
	labels := Map(numbers, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Println("labels:", labels)

	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens:", evens)

	fmt.Println("\nContains(numbers, 3):", Contains(numbers, 3))
	fmt.Println("Contains(numbers, 99):", Contains(numbers, 99))

	fmt.Println("\nSum(ints):", Sum([]int{1, 2, 3}))
	fmt.Println("Sum(floats):", Sum([]float64{1.5, 2.5}))

	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	value, _ := intStack.Pop()
	fmt.Println("\npopped from intStack:", value)

	var stringStack Stack[string]
	stringStack.Push("a")
	stringStack.Push("b")
	value2, _ := stringStack.Pop()
	fmt.Println("popped from stringStack:", value2)

	words := []string{"foo", "bar", "zoo"}
	fmt.Println("\nSlicesIndex(words, \"zoo\"):", SlicesIndex(words, "zoo"))

	var linkedList LinkedList[int]
	linkedList.Push(10)
	linkedList.Push(13)
	linkedList.Push(23)
	fmt.Println("linkedList.AllElements():", linkedList.AllElements())

	// Dart comparison: this is the one topic where Dart already
	// gave you the mental model — List<T>, Map<K, V>, and your own
	// generic classes work the same way conceptually. Go's syntax
	// differs (square brackets after the name: Stack[T any] instead
	// of Stack<T>) and Go adds "constraints" as their own concept —
	// interfaces like Number or comparable that limit which types
	// are allowed, roughly like Dart's "T extends Something" bound.

	// Backend connection: generics are common for reusable
	// utilities — a generic Repository[T] pattern, pagination
	// helpers, or Map/Filter/Contains-style helpers like these,
	// instead of writing the same function once per model type.
}
