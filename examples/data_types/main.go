// Program data_types demonstrates Go's most commonly used data types.
package main

import "fmt"

func main() {
	// Go infers these types from their values.
	var num int = 56
	name := "Ada"
	age := 30
	height := 1.65
	isLearning := true
	initial := 'A'

	// Printf verbs (the %-codes) each format a value differently:
	//   %v  - "default" format, works for any type
	//   %q  - double-quoted string/rune, with escapes made visible
	//   %T  - the Go type of the value, not its value
	//   %d  - base-10 integer
	//   %f  - decimal float (%.2f rounds to 2 places after the point)
	fmt.Println("Basic types:")
	fmt.Printf("num: %v (%T)\n", num, num)

	fmt.Printf("name: %q (%T)\n", name, name)
	fmt.Printf("age: %v (%T)\n", age, age)
	fmt.Printf("height: %.2f (%T)\n", height, height)
	fmt.Printf("isLearning: %v (%T)\n", isLearning, isLearning)
	fmt.Printf("initial: %q (%T)\n", initial, initial)

	// Variables declared without values receive their type's zero value.
	var count int
	var message string
	var complete bool
	fmt.Println("\nZero values:")
	fmt.Printf("count=%d message=%q complete=%v\n", count, message, complete)

	// Numeric conversions must be explicit.
	ageAsDecimal := float64(age)
	fmt.Println("\nType conversion:")
	fmt.Printf("ageAsDecimal: %.1f (%T)\n", ageAsDecimal, ageAsDecimal)

	// Composite types organize collections of values.
	topics := []string{"variables", "functions", "slices"}
	topics = append(topics, "structs")

	scores := map[string]int{
		"Ada":   95,
		"Grace": 98,
	}

	// A struct groups named fields into one type, like an object with
	// no methods. Here it's declared inside main() just for this demo;
	// normally you'd declare it at package level so other functions
	// can use it too.
	type Person struct {
		Name string
		Age  int
	}

	student := Person{Name: "Ada", Age: 30}

	fmt.Println("\nComposite types:")
	fmt.Printf("topics: %v (%T)\n", topics, topics)
	fmt.Printf("scores: %v (%T)\n", scores, scores)
	// %+v is %v plus field names, useful for structs.
	fmt.Printf("student: %+v (%T)\n", student, student)
}
