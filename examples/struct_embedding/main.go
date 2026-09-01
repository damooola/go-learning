// Program struct_embedding demonstrates Go's substitute for
// inheritance: embedding one struct inside another.
package main

import "fmt"

type Base struct {
	Name string
}

func (b Base) Describe() string {
	return "I am " + b.Name
}

// Employee "embeds" Base by naming its type with no field name.
// This is composition (Employee HAS a Base), but Go promotes Base's
// fields and methods so they're also reachable directly on Employee.
type Employee struct {
	Base
	Role string
}

// Describer is any type with a Describe() method — both Base and,
// thanks to promotion, Employee satisfy it automatically.
type Describer interface {
	Describe() string
}

func announce(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	employee := Employee{
		Base: Base{Name: "Ada"},
		Role: "Engineer",
	}

	// Name and Describe() are promoted from Base, so they're
	// reachable directly on employee without going through Base.
	fmt.Println("employee.Name:", employee.Name)
	fmt.Println("employee.Describe():", employee.Describe())

	// The embedded struct is still there under its type name if you
	// ever need to reach it explicitly.
	fmt.Println("employee.Base.Name:", employee.Base.Name)

	// Because Describe() was promoted, Employee satisfies Describer
	// even though Employee itself never declared a Describe method.
	announce(employee)

	// Important gotcha: this is NOT inheritance with virtual
	// methods. If Employee defines its OWN Describe(), calling it
	// on an Employee uses Employee's version — but Base's Describe
	// still only knows about Base, with no way to "see" Employee's
	// override. There's no dynamic dispatch like Dart's method
	// overriding gives you.
	manager := EmployeeWithOwnDescribe{Base: Base{Name: "Grace"}, Role: "Manager"}
	fmt.Println("\nmanager.Describe():", manager.Describe())

	// Dart comparison: Dart's "extends" creates a true is-a
	// relationship with polymorphism — an overridden method is used
	// even when called through code that only knows about the
	// parent class. Go's embedding is composition wearing
	// inheritance's clothes: you get field/method promotion for
	// convenience, but no virtual dispatch, and no is-a relationship
	// enforced by the type system (Employee is not actually a Base
	// as far as Go's type checker is concerned, e.g. you can't pass
	// an Employee where a Base is required).

	// Backend connection: embedding shows up for sharing common
	// fields across models (e.g. a base struct with ID/CreatedAt/
	// UpdatedAt embedded into every database model) and for
	// composing behavior, such as embedding sync.Mutex into a
	// struct to give it Lock()/Unlock() methods for free.
}

type EmployeeWithOwnDescribe struct {
	Base
	Role string
}

func (e EmployeeWithOwnDescribe) Describe() string {
	return fmt.Sprintf("I am %s, a %s", e.Name, e.Role)
}
