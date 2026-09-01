// Program strings_and_runes goes deeper into how Go strings are
// actually stored, building on the string/range basics from
// range_builtin_types.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const word = "résumé"

	// A for loop can count a string's characters manually by
	// ranging over it and counting iterations — range gives one
	// iteration per rune, not per byte, so this matches
	// utf8.RuneCountInString rather than len().
	characterCount := 0
	for range word {
		characterCount++
	}
	fmt.Println("character count via for range:", characterCount)

	// A plain byte-indexed for loop instead counts bytes, matching
	// len(word) — included here just to show the contrast.
	byteCount := 0
	for i := 0; i < len(word); i++ {
		byteCount++
	}
	fmt.Println("byte count via classic for:", byteCount)

	// A Go string is really just a read-only sequence of bytes.
	// len() counts those bytes, and indexing with s[i] gives you one
	// raw byte, not one character.
	fmt.Println("word:", word)
	fmt.Println("len(word):", len(word), "bytes")
	fmt.Println("word[0]:", word[0], "(a byte, printed as a number)")

	// Printing every byte in hex shows the two accented "é"
	// characters each take 2 bytes in UTF-8, while plain ASCII
	// letters take 1.
	fmt.Print("bytes in hex: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i])
	}
	fmt.Println()

	// A rune is Go's name for a single Unicode code point (one
	// character), stored as an int32. utf8.RuneCountInString counts
	// characters correctly, unlike len().
	fmt.Println("\nutf8.RuneCountInString(word):", utf8.RuneCountInString(word), "characters")

	// range over a string decodes the UTF-8 bytes for you and gives
	// each rune with the byte index it starts at.
	fmt.Println("\nranging over runes:")
	for index, r := range word {
		fmt.Printf("%q starts at byte %d\n", r, index)
	}

	// Converting a string to []rune gives you a slice you CAN index
	// by character, because now every element is a full rune.
	letters := []rune(word)
	fmt.Println("\n[]rune(word):", letters, "len:", len(letters))
	fmt.Printf("letters[1]: %q (the accented e, not a broken byte)\n", letters[1])

	// Converting back to a string reassembles it.
	fmt.Println("string(letters):", string(letters))

	// Dart comparison: Dart strings are UTF-16, and word[i] in Dart
	// gives you a code unit that is usually (not always) a full
	// character, which can hide this problem entirely for typical
	// text. Go's raw-UTF-8 model always forces you to think about
	// bytes vs. characters, which matters once you handle text
	// beyond plain ASCII (names, emoji, other languages).

	// Backend connection: this matters whenever your API accepts
	// free-text input (names, comments, search queries) — validating
	// or truncating by len() truncates by BYTES and can cut a
	// multi-byte character in half; use RuneCountInString or []rune
	// when you actually mean "characters".
}
