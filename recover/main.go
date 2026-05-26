package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil { // recover() will return nil if there is no panic, otherwise it will return the value passed to panic()
		fmt.Println("Panic happened, but we recovered from it:", r)
	}
}

func divide(a, b int) {
	defer handlePanic() // defer will ensure that handlePanic called even if there is a panic in this function
	fmt.Printf("Dividing %d by %d\n", a, b)
	result := a / b // this will trigger panic if b = 0
	fmt.Println("Result: ", result)
}

func main() {
	fmt.Println("Recover example")
	divide(10, 2) // output normal
	divide(5, 0)  // output will trigger panic, but recover will handle it and prevent the program from crashing
	fmt.Println("Program finish securely without crashing due to panic")
}
