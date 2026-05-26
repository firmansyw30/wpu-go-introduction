package main

import (
	"fmt"
)

// Array comparison
func compareArray() bool {
	fmt.Println("Array Comparison")
	var isDifferent bool
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [4]int{1, 2, 3, 4}
	// Using range to modernize loop instead of traditional for loop
	for i := range arr1 {
		if arr1[i] == arr2[i] {
			fmt.Println("Element at index", i, "is equal in both arrays:", arr1[i], arr2[i])
			isDifferent = false
		} else {
			fmt.Println("Element at index", i, "is not equal in both arrays:", arr1[i], arr2[i])
			isDifferent = true
		}
	}
	return isDifferent
}

func ArrayFunc() {
	// Example array manipulation
	fmt.Println("Example Array Manipulation")
	var exampleArray [4]int // Quick way to declare and initialize an array until 4 elements
	for i := 0; i < len(exampleArray); i++ {
		exampleArray[i] = i * 2
	}
	fmt.Println("Example array after manipulation:", exampleArray)

	for index, value := range exampleArray {
		fmt.Println("Element at index", index, "is", value)
	}
}

func anotherArrayComparison() {
	// Another example of array manipulation
	fmt.Println("Another Array Comparison")
	array1 := [4]int{1, 2, 3, 4}
	array2 := [4]int{4, 3, 2, 1}
	// Comparing arrays
	fmt.Println("Array1:", array1)
	fmt.Println("Array2:", array2)
	fmt.Println("Is array1 equal to array2?", array1 == array2)
	fmt.Println("Is array1 not equal to array2?", array1 != array2)
}
func main() {
	// Declare and initialize an array (without looping)
	fmt.Print("Declare and initialize an array (without looping)\n")
	var someNumber [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(someNumber)
	fmt.Println("Element at index 0 is:", someNumber[0])
	fmt.Println("")
	someNumber[2] = 6
	fmt.Println("Updated element at index 2 is:", someNumber[2])

	// Declare and initialize an array with looping
	fmt.Println("Declare and initialize an array with looping")
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Access and print array elements
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, "is", numbers[i])
	}

	// Counting element in the array
	fmt.Println("")
	fmt.Println("Counting element in the array")

	var newArray1 [4]int = [4]int{1, 2, 3, 4}
	var countArray1 int //used for counting elements
	for i := 0; i < len(newArray1); i++ {
		countArray1++
	}
	fmt.Println("From this Array:", newArray1)
	fmt.Println("Total element in the array is:", countArray1)
	// Compare two arrays by calling function
	fmt.Print("\n")
	result := compareArray()
	fmt.Println("Arrays are different:", result)

	fmt.Print("\n")
	anotherArrayComparison()

	fmt.Print("\n")
	ArrayFunc()
}
