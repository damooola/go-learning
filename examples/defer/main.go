// Program defer explains defer on its own — it's already been used
// throughout the concurrency lessons (defer wg.Done(), defer
// c.mu.Unlock()) without a dedicated explanation until now.
package main

import (
	"fmt"
	"os"
)

func main() {
	// defer schedules a call to run right before the surrounding
	// function returns — not immediately, no matter where the defer
	// statement itself sits in the function body.
	defer fmt.Println("this prints LAST, even though it's written first")
	fmt.Println("this prints first")

	// Multiple defers run in LIFO order — Last In, First Out. Think
	// of it as a stack: each defer statement pushes a call onto it,
	// and they run in reverse as the function returns.
	for i := 1; i <= 3; i++ {
		defer fmt.Println("deferred:", i)
	}
	fmt.Println("main body finishing")
	// Expect: "deferred: 3", "deferred: 2", "deferred: 1", THEN the
	// very first defer's message — it was registered first, so LIFO
	// runs it last.

	argGotcha()
	writeAndClean()
}

// argGotcha shows a common surprise: a deferred call's ARGUMENTS are
// evaluated immediately, when the defer statement runs — only the
// CALL itself is delayed.
func argGotcha() {
	x := 1
	defer fmt.Println("\ndeferred x was", x) // x captured as 1, right now
	x = 2
	fmt.Println("x is now", x)
	// Prints "x is now 2" first, then "deferred x was 1" — not 2.
}

// writeAndClean is the idiom defer exists for: guaranteeing a
// resource gets released on every return path, without repeating the
// cleanup call before each one.
func writeAndClean() error {
	f, err := os.CreateTemp("", "greeting-*.txt")
	if err != nil {
		return err
	}
	// Registered in this order: Remove, then Close. LIFO means Close
	// runs FIRST (last registered), Remove runs second — closing the
	// file before deleting it, which matters on some platforms.
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.WriteString("hello from the defer lesson\n"); err != nil {
		return err
	}

	fmt.Println("\nwrote and will clean up:", f.Name())
	return nil
	// defer's guarantee is what makes this safe even though there
	// are two different return points above — f.Close() and
	// os.Remove() run no matter which one fires.
}

// Dart comparison: defer is closest to Dart's try/finally — both
// guarantee cleanup code runs regardless of how the surrounding code
// exits. The difference is structural: try/finally wraps a block
// directly, while defer attaches to the nearest enclosing FUNCTION's
// return, wherever in that function the defer statement appears.

// Backend connection: defer is the standard way to close a database
// row set, an HTTP response body (resp.Body.Close()), or release a
// mutex (mutexes/main.go's defer c.mu.Unlock()) — anything acquired
// that must be released no matter which path the function takes out.
