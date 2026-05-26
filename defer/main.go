package main

import "fmt"

func initialDefer() {
	fmt.Println("Initial Defer Example")
	defer fmt.Println("A: Last to execute") // Act like Stack, LIFO (Last In First Out)
	fmt.Println("B: First to execute")
}

func noDefer() {
	fmt.Println("No Defer Example")
	fmt.Println("A: First to execute")
	fmt.Println("B: Second to execute")
}

func multipleDefer() {
	fmt.Println("Multiple Defer Example")
	x := 1
	defer fmt.Println("Defer-1, x = ", x) // Argument x is evaluated at the time of defer statement, not at the time of execution
	x = 2
	defer fmt.Println("Defer-2, x = ", x) // Last defer statement will execute first, so it will print x = 2
	fmt.Println("x = ", x)                // This will print x = 2
}

func main() {
	initialDefer()
	fmt.Println()
	noDefer()
	fmt.Println()
	multipleDefer()
}
