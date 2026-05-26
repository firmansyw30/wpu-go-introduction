package main

import "fmt"

func expressionSwitch() {
	var choice int

	fmt.Println("Switch Case Example")
	fmt.Println("1. Custom Comparison")
	fmt.Println("2. Logical Comparison")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		customComparison()
	case 2:
		logicalComparison()
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}

func customComparison() {
	var num1, num2 int

	fmt.Println("Custom Comparison")
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("\nComparison Result")
	fmt.Printf("%d == %d ? %v\n", num1, num2, num1 == num2)
	fmt.Printf("%d != %d ? %v\n", num1, num2, num1 != num2)
	fmt.Printf("%d < %d ? %v\n", num1, num2, num1 < num2)
	fmt.Printf("%d > %d ? %v\n", num1, num2, num1 > num2)
	fmt.Printf("%d <= %d ? %v\n", num1, num2, num1 <= num2)
	fmt.Printf("%d >= %d ? %v\n", num1, num2, num1 >= num2)
}

func logicalComparison() {
	var num1, num2, num3, num4 int

	fmt.Println("Logical Comparison")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter third number: ")
	fmt.Scanln(&num3)

	fmt.Print("Enter fourth number: ")
	fmt.Scanln(&num4)
	greaterFirst := num1 > num2
	lessFirst := num1 < num2
	greaterSecond := num3 > num4
	lessSecond := num3 < num4

	fmt.Println("\nComparison Result")

	fmt.Println("\nAnd (&&)")
	fmt.Printf("%d > %d && %d > %d ? %v\n", num1, num2, num3, num4, greaterFirst && greaterSecond)
	fmt.Printf("%d > %d && %d < %d ? %v\n", num1, num2, num3, num4, greaterFirst && lessSecond)
	fmt.Printf("%d < %d && %d > %d ? %v\n", num1, num2, num3, num4, lessFirst && greaterSecond)
	fmt.Printf("%d < %d && %d < %d ? %v\n", num1, num2, num3, num4, lessFirst && lessSecond)

	fmt.Println("\nOr (||)")
	fmt.Printf("%d > %d || %d > %d ? %v\n", num1, num2, num3, num4, greaterFirst || greaterSecond)
	fmt.Printf("%d > %d || %d < %d ? %v\n", num1, num2, num3, num4, greaterFirst || lessSecond)
	fmt.Printf("%d < %d || %d > %d ? %v\n", num1, num2, num3, num4, lessFirst || greaterSecond)
	fmt.Printf("%d < %d || %d < %d ? %v\n", num1, num2, num3, num4, lessFirst || lessSecond)

	fmt.Println("\nNot (!)")
	fmt.Printf("!(%d > %d) ? %v\n", num1, num2, !greaterFirst)
}

func shortStatementSwitch() {
	var word string

	fmt.Println("Short Statement Switch Example")
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)

	switch wordLength := len(word); {
	case wordLength > 10:
		fmt.Printf("The Word \"%s\" is longer than 10 characters.\n", word) // the %s are used to format the string, %s is for string and %d is for integer
	case wordLength >= 5:
		fmt.Printf("The Word \"%s\" is between 5 and 10 characters.\n", word)
	default:
		fmt.Printf("The Word \"%s\" is shorter than 5 characters.\n", word)
	}
}

func main() {
	expressionSwitch()
	fmt.Println("")
	shortStatementSwitch()
}
