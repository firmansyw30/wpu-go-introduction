package main

import "fmt"

// First Example of Pointer in Go
func changeName(name *string) {
	*name = "Firman" // change directly the value of name variable in main function using pointer
}

func callChangeName() {
	fmt.Println("Sample Pointer in Go")
	name := "John"
	fmt.Println("Name before change: ", name)
	changeName(&name) // pass the address of name variable to changeName function
	fmt.Println("Name after change: ", name)
}

// Second Example of Pointer in Go
func heal(hp *int) {
	*hp += 20 // Add HP directly to memory address of health variable in main function using pointer
	fmt.Println("Player healed +20 HP!")
}

func attack(hp *int, damage int) {
	*hp -= damage // Subtract HP directly to memory address of health variable in main function using pointer
	fmt.Printf("Player attacked with %d damage!\n", damage)
	if *hp <= 0 {
		fmt.Println("Game Over! Player has been defeated.")
	}
}

func playGame() {
	fmt.Println("Sample Pointer in Go for simple game simulation")
	hp := 50 // Initial HP of player
	fmt.Printf("Player's initial HP: %d\n", hp)
	heal(&hp) // Heal player by passing the address of hp variable to heal function
	fmt.Printf("Player's HP after healing: %d\n", hp)
	attack(&hp, 30) // Attack player by passing the address of hp variable and damage value to attack function
	fmt.Printf("Player's HP after attack: %d\n", hp)
	attack(&hp, 25) // Attack player again to trigger game over
	fmt.Printf("Player's HP after second attack: %d\n", hp)
	attack(&hp, 15) // Attack player again to trigger game over
	fmt.Printf("Player's HP after third attack: %d\n", hp)
}

func main() {
	callChangeName()
	fmt.Println()
	playGame()
}
