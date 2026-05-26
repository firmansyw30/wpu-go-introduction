package main

import "fmt"

func greating() {
	fmt.Println("Hello, World!")
}

func greating2(name string) {
	fmt.Println("Hello, " + name + "! Welcome to Go Language")
}

func greating3(times, name string) {
	fmt.Println("Hello, " + name + "! Good " + times + "! Welcome to Go Language!")
}

func sampleAddition(a, b int) int {
	return a + b
}

func multipleReturnFunction() (string, int) {
	return "Firman", 24
}

func sampleSubstraction(a, b int) (result int) {
	result = a - b
	return
}

func variadicFunction(numbers ...int) int {
	sum := 0
	for _, i := range numbers { // _ used to ignore the index
		sum += i // equal to sum = sum + i
	}
	return sum
}

func main() {
	greating()
	greating2("Firman")
	greating3("Morning", "Firman")

	result := sampleAddition(10, 20)
	fmt.Println(result)

	name, age := multipleReturnFunction()
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	fmt.Println(sampleSubstraction(10, 6))

	fmt.Println(variadicFunction(1, 2, 3))       // Output 6
	fmt.Println(variadicFunction(10, 20, 30, 4)) // Output 64

}
