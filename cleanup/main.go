package main

import "fmt"

func cleanupExample() {
	fmt.Println("Cleanup: Closing resource..")
}

func readConfig(fileName string) {
	// Defer always called even there's panic
	defer cleanupExample()

	// Recover stored inside defer to handle panic
	defer func() { // Anonymous function
		if r := recover(); r != nil {
			fmt.Println("Error occured: ", r)
		}
	}()

	if fileName == "" {
		panic("Filename shouldn't be empty")
	}

	fmt.Println("Reading config from file: ", fileName)
}

func main() {
	readConfig("")
	fmt.Println("Program still running after panic because we handled it with recover and cleanup is also executed")
}
