# Learning Go

This is our evolving guide. Each topic gets:

- a short explanation here;
- a runnable example in `examples/`;
- a few small exercises to try.

Because you already know Dart, lessons may include short **Dart comparison**
notes. These comparisons are learning aids; we will also point out where the
two languages behave differently so Dart habits do not cause confusing Go
bugs.

## Our goal

The long-term goal is to understand and contribute safely to the
[`neocore`](https://github.com/stmfb/neocore) Go backend.

`neocore` is a large production application, so it is normal that its code is
hard to read as a beginner. We will not try to understand it all at once.
Instead, each lesson will teach one small idea and later connect that idea to a
pattern used in the backend.

### Learning path

- [x] Set up a Go module and run a program
- [x] Learn data types
- [ ] Learn variables, constants, and operators
- [x] Make decisions with `if` and `else`
- [ ] Make decisions with `switch` (current lesson)
- [x] Repeat work with `for`
- [ ] Write and call functions
- [ ] Work confidently with arrays, slices, and maps
- [ ] Model data with structs
- [ ] Understand pointers and methods
- [ ] Split code into packages
- [ ] Handle errors
- [ ] Use interfaces
- [ ] Write tests
- [ ] Learn HTTP requests, responses, and status codes
- [ ] Encode and decode JSON
- [ ] Build a small REST API
- [ ] Add middleware and request validation
- [ ] Store and retrieve data
- [ ] Learn `context` and basic concurrency
- [ ] Read a small feature in `neocore` from route to database

The later backend lessons will mirror the broad shape found in `neocore`:

```text
HTTP request -> route -> handler -> service -> database/integration
                    |          |
                  JSON     business rules
```

We will build a much smaller practice API before working directly in that
codebase.

## Learning resources

- [Go by Example](https://gobyexample.com/) — short, runnable examples for
  individual Go features.
- [Go Developer Roadmap](https://roadmap.sh/golang) — a broad checklist of
  topics used by Go developers.
- [Official Go documentation](https://go.dev/doc/) — the authoritative
  reference when we need more detail.

We will not blindly copy their order or examples. Go by Example is concise and
sometimes omits error handling to keep a demonstration small. Our local lessons
will add beginner-friendly explanations, exercises, and safe backend practices.

## Running the code

From the repository root:

```powershell
go run .
go run ./examples/data_types
go run ./examples/for_loops
go run ./examples/if_else
go run ./examples/switch
```

## 1. Data types

A data type tells Go what kind of value a variable holds and which operations
are valid for that value. Go is statically typed, so a variable's type is known
when the program is compiled.

### Common basic types

| Type | Example | Used for |
| --- | --- | --- |
| `string` | `"Ada"` | Text |
| `bool` | `true` | True/false values |
| `int` | `42` | Whole numbers |
| `float64` | `3.14` | Decimal numbers |
| `byte` | `'A'` | One byte; alias for `uint8` |
| `rune` | `'界'` | One Unicode code point; alias for `int32` |

Go also has specifically sized numbers such as `int8`, `int16`, `int32`,
`int64`, and unsigned versions such as `uint8` and `uint64`. For ordinary
whole numbers, start with `int`. For ordinary decimal numbers, start with
`float64`.

### Declaring variables

```go
var name string = "Ada" // explicit type
var age = 30            // Go infers int
score := 92.5            // short declaration; only inside functions
```

Use `const` for a value that should not change:

```go
const language = "Go"
```

### Zero values

Variables declared without a value receive a predictable default:

| Type | Zero value |
| --- | --- |
| Numbers | `0` |
| `string` | `""` |
| `bool` | `false` |
| Pointers, slices, maps, functions, interfaces, channels | `nil` |

```go
var count int      // 0
var title string   // ""
var finished bool  // false
```

### Converting types

Go does not automatically mix different numeric types. Convert explicitly:

```go
age := 30
ageAsDecimal := float64(age)
```

Converting a decimal number to `int` removes the fractional part; it does not
round it:

```go
n := int(3.9) // 3
```

### Composite types

These types hold or organize other values:

```go
numbers := [3]int{10, 20, 30}             // array: fixed length
colors := []string{"red", "blue"}         // slice: flexible length
ages := map[string]int{"Ada": 30}         // map: key/value pairs
person := struct { Name string }{"Grace"} // struct: grouped fields
```

Slices and maps are especially common. Structs are how Go programs model
things such as a user, product, or transaction.

### Printing a value and its type

With `fmt.Printf`:

```go
fmt.Printf("value=%v type=%T\n", age, age)
```

- `%v` prints a value in a general format.
- `%T` prints its type.
- `%q` prints quoted strings and characters.
- `%d` prints decimal integers.
- `%f` prints floating-point numbers.

See the complete runnable example in
[`examples/data_types/main.go`](examples/data_types/main.go).

Related reading:

- [Go by Example: Values](https://gobyexample.com/values)
- [Go by Example: Variables](https://gobyexample.com/variables)

### Try it yourself

1. Add variables for your name, age, height, and whether you are learning Go.
2. Print each variable's value and type.
3. Create a slice containing three topics you want to learn.
4. Add another topic with `append`.
5. Change `price := 19.95` into an `int` and observe the result.

## 2. Switch

A `switch` chooses one branch by comparing a value with several cases. It is
often clearer than a long chain of `if`/`else if` statements.

```go
day := "Saturday"

switch day {
case "Saturday", "Sunday":
	fmt.Println("Weekend")
case "Monday":
	fmt.Println("Start of the work week")
default:
	fmt.Println("Weekday")
}
```

Important differences from Dart:

- Go does not require parentheses around the value.
- A case can contain multiple comma-separated values.
- Go automatically stops after a matching case, so you do not write `break`.
- `fallthrough` exists, but it is explicit and rarely needed.
- `default` runs when no case matches.

Go can also use `switch` without a value. Each case is then a Boolean
condition, making it a tidy alternative to an `if`/`else if` chain:

```go
switch {
case score >= 90:
	fmt.Println("A")
case score >= 80:
	fmt.Println("B")
default:
	fmt.Println("Keep practising")
}
```

Backend programs often use switches for values such as HTTP status codes,
request methods, transaction states, or error categories.

See [`examples/switch/main.go`](examples/switch/main.go).

Related reading: [Go by Example: Switch](https://gobyexample.com/switch)

### Try it yourself

1. Change `day` and predict the output before running the program.
2. Add `http.StatusUnauthorized` to the HTTP status switch.
3. Add grade `D` for scores from 60 through 69.
4. Create a traffic-light switch for `"red"`, `"yellow"`, and `"green"`.

## Immediate next topics

1. Finish the `switch` exercises
2. Variables, constants, and operators
3. Functions
4. Arrays, slices, and maps in more depth
5. Structs and methods
