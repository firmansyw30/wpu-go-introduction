package main

import "fmt"

func panicExample() {
	defer fmt.Println("Defer will be run before the program terminated")
	fmt.Println("Before panic")
	panic("There's something fatal happened")

	// The code below will never be executed because panic will stop the normal execution flow of the program
	fmt.Println("After panic") // This will never be printed because execution stopped in panic
}

func main() {
	panicExample()
}
