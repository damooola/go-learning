# Learning Go

This is our evolving guide. Each topic gets:

- a short explanation here;
- a runnable example in `examples/`;
- a few small exercises to try.

## Running the code

From the repository root:

```powershell
go run .
go run ./examples/data_types
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

### Try it yourself

1. Add variables for your name, age, height, and whether you are learning Go.
2. Print each variable's value and type.
3. Create a slice containing three topics you want to learn.
4. Add another topic with `append`.
5. Change `price := 19.95` into an `int` and observe the result.

## Suggested next topics

1. Variables, constants, and operators
2. `if`, `switch`, and loops
3. Functions
4. Arrays, slices, and maps in more depth
5. Structs and methods
6. Pointers
7. Interfaces and error handling

