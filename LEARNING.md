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
- [x] Close channels in depth, including their broadcast behaviour
- [x] Range over channels in more depth
- [x] Use `time.Timer` for one-shot delayed events
- [x] Use `time.Ticker` for repeating events
- [x] Build a worker pool from goroutines and channels
- [x] Use `sync.WaitGroup` to wait on a dynamic number of goroutines
- [x] Throttle work with rate limiting
- [ ] Share a counter safely with `sync/atomic` (current lesson)
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

## 10. Range over channels in more depth

`channels/main.go` already used `for value := range moreJobs` briefly. The
short version: `range` receives one value per iteration and keeps going until
the channel is both closed and drained, then the loop ends on its own — no
`ok` check or `break` needed, unlike the manual `for { value, ok := <-jobs
... }` loop it replaces.

```go
queue := make(chan string, 3)
go queueWork([]string{"wash", "dry", "fold"}, queue) // sends, then closes

for task := range queue {
	fmt.Println("-", task)
}
```

Two details worth calling out on their own:

- `range` over a channel gives you a single loop variable — the received
  value. There is no index, unlike `range` over a slice.
- If the producer never calls `close`, this loop hangs forever after the
  last real value — there's no other way for `range` to know the producer is
  finished. Forgetting to close a channel a `range` loop depends on is one of
  the most common goroutine leaks in Go.

This is close to Dart's `await for (var task in stream)` over a `Stream` —
both consume values one at a time until the source signals completion
(`Stream`'s "done" event, Go's channel close).

`range` over a channel is the natural way to drain a queue of work or a
stream of results from another goroutine — it's exactly what the worker
pool lesson uses inside each worker.

See [`examples/range_over_channels/main.go`](examples/range_over_channels/main.go).

