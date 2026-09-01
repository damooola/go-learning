// Program range_builtin_types demonstrates range over types other
// than slices and maps: strings and integers.
package main

import "fmt"

func main() {
	// Ranging over a string gives you a byte INDEX and a rune VALUE,
	// not a byte value. Go decodes the UTF-8 bytes for you and skips
	// straight to the next character each time.
	word := "héllo"
	fmt.Println("ranging over \"héllo\":")
	for index, character := range word {
		fmt.Printf("index %d: %q\n", index, character)
	}
	// Notice index jumps from 1 to 3: 'é' takes 2 bytes in UTF-8,
	// so the next character starts at byte position 3, not 2.

	fmt.Println("\nlen(\"héllo\"):", len(word), "(byte count, not character count)")

	// Dart comparison: Dart strings are UTF-16 and word[i] gives you
	// a code unit directly. Go strings are just raw UTF-8 bytes —
	// word[i] gives you a single byte, which can slice a multi-byte
	// character in half. range is what safely walks character by
	// character; runes are covered in more depth in the next topic.

	// Since Go 1.22, ranging over a plain integer counts from 0 up
	// to (not including) that number — a concise way to write a
	// simple counting loop.
	fmt.Println("\nrange over an int:")
	for i := range 3 {
		fmt.Println(i)
	}

	// Equivalent to the classic for loop:
	// for i := 0; i < 3; i++ { fmt.Println(i) }

	// When you don't need the counter at all, drop the variable.
	fmt.Println("\nrange over an int, count only:")
	count := 0
	for range 3 {
		count++
	}
	fmt.Println("looped", count, "times")

	// range also works over channels and (since Go 1.23) custom
	// iterator functions — those come later, alongside concurrency.
}
