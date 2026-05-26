package main

import "fmt"

func initialForLoop() {
	fmt.Println("Initial For Loop")
	someArray := [5]string{"Go", "is", "a", "great", "language"}
	for i := 0; i < len(someArray); i++ {
		fmt.Println("Printed array is:", someArray[i])
	}
}

func initialForLoop2() {
	fmt.Println("Initial For Loop 2")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
}

func initialForLoop3() { // Like while loop in other languages
	fmt.Println("Initial For Loop 3")
	i := 1
	for i <= 5 {
		fmt.Printf("Iteration %d\n", i)
		i++
	}
}

func forLoopWithBreak() {
	fmt.Println("For Loop with Break")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at iteration", i)
			break
		}
		fmt.Printf("Iteration %d\n", i)
	}
}

func forLoopWithContinue() {
	fmt.Println("For Loop with Continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Skipping even iteration %d\n", i)
			continue
		}
		fmt.Printf("Iteration %d\n", i)
	}
}

func forLoopWithRange() {
	fmt.Println("For Loop with Range")
	someArray := []string{"Go", "is", "a", "great", "language"}
	for index, value := range someArray {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}

func forLoopOddNumber() {
	var even int
	fmt.Println("For Loop Even Number with while style")
	fmt.Print("Enter the Desired end (Even Number): ")
	fmt.Scanln(&even)

	if even%2 == 0 {
		fmt.Printf("The Desired number %d is not an even number. Please enter an even number.\n", even)
		fmt.Scanln(&even)
	} else if even < 0 {
		fmt.Printf("The Desired number %d is not a positive number. Please enter a positive even number. \n", even)
		fmt.Scanln(&even)
	} else {
		fmt.Printf("The Even Number Loop from 2 to %d is", even)
		fmt.Println("")
		for i := 2; i <= even; i += 2 {
			fmt.Printf("%d, ", i)
		}
	}
}

func forLoopCountDown() {
	fmt.Println("For Loop Countdown to 0")
	var num int
	fmt.Print("Enter a number to count down from: ")
	fmt.Scanln(&num)

	fmt.Println("Countdown from", num, ":")
	for {
		fmt.Printf("%d\n", num)
		num--
		if num == 0 {
			break
		}
	}
}
func main() {
	initialForLoop()
	initialForLoop2()
	initialForLoop3()
	forLoopWithBreak()
	forLoopWithContinue()
	forLoopWithRange()
	forLoopOddNumber()
	forLoopCountDown()
}