Related reading:
[Go by Example: Range over Channels](https://gobyexample.com/range-over-channels)

### Try it yourself

1. Add a fourth item to the producer's list and confirm the loop still ends
   cleanly.
2. Comment out `close(queue)` inside `queueWork`, run the program, and
   observe it hangs — then restore the `close` call.
3. Change the sum-of-squares loop to also count how many values it received.
4. Rewrite the sum-of-squares loop as a manual `for { v, ok := <-numbers
   ...}` loop and confirm it produces the same total.

## 11. Timers

The timeouts lesson used `time.After(duration)` without explaining what it
actually is: shorthand for `time.NewTimer(duration).C`. `time.Timer` is the
full type behind it — a one-shot event that fires once after a delay, and
can be cancelled before it does.

```go
timer1 := time.NewTimer(150 * time.Millisecond)
<-timer1.C // blocks until the duration elapses, then receives the time
fmt.Println("timer1 fired")
```

`NewTimer` starts counting immediately. `timer.C` is a receive-only channel
that gets exactly one value once the duration has passed — that's the whole
mechanism `time.After` was built on, without a name attached to the timer
itself.

Because you have that name, you can stop it before it fires:

```go
timer2 := time.NewTimer(200 * time.Millisecond)
stopped := timer2.Stop() // true: prevented it from firing
```

`Stop()` returns `true` if it successfully prevented the firing, `false` if
the timer had already fired or already been stopped. Use `time.After` when
you only need the channel (as in the timeouts lesson); reach for
`time.NewTimer` directly when you might need to cancel it first.

Dart's closest match is `Timer(duration, callback)` — both schedule a
one-shot event and both can be cancelled before firing (`timer.cancel()` in
Dart, `timer.Stop()` in Go). The difference is what "firing" looks like:
Dart runs your callback directly, while Go delivers a value on a channel
that your own code then receives from.

Backend code reaches for timers directly (rather than `time.After`) whenever
a scheduled action might need to be cancelled if it turns out to be
unnecessary, e.g. debouncing repeated requests.

See [`examples/timers/main.go`](examples/timers/main.go).

Related reading: [Go by Example: Timers](https://gobyexample.com/timers)

### Try it yourself

1. Change timer1's duration and confirm the "fired" message is delayed to
   match.
2. Remove the `timer2.Stop()` call and confirm "timer2 fired" now prints
   from the goroutine.
3. Call `timer2.Stop()` a second time and print its return value — what do
   you expect, and does the result match?
4. Rewrite the timer1 wait using `select` with a `default` case (from the
   non-blocking channels lesson) so the program prints "still waiting"
   once before the timer fires, instead of blocking silently.

## 12. Tickers

A `time.Timer`'s channel fires exactly once. `time.Ticker` is the repeating
version: its channel fires again and again, once per interval, until you
stop it.

```go
ticker := time.NewTicker(100 * time.Millisecond)
done := make(chan bool)

go func() {
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("tick at", t.Format("15:04:05.000"))
		}
	}
}()

time.Sleep(550 * time.Millisecond)
ticker.Stop()
done <- true
```

The `select` inside the goroutine combines two ideas from earlier lessons:
waiting on multiple channels at once (select) and a dedicated signal channel
to ask a goroutine to stop (`done`, the same pattern as
`channel_synchronization/main.go`).

The important warning: forgetting `ticker.Stop()` leaks the ticker's
underlying goroutine and timer for as long as the program keeps running.
Unlike a `Timer`, which naturally produces no more events after its one
fire, a `Ticker` keeps firing forever if nothing stops it.

Dart's `Timer.periodic(duration, callback)` is the direct match — both fire
repeatedly on an interval, and both need an explicit cancel (`ticker.Stop()`
vs. `timer.cancel()`) or they run indefinitely.

Tickers drive periodic background work: polling a queue every few seconds,
refreshing a cache, emitting metrics on a schedule. The rate limiting lesson
next reuses this same "channel that fires on an interval" idea for a
different purpose — throttling instead of scheduling.

See [`examples/tickers/main.go`](examples/tickers/main.go).

Related reading: [Go by Example: Tickers](https://gobyexample.com/tickers)

### Try it yourself

1. Change the ticker's interval to 200ms and the sleep to 1 second, and
   predict how many ticks will print before you run it.
2. Remove `ticker.Stop()` (but keep `done <- true`) and confirm the
   goroutine still exits via the `done` case — then restore `Stop()`.
3. Add a counter that stops the ticker automatically after 3 ticks, instead
   of relying on `time.Sleep`.
4. Explain out loud why a `Ticker` needs a separate `done` channel to stop
   its goroutine, while a `Timer` used the same way would not.

## 13. Worker pools

A worker pool combines several earlier lessons into the pattern they were
building toward: a fixed number of goroutines all pulling jobs off one
shared channel and pushing results onto another.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // range_over_channels: keeps pulling until closed
		results <- j * 2
	}
}

jobs := make(chan int, 5)
results := make(chan int, 5)

for w := 1; w <= 3; w++ {
	go worker(w, jobs, results) // 3 workers sharing one jobs channel
}

for j := 1; j <= 5; j++ {
	jobs <- j
}
close(jobs) // lets every worker's range loop end (closing_channels: broadcast)
```

`jobs <-chan int` / `results chan<- int` are the direction restrictions from
`channel_direction/main.go`, now used to document each worker's role.
`range` over `jobs` is `range_over_channels/main.go`'s pattern, now shared by
several goroutines instead of one. Closing `jobs` relies on
`closing_channels/main.go`'s broadcast property: all 3 workers notice the
close once they've each drained the jobs meant for them, not just the first
one to notice.

Only 3 goroutines run at once here regardless of how many jobs are queued —
that's the actual point of the pattern. It bounds concurrency: e.g. never
running more than 3 downstream API calls or database queries at the same
time, no matter how much work is waiting, instead of spawning one goroutine
per job and letting concurrency grow unbounded.

Dart has no built-in worker pool primitive — the closest match is manually
spinning up several `Isolate`s with `SendPort`/`ReceivePort` pairs, which is
considerably more setup than a few goroutines sharing one channel. That gap
exists because Go's goroutines are so cheap to create (see
`goroutines/main.go`) that "just launch a few and share a channel" is
already lightweight.

See [`examples/worker_pools/main.go`](examples/worker_pools/main.go).

Related reading: [Go by Example: Worker Pools](https://gobyexample.com/worker-pools)

### Try it yourself

1. Change the worker count from 3 to 1 and observe the jobs now run one at a
   time instead of overlapping.
2. Increase `numJobs` to 9 and watch how the 3 workers split the extra work.
3. Make each worker sleep for a random duration instead of a fixed 50ms, and
   observe the finishing order becomes less predictable.
4. Add a `sync.WaitGroup` (from the next lesson) instead of a counted
   results loop, so the program doesn't need to know `numJobs` in advance.

## 14. WaitGroups in more depth

`goroutines/main.go` already used `sync.WaitGroup` for the basic case: launch
N goroutines, `Add(1)` before each one, `Done()` when each finishes, `Wait()`
blocks until every `Add` has a matching `Done`. This lesson looks at two
details that matter once WaitGroups get used for real work: how a worker
receives the WaitGroup, and how to close a channel only once every worker is
actually finished.

```go
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}

