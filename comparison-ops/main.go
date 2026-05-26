package main

import "fmt"

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
	fmt.Printf("!(%d < %d) ? %v\n", num1, num2, !lessFirst)
	fmt.Printf("!(%d > %d) ? %v\n", num3, num4, !greaterSecond)
	fmt.Printf("!(%d < %d) ? %v\n", num3, num4, !lessSecond)
}

func main () {
	//customComparison()
	logicalComparison()
	// fmt.Println("")
	// a := 5
	// b := 3
	// fmt.Println("Does a equal to b? The answer is",a==b)
}