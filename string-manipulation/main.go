package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String Manipulation")
	fmt.Print("\n")
	// String Manipulation
	var text string = "Hello, World!"
	fmt.Println("Original Text:", text)

	// Convert to Uppercase
	upper := strings.ToUpper(text)
	fmt.Println("Uppercase:", upper)

	// Convert to Lowercase
	lower := strings.ToLower(text)
	fmt.Println("Lowercase:", lower)

	// Trim Whitespace
	trimmed := strings.TrimSpace("   Hello, World!   ")
	fmt.Println("Trimmed:", trimmed)

	// Split String & remove the delimiter
	words := strings.Split(text, " ")
	fmt.Println("Split:", words)

	// Join String
	joined := strings.Join(words, "-")
	fmt.Println("Joined:", joined)

	// Check does string started with "Hello" (using variable)
	startsWithHello := strings.HasPrefix(text, "Hello")
	fmt.Println("Starts with 'Hello'? :", startsWithHello)

	// Check does string contained "World" with single line
	fmt.Println("Contains 'World'? :", strings.Contains(text, "World"))

	// Replace string "World" with "Go", the last number 1 define the index of the text
	fmt.Println("Replace 'World' with 'Go':", strings.Replace(text, "World", "Go", 1))

}
