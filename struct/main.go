package main

import "fmt"

// Struct is a composite data type that groups together variables under a single name. Each variable in a struct is called a field.
// Structs are used to create complex data types that represent real-world entities.

// Sample Composition (Embedding) in Go
type Address struct {
	City, State, Country string
	PostalCode           int
}

type User struct {
	Name    string
	Email   string
	Age     int
	Address // embedding struct Address into User, so that we can access its fields directly from User
}

func callUserStruct() {
	fmt.Println("Sample Struct Composition (Embedding) in Go")
	user1 := User{
		Name:  "Firman",
		Email: "firmansyahwicaksono30@gmail.com",
		Age:   25,
		Address: Address{
			City:       "Jakarta",
			State:      "DKI Jakarta",
			Country:    "Indonesia",
			PostalCode: 12345,
		},
	}

	fmt.Println("Name : ", user1.Name)
	fmt.Println("Email : ", user1.Email)
	fmt.Println("Age : ", user1.Age)
	fmt.Println("City : ", user1.City)              // We can access City directly from User struct because of embedding
	fmt.Println("State : ", user1.State)            // We can access State directly from User struct because of embedding
	fmt.Println("Country : ", user1.Country)        // We can access Country directly from User struct because of embedding
	fmt.Println("Postal Code : ", user1.PostalCode) // We can access PostalCode directly from User struct because of embedding
}

// New Struct Example
type Rectangle struct { // Persegi Panjang
	width  float64 // Panjang
	height float64 // Lebar
}

// Method to calculate area of rectangle
func (r Rectangle) Area() float64 { // Receiver function to calculate area of rectangle. THe actual function is Area, and the receiver is Rectangle (r Rectangle)
	return r.width * r.height
}

func countArea() {
	fmt.Println("Sample Struct with Method in Go for count Area of Rectangle")
	rect := Rectangle{width: 5.0, height: 3.0}
	fmt.Printf("Width of rectangle: %.2f\n", rect.width)
	fmt.Printf("Height of rectangle: %.2f\n", rect.height)
	fmt.Printf("Area of rectangle: %.2f\n", rect.Area())
}

func main() {
	callUserStruct()
	fmt.Println()
	countArea()
}