var wg sync.WaitGroup
for w := 1; w <= 3; w++ {
	wg.Add(1)             // registered BEFORE the goroutine starts
	go worker(w, jobs, results, &wg)
}
```

Two rules worth calling out explicitly:

- `worker` takes `*sync.WaitGroup` — a pointer, not `sync.WaitGroup` by
  value. A `WaitGroup` tracks internal counters that must never be copied;
  `go vet` will warn if you accidentally pass one by value. This is the same
  pointer-passing reasoning as pointer receivers in `methods/main.go`.
- `wg.Add(1)` runs in the loop, before `go worker(...)` — never inside
  `worker` itself. Goroutines aren't guaranteed to start in the order they're
  launched, so if `Add` ran inside the goroutine, `wg.Wait()` could
  potentially run before any `Add` happens at all, and return immediately
  having never actually waited for anything.

`worker_pools/main.go` collected exactly `numJobs` results because it
already knew that count in advance. Combining a `WaitGroup` with a second
goroutine removes that requirement entirely:

```go
go func() {
	wg.Wait()      // blocks until every worker has called Done()
	close(results) // only now is it safe to say "no more results are coming"
}()

for r := range results { // range_over_channels: no job count needed
	total += r
}
```

This fan-out/fan-in shape — launch several workers, wait for all of them,
then close the output channel — is the standard way to run a batch of
concurrent work and collect every result exactly once, without hardcoding
how many results to expect.

Dart's nearest match is `Future.wait([...])` — both wait for a group of
concurrent operations to finish. The difference: `Future.wait` takes a
known, fixed list of `Future`s up front, while a `WaitGroup`'s count can grow
dynamically through more `Add` calls even after some goroutines are already
running.

See [`examples/wait_groups/main.go`](examples/wait_groups/main.go).

Related reading: [Go by Example: WaitGroups](https://gobyexample.com/waitgroups)

### Try it yourself

1. Move `wg.Add(1)` inside `worker` (after `defer wg.Done()`) and run the
   program several times — with enough runs, watch for it finishing "too
   early" with an incomplete sum, then restore the original order.
2. Change `worker` to take `sync.WaitGroup` by value instead of
   `*sync.WaitGroup` and run `go vet ./...` to see the warning it produces.
3. Increase the worker count to 5 and `numJobs` to 20, and confirm the sum
   is still correct.
4. Add a fourth worker that's launched without a matching `wg.Add(1)` and
   observe `wg.Wait()` returns before that worker finishes.

## 15. Rate limiting

`tickers/main.go` used `time.Ticker` to run a repeating action. Rate
limiting reuses the exact same ticker, but for a different purpose:
throttling how fast a loop is *allowed* to proceed, regardless of how fast
the work it's processing actually arrives.

```go
limiter := time.NewTicker(200 * time.Millisecond)
defer limiter.Stop()

