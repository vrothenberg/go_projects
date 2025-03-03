package main

import (
	"fmt"
)

func main() {
	// ASCII vs Unicode characters
	str1 := "Hello"        // ASCII (each letter is 1 byte)
	str2 := "你好世界"        // Unicode (each character takes multiple bytes in UTF-8)

	fmt.Println("ASCII String:", str1)
	fmt.Println("Unicode String:", str2)

	// Byte representation: A string in Go is stored as a sequence of bytes.
	bytes1 := []byte(str1) // Convert ASCII string to a byte slice
	bytes2 := []byte(str2) // Convert Unicode string to a byte slice

	fmt.Println("Bytes of ASCII String:", bytes1) // Each letter is 1 byte
	fmt.Println("Bytes of Unicode String:", bytes2) // Each character is multiple bytes in UTF-8

	// Rune representation: A rune represents a Unicode code point.
	runes1 := []rune(str1) // Convert ASCII string to runes (Unicode points)
	runes2 := []rune(str2) // Convert Unicode string to runes

	fmt.Println("Runes of ASCII String:", runes1) // Each ASCII character is a single-byte rune
	fmt.Println("Runes of Unicode String:", runes2) // Each Chinese character is represented as a single rune

	// Checking individual bytes and runes
	fmt.Printf("First byte of '%s': %v (ASCII: %q)\n", str1, bytes1[0], bytes1[0])
	fmt.Printf("First rune of '%s': %v (Unicode: U+%04X)\n", str2, runes2[0], runes2[0])

	// Length difference
	fmt.Println("Length of str1 (bytes):", len(str1))       // 5 bytes (ASCII: each character is 1 byte)
	fmt.Println("Length of str2 (bytes):", len(str2))       // 12 bytes (UTF-8: each Chinese character is 3 bytes)
	fmt.Println("Rune count of str2 (characters):", len(runes2)) // 4 runes (each Chinese character is 1 rune)

	// UTF-8 encoding breakdown
	fmt.Println("\nUTF-8 Encoding Breakdown for '你好世界':")
	for _, r := range runes2 {
		utf8Bytes := []byte(string(r)) // Convert rune back to UTF-8 bytes
		fmt.Printf("Rune: %c (U+%04X) -> UTF-8 Bytes: %v\n", r, r, utf8Bytes)
	}
}
