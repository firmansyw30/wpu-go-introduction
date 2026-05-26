package main

import "fmt"

func main () {
	//Integer Type
	fmt.Println("Math Operations Using Integer Data Type")
	a := 10 
	b := 4

	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b) //
	fmt.Println("a / b = ", a/b) //Divide (Integer)
	fmt.Println("a % b = ", a%b) //Modulus

	//Float Type
	x := 7.0
	y := 2.0
	fmt.Println("x / y = ", x/y)

	// Assignment Operator using integer data type
	fmt.Println("Assignment Operator Using Integer Data Type")
	z := 6 // Placeholder variable
	z = 10 // Actual Variable
	fmt.Println(z)

	z+=2
	fmt.Println("z += 2 : ", z) // equal to z = z + 2

	z-=3
	fmt.Println("z -= 3 : ", z) // equal to z = z - 3

	z*=4
	fmt.Println("z *= 4 : ", z) // equal to z = z * 4

	z/=5
	fmt.Println("z /= 5 : ", z) // equal to z = z / 5

	// Increment Operator using integer data type
	fmt.Println ("Increment Operator Using Integer Data Type")
	a1 := 2
	a2 := 3
	a1++
	a2--
	fmt.Println("a1++: ", a1) //equal to a1 = a1 + 1
	fmt.Println("a2--: ",a2) //equal to a2 = a2 - 1
}