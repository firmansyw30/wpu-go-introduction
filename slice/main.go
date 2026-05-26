// Slice concept is like a part of pizza that flexible.
// For example, if you have a whole pizza and you want to give a part of it to your friend, you can slice it and give only the part that you want to share.
// In Go, slice is a data structure that allows you to access a subset of elements in an array or string.
// You can define the start and end index of the slice, and even omit the end index if you want to slice up to the end of the array or string.
package main

import "fmt"

func sliceBasic() {
	fmt.Println("Slice Basic")
	// Define Array for slice (require 3 args : data type, length, capacity)
	arr := make([]int, 10, 20)
	fmt.Println(arr)
	fmt.Println("Length:", len(arr))
	fmt.Println("Capacity:", cap(arr))

	fmt.Println("")
	fmt.Println("Append Operations in slice")
	arr = append(arr, 1, 2, 3, 4)
	fmt.Println("After append: ", arr)

	fmt.Println("")
	fmt.Println("Copy Operations in slice")
	sliceArr2 := make([]int, 3)
	sliceArr3 := copy(sliceArr2, arr)
	fmt.Println("Copied: ", sliceArr2)
	fmt.Println("Number of elements copied: ", sliceArr3)
}

func customSlicing() {
	fmt.Println("Custom Slicing")
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := number[2:5] // Slicing from index 2 to 5 (index 5 is not included)
	fmt.Printf("The Original Array (Number): %v\n", number)
	fmt.Println("Sliced from index 2 to 5: ", slice)
}
func main() {
	fmt.Println("Slice Introduction")
	// Define Array
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Define slice operation
	sliceAllArray := arr[:]              // Slice all elements
	sliceFromIndexToCustom := arr[:8]    // Slice from beginning to index 8, but the index 8 isn't included
	sliceFromIndexCustomToEnd := arr[2:] // Slice from index 2 to 10
	sliceFromIndexToIndex := arr[2:6]    // Slice from index 2 to 6

	// Define slice function
	returnSliceElement := len(sliceFromIndexToIndex) // Using len slice function to return the length of the slice (count of existing elements)
	returnSliceCapacity := cap(make([]int, len(sliceFromIndexToIndex)))
	appendNewElementToSlice := append(sliceFromIndexToIndex, 11)                      // Using append slice function to add new element to the slice
	copySlice := copy(make([]int, len(sliceFromIndexToIndex)), sliceFromIndexToIndex) // Using copy slice function to copy the slice to a new slice
	makeNewSlice := make([]int, 4, 6)
	makeNewSliceFromIndex := []int{1, 2, 3}
	makeNewSliceFromIndex = append(makeNewSliceFromIndex, 4) // Using make slice function to create a new slice with length and capacity, and then using append to add new element to the slice

	// Print Array
	fmt.Printf("Original Array: %v\n", arr)

	// Print slice operation
	fmt.Println("Slice Operation")
	fmt.Printf("Slice All Array: %v\n", sliceAllArray)
	fmt.Printf("Slice from index 0 to index %d: %v\n", len(sliceFromIndexToCustom), sliceFromIndexToCustom)
	fmt.Printf("Slice from index %d to the end: %v\n", sliceFromIndexCustomToEnd[0], sliceFromIndexCustomToEnd)
	fmt.Printf("Slice from index %d to %d: %v\n", sliceFromIndexToIndex[0]-1, sliceFromIndexToIndex[len(sliceFromIndexToIndex)-1], sliceFromIndexToIndex)

	fmt.Println("")

	// Print slice function
	fmt.Println("Slice Function")
	fmt.Printf("Length of slice (element) from index %d to %d: %d\n", sliceFromIndexToIndex[0]-1, sliceFromIndexToIndex[len(sliceFromIndexToIndex)-1], returnSliceElement)
	fmt.Printf("Capacity of slice from index %d to %d: %d\n", sliceFromIndexToIndex[0]-1, sliceFromIndexToIndex[len(sliceFromIndexToIndex)-1], returnSliceCapacity)
	fmt.Printf("Slice after appending new element: %v\n", appendNewElementToSlice)
	fmt.Printf("Copied slice: %v\n", copySlice)
	fmt.Printf("New slice: %v\n", makeNewSlice)
	fmt.Printf("New slice from index: %v\n", makeNewSliceFromIndex)

	fmt.Println("")
	sliceBasic()

	fmt.Println("")
	customSlicing()
}
