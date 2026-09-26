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
- [x] Make decisions with `switch`
- [x] Repeat work with `for`
- [x] Write and call functions
- [x] Understand closures
- [x] Understand recursion
- [x] Work with arrays, slices, and maps
- [x] Model data with structs
- [x] Understand pointers and methods
- [ ] Split code into packages
- [x] Handle errors
- [x] Use interfaces
- [x] Learn struct embedding, enums, and generics
- [ ] Write tests
- [ ] Learn HTTP requests, responses, and status codes
- [ ] Encode and decode JSON
- [ ] Build a small REST API
- [ ] Add middleware and request validation
- [ ] Store and retrieve data
- [x] Start concurrency with goroutines and channels
- [x] Learn channel buffering and synchronization
- [x] Restrict channels with channel directions
- [x] Wait on multiple channels with `select`
- [x] Add timeouts to channel operations
- [x] Use non-blocking channel operations with `select`/`default`
- [ ] Close channels in depth, including their broadcast behaviour (current
      lesson)
- [ ] Learn `context`
- [ ] Read a small feature in `neocore` from route to database

The later backend lessons will mirror the broad shape found in `neocore`:

```text
HTTP request -> route -> handler -> service -> database/integration
                    |          |
                  JSON     business rules
```

We will build a much smaller practice API before working directly in that
codebase.

## Neocore case studies

- [Bank-statement HTTP handlers](case_studies/bank_statement_handler/README.md)
  — an annotated 533-line production file covering handler closures, request
  validation, pointers, structs, slices, maps, methods, errors, switch, JSON
  responses, and your connector-fallback change.

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
# Replace <topic> with a folder name, such as channels.
go run ./examples/<topic>

# Compile every lesson at once.
go test ./...
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

## 3. Closures

A closure is a function that remembers variables from the surrounding scope,
even after that surrounding function has finished running.

```go
func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
```

`makeCounter` returns a value whose type is `func() int`: a function taking no
arguments and returning an `int`. The returned function captures `count`, so
the value remains available and changes on every call.

```go
counter := makeCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
```

This is similar to a Dart function returning another function:

```dart
int Function() makeCounter() {
  var count = 0;
  return () => ++count;
}
```

Each call to `makeCounter()` creates a separate captured `count`. Closures can
also be passed into other functions, allowing callers to provide a small piece
of behaviour.

Backend code often uses closures for configured handlers, middleware, logging,
validation, and callbacks. A closure that modifies captured state is not
automatically safe when multiple goroutines call it; we will handle that later
when learning concurrency.

See [`examples/closures/main.go`](examples/closures/main.go).

