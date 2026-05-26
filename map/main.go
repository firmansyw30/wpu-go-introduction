package main

import "fmt"

func mapLiteral() {
	fmt.Println("Map Literal")
	// Define Map (Literal)
	mapData := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Define map operations
	value, exists := mapData["two"] // Accessing value by key and checking if the key exists
	mapData["four"] = 4             // Adding new key-value pair to the map
	delete(mapData, "one")          // Deleting a key-value pair from the map

	// Print Map
	fmt.Printf("Map Data: %v\n", mapData)
	fmt.Printf("Value for key 'two': %d, Exists: %t\n", value, exists)
}

func mapWithMake() {
	fmt.Println("Map with Make")
	// Define Map (with make)
	mapData := make(map[string]int)
	mapData["one"] = 1
	mapData["two"] = 2
	mapData["three"] = 3

	// Print Map
	fmt.Printf("Map Data: %v\n", mapData)
}

func mapExample() {
	fmt.Println("Map Example")
	// Define Map (Literal)
	mapData := map[string]string{
		"Name":   "John",
		"Class":  "10th",
		"School": "ABC High School",
	}

	// Print Map
	fmt.Printf("Map Data: %v\n", mapData)
	fmt.Printf("Name: %s\n", mapData["Name"])
	fmt.Printf("Class: %s\n", mapData["Class"])
	fmt.Printf("School: %s\n", mapData["School"])
}

// Call all map function
func main() {
	mapLiteral()
	fmt.Println("")
	mapWithMake()
	fmt.Println("")
	mapExample()

}