for req := range requests {
	<-limiter.C // blocks until the next tick — this IS the throttle
	fmt.Println("request", req, "at", time.Now())
}
```

Each iteration waits for the next tick before continuing, so at most one
request is processed every 200ms no matter how many are already queued in
`requests`.

Real traffic usually arrives in bursts rather than evenly spaced, so a
second pattern layers a small buffered channel — a **token bucket** — on
top:

```go
burstyLimiter := make(chan time.Time, 3) // 3 tokens available immediately
for i := 0; i < 3; i++ {
	burstyLimiter <- time.Now()
}

go func() {
	for t := range burstyTicker.C { // refills one token per tick
		burstyLimiter <- t
	}
}()
```

The first 3 requests can go through immediately (using the pre-filled
tokens); after that, requests fall back to waiting on the ticker's rate,
same as the basic version. This is the same shape you'll see in real
libraries, just built here from primitives you already have (`Ticker` + a
buffered channel).

Dart has no built-in rate limiter either — you'd hand-roll the same idea
with `Timer.periodic` refilling a token count, or reach for a package.
Nothing about the concept is Go-specific, only the specific primitives used
to build it.

Rate limiting protects a service two ways: capping outbound requests to a
third-party API that enforces its own limit, and protecting your own
service from being overwhelmed by a sudden spike of incoming requests.

See [`examples/rate_limiting/main.go`](examples/rate_limiting/main.go).

Related reading:
[Go by Example: Rate Limiting](https://gobyexample.com/rate-limiting)

### Try it yourself

1. Change the basic limiter's interval to 500ms and observe the requests
   spread out further apart.
2. Increase `burstyLimiter`'s buffer from 3 to 5 and confirm all 5 bursty
   requests now print almost immediately.
3. Add a 6th bursty request and predict, before running it, how long after
   the burst it will print.
4. Explain out loud why `burstyLimiter`'s refill goroutine needs `range` over
   `burstyTicker.C` instead of a single receive.

## 16. Atomic counters

Every concurrency lesson so far has used channels or a `WaitGroup` to
coordinate goroutines. `sync/atomic` solves a narrower, more direct problem:
letting many goroutines safely share a single number, with no channel and no
mutex.

```go
var ops atomic.Uint64
var wg sync.WaitGroup

for i := 0; i < 50; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := 0; c < 1000; c++ {
			ops.Add(1)
		}
	}()
}

wg.Wait()
fmt.Println("ops:", ops.Load()) // always exactly 50000
```

`atomic.Uint64` (not a plain `uint64`) is the entire point: its `Add` and
`Load` methods are guaranteed safe to call from many goroutines at once. A
plain `ops++` shared across 50 goroutines would be a **data race** — two
goroutines could both read the same old value before either writes back,
silently losing updates. Swap `atomic.Uint64`/`Add`/`Load` for a plain
`uint64`/`ops++`, and the final count becomes unpredictable and changes
between runs, usually landing a bit under 50000.

Dart rarely needs anything like this. Its single-threaded event loop means
only one piece of code ever runs at a time (`Isolate`s aside, and they don't
share memory anyway), so there's no equivalent of two pieces of code racing
to update the same variable. Go's goroutines genuinely run in parallel
across CPU cores and share memory directly — that's exactly the situation
atomic operations protect against.

Atomic counters are the lightweight option when the shared state is just a
number — request counts, active connections, simple metrics. For anything
more than a single number (several related fields, a map, a struct), a
`sync.Mutex` is the more general tool.

See [`examples/atomic_counters/main.go`](examples/atomic_counters/main.go).

Related reading:
[Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)

### Try it yourself

1. Uncomment the `unsafeOps`/`ops++` block, run the program several times,
   and note that the printed count is unreliable and varies between runs.
2. Run the unsafe version with `go run -race .` (the race detector) and read
   the race report it produces.
3. Change the goroutine count from 50 to 200 and the inner loop from 1000 to
   500, and confirm the atomic total is still exactly correct.
4. Replace `atomic.Uint64` with a `sync.Mutex`-protected plain `uint64`
   (`mu.Lock(); ops++; mu.Unlock()`) and confirm it's also always exactly
   50000 — two different tools solving the same problem.

## Immediate next topics

1. `context` for cancellation and deadlines