Related reading: [Go by Example: Closures](https://gobyexample.com/closures)

### Try it yourself

1. Change the counter so it increases by `2`.
2. Write `makeCountdown(start int)` that returns `start`, then one less each
   time it is called.
3. Create two counters and confirm that changing one does not change the other.
4. Pass a closure to `transformAll` that triples every number.

## 4. Recursion

Recursion happens when a function calls itself. Every recursive function needs
two important parts:

1. A **base case** that stops the recursion.
2. A **recursive step** that moves the input closer to the base case.

```go
func factorial(n int) int {
	if n == 0 { // base case
		return 1
	}
	return n * factorial(n-1) // recursive step
}
```

For `factorial(3)`, Go builds up the calls and then returns through them:

```text
factorial(3)
3 * factorial(2)
3 * 2 * factorial(1)
3 * 2 * 1 * factorial(0)
3 * 2 * 1 * 1 = 6
```

The idea is the same in Dart. The main differences are Go's function syntax
and explicit types:

```dart
int factorial(int n) {
  if (n == 0) return 1;
  return n * factorial(n - 1);
}
```

A missing or unreachable base case causes calls to continue until the program
runs out of stack space. Deep recursion can also consume a lot of memory. For
ordinary repetition, a `for` loop is usually simpler in Go. Recursion is most
helpful when the data is naturally nested, such as trees, folder structures,
menus, and parent/child records.

See [`examples/recursion/main.go`](examples/recursion/main.go).

Related reading: [Go by Example: Recursion](https://gobyexample.com/recursion)

### Try it yourself

1. Change `countdown` to start at `10`.
2. Trace `factorial(4)` on paper before running it.
3. Write `sumTo(n int)` that calculates `n + (n-1) + ... + 1`.
4. Add a base case that makes `sumSlice` work with an empty slice.

## 5. Channel directions

A normal `chan T` can both send and receive values. A function often needs
only one of those abilities, so Go lets its parameter type restrict the
channel's direction:

```go
func send(ch chan<- string) { // can only send
	ch <- "hello"
}

func receive(ch <-chan string) { // can only receive
	fmt.Println(<-ch)
}
```

Use the arrow as a visual guide:

- `chan<- T`: the arrow points into the channel, so it is send-only.
- `<-chan T`: the arrow points out of the channel, so it is receive-only.
- `chan T`: no restriction; it can send and receive.

`make` creates a bidirectional channel. Go automatically narrows it when it is
passed to a function expecting a send-only or receive-only channel. Trying the
wrong operation inside that function is a compile-time error.

Dart does not have directly equivalent channel type syntax. `SendPort` and
`ReceivePort` separate responsibilities in isolate communication, but Go can
restrict two views of the same channel through function parameter types.

Direction restrictions do not change buffering or blocking. They document who
owns each operation and let the compiler catch mistakes. As a common rule, the
producer that sends values should also close the channel.

See [`examples/channel_direction/main.go`](examples/channel_direction/main.go).

Related reading:
[Go by Example: Channel Directions](https://gobyexample.com/channel-directions)

### Try it yourself

1. Add another number to the producer.
2. Change the multiplier from `2` to `3`.
3. Try receiving inside `produce`, observe the compiler error, and undo it.
4. Try sending inside `multiply`, observe the compiler error, and undo it.

## 6. Select

A `select` waits for channel operations. It looks like `switch`, but its cases
are sends or receives instead of ordinary values:

```go
select {
case message := <-channelA:
	fmt.Println("A:", message)
case message := <-channelB:
	fmt.Println("B:", message)
}
```

If neither channel is ready, `select` blocks until one becomes ready. If
several cases are ready at the same time, Go chooses one; do not rely on a
specific ordering. A single `select` runs only one case, so put it in a loop
when you need to receive multiple results.

The closest everyday Dart analogy is `Future.any`, because both can react to
whichever operation finishes first. The match is not exact: `select` operates
on channel sends and receives, can be repeated, and does not cancel the other
operations.

Backend services use `select` to combine results from concurrent work, wait
for messages, handle cancellation, and enforce timeouts. We will add timeout
and cancellation cases in later lessons.

See [`examples/select/main.go`](examples/select/main.go).

Related reading: [Go by Example: Select](https://gobyexample.com/select)

### Try it yourself

1. Reverse the two delays and predict which message arrives first.
2. Give both goroutines the same delay and run the program several times.
3. Add a third channel and a third `select` case.
4. Change the loop so it waits for three total messages.

## 7. Timeouts

A timeout prevents a program from waiting forever. Combine `select` with
`time.After` to wait for either a result or a deadline:

```go
select {
case result := <-results:
	fmt.Println("received:", result)
case <-time.After(100 * time.Millisecond):
	fmt.Println("timed out")
}
```

`time.After(duration)` returns a receive-only channel. After the duration has
passed, that channel receives a time value and its `select` case becomes ready.
Whichever case becomes ready first wins.

This resembles Dart's `Future.timeout`, although Go expresses the race using
channels and `select`.

Important: a timeout stops the caller from **waiting**, but it does not cancel
the goroutine doing the work. The example uses a result channel with capacity
one so a late worker can still send its result and exit. In real backend code,
`context` is commonly used when the underlying operation must also be
cancelled.

Timeouts protect backend handlers when databases, APIs, queues, or other
services respond too slowly.

See [`examples/timeouts/main.go`](examples/timeouts/main.go).

Related reading: [Go by Example: Timeouts](https://gobyexample.com/timeouts)

### Try it yourself

1. Increase the first timeout so its operation succeeds.
2. Decrease the second timeout so it fails.
3. Print the value received from the `time.After` channel.
4. Add a third operation whose delay and timeout are equal, then run it several
   times and observe why relying on exact timing is risky.

## 8. Non-blocking channel operations

A plain send or receive blocks until the other side is ready. Adding a
`default` case to `select` lets you attempt a channel operation and fall
through immediately if it is not ready, instead of waiting:

```go
select {
case msg := <-messages:
	fmt.Println("received message:", msg)
default:
	fmt.Println("no message received")
}
```

Because nothing is sending on `messages`, the receive case is never ready, so
`select` runs `default` right away rather than blocking. The same pattern
works for a non-blocking send, and several non-blocking cases can share one
`default`; `select` still only runs one case per attempt.

This differs from the timeouts lesson: a timeout (`select` with
`time.After`) waits up to a bounded duration before giving up, while a
non-blocking `select` waits zero time. It is for checking something — a
shutdown signal, an optional message — once per loop iteration without ever
stalling that iteration on the check.

Dart has no direct equivalent. `Stream` is push/callback based: you subscribe
and get notified later, but there is no built-in "is a value available right
now, otherwise move on" check. Approximating this in Dart would mean
maintaining your own flag or buffer alongside the stream.

Backend services use this pattern to poll for a signal or a piece of optional
work alongside a larger loop, without ever pausing that loop to wait for it.

See [`examples/non_blocking_channels/main.go`](examples/non_blocking_channels/main.go).

Related reading:
[Go by Example: Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations)

### Try it yourself

1. Start a goroutine that sends on `messages` after a short delay, then run
   the first `select` again and observe it now receives instead of hitting
   `default`.
2. Add a `count` variable, and inside a `for` loop of several iterations, use
   a non-blocking receive on `messages` to print whenever one arrives while
   still printing something every iteration.
3. Change the combined `select` so `signals` is ready instead of `messages`,
   and confirm the `signals` case runs.
4. Remove `default` from one of the selects and explain out loud why the
   program now blocks forever.

## 9. Closing channels in depth

`channels/main.go` already showed the basics: `close(ch)` means "no more
values are coming," and receiving from a closed, drained channel returns the
zero value with `ok` set to `false` instead of blocking forever. What that
lesson didn't cover is what makes `close` different from an ordinary send.

A sent value only reaches **one** receiver — whichever goroutine happens to
be ready for it. Closing a channel is a **broadcast**: every goroutine
currently blocked on a receive from that channel wakes up at the same
moment, and every future receive on it keeps returning immediately too.

```go
ready := make(chan bool)

for id := 1; id <= 3; id++ {
	go func(id int) {
		<-ready // blocks until ready is closed
		fmt.Printf("worker %d: starting\n", id)
	}(id)
}

close(ready) // releases all three goroutines at once
```

Closing is also safe to do while another goroutine is mid-loop receiving from
that same channel — the pending values already sent are still delivered
first, and only once those are drained does `ok` become `false`. Receiving
again after that keeps returning `ok=false`; it never panics from repeating
the receive. (Closing an already-closed channel, or sending to a closed
channel, are the two operations that *do* panic.)

Dart's closest match is `StreamController.close()`: both notify every current
subscriber that no more values are coming. Dart delivers that as a stream
"done" callback; Go represents it as an ordinary channel receive that comes
back with `ok=false`.

The broadcast property is what makes `close` useful as a stop/cancel signal
for an entire worker pool — closing one shared channel tells every worker
goroutine to stop at once, which is simpler than sending each of them an
individual value.

See [`examples/closing_channels/main.go`](examples/closing_channels/main.go).

Related reading:
[Go by Example: Closing Channels](https://gobyexample.com/closing-channels)

### Try it yourself

1. Change the worker count from 3 to 5 and confirm all five still start
   together right after `close(ready)`.
2. Replace `close(ready)` with sending `true` on `ready` instead, and observe
   that only one worker starts while the others block forever (a deadlock —
   stop the program manually).
3. Add a second `_, ok := <-jobs` after the first one and confirm `ok` is
   still `false`, proving it's safe to repeat.
4. Try `close(jobs)` a second time and read the panic message it produces.

## Immediate next topics

1. Worker pools and wait groups
2. `context` for cancellation and deadlines
