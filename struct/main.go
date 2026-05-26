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

func main() {
	user1 := User{
		Name:  "Firman",
		Email: "firmansyahwicaksono30@gmail.com",
		Age:   25,
		Address: Address{
			City:      "Jakarta",
			State:     "DKI Jakarta",
			Country:   "Indonesia",
			PostalCode: 12345,
		},
	}

	fmt.Println("Name : ", user1.Name)
	fmt.Println("Email : ", user1.Email)
	fmt.Println("Age : ", user1.Age)
	fmt.Println("City : ", user1.City)
	fmt.Println("State : ", user1.State)
	fmt.Println("Country : ", user1.Country)
	fmt.Println("Postal Code : ", user1.PostalCode)
}
