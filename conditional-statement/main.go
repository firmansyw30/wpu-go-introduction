package main

import "fmt"

func driverLicenseCheck(age int) {
	fmt.Println("Driver's License Check")
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are eligible for a driver's license.")
	} else {
		fmt.Println("You are not eligible for a driver's license. Because your age is,", age, "years old.")
	}
}

func checkEvenOddNumber() {
	var num int
	fmt.Println("Check Even or Odd Number (Ganjil atau Genap)")
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d is an even (genap) number.\n", num)
	} else {
		fmt.Printf("%d is an odd (ganjil) number.\n", num)
	}
}

func basicConditionalIf() {
	a := 5
	fmt.Println("Basic Conditional with if")
	if a > 3 {
		fmt.Println("a =", a)
		fmt.Println("a is greater than 3")
	}
}
func basicConditionalIfElse() {
	var a int = 10
	fmt.Println("Basic Conditional with else")
	if a > 5 {
		fmt.Println("a =", a)
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a =", a)
		fmt.Println("a is less than or equal to 5")
	}
}

func basicConditionalWithElseIf() {
	var a int
	fmt.Println("Basic Conditional with else if")
	fmt.Print("Enter a number: ")
	fmt.Scanln(&a)

	if a > 15 {
		fmt.Println("a =", a)
		fmt.Println("a is greater than 15")
	} else if a > 5 {
		fmt.Println("a =", a)
		fmt.Println("a is greater than 5 but less than or equal to 15")
	} else {
		fmt.Println("a =", a)
		fmt.Println("a is less than or equal to 5")
	}
}

func ifShortStatement() {
	var num int
	fmt.Println("If Short Statement")
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if num := num * 2; num > 5 {
		fmt.Println("num =", num)
		fmt.Println("num is greater than 5")
	}
}
func main() {
	var age int
	driverLicenseCheck(age)
	fmt.Println("")
	checkEvenOddNumber()
	fmt.Println("")
	basicConditionalIf()
	fmt.Println("")
	basicConditionalIfElse()
	fmt.Println("")
	basicConditionalWithElseIf()
	fmt.Println("")
	ifShortStatement()
	fmt.Println("")
}
