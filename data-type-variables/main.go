package main

import (
	"fmt"
)

func main() {
	// declaration (classic style)
	var firstname string = "Firmansyah"
	var lastname string = "Wicaksono"
	var age int = 24
	var interest = []string{"Coding", "Music", "Sports"}

	// inference (automatically guess the data type)
	city := "Bandung"
	year := 2002
	hobbies := []string{"Reading", "Gaming", "Traveling"}
	firstAndLastName := firstname + " " + lastname

	// constanta (cannot)
	const zodiac = "Aquarius"

	// signed integers
	si8 := int8(-30)
	si16 := int16(-30000)
	si32 := int32(-46273)
	si64 := int64(-86854860)
	sint := int(-100) // platform-dependent (32 or 64 bit)

	// unsigned integers
	ui8 := uint8(30)
	ui16 := uint16(30000)
	ui32 := uint32(307367)
	ui64 := uint64(26358260)
	uint := uint(100) // platform-dependent (32 or 64 bit)

	// floating point data type
	fp32 := float32(3.14)
	fp64 := float64(3.141592653589793)
	fp := 2.71828

	var floatingpoint32 float32 = 3.15
	var floatingpoint64 float64 = 3.151592653589793
	var floatingpoint float64 = 2.35262

	// boolean data type
	isMarried := false
	hasChildren := false
	var isAlive bool = true
	var isStudent bool = false

	// string data type using backtick (``) for raw text
	var aboutme string = `
Hello, my name is Firman a DevOps Engineer from Bandung.
I am interested in technologies such as Docker, Kubernetes, and Cloud Computing.
I have experience in developing and deploying applications on various cloud platforms.
I am also skilled in containerization and orchestration.
I have a strong passion for automation and have developed scripts to automate various tasks.
I am a team player and enjoy working with others to achieve a common goal.
I am a quick learner and always eager to learn new technologies and skills.
I am a problem solver and enjoy debugging code to find the root cause of an issue.
I am a good communicator and enjoy explaining complex concepts to non-technical people.
I am a leader and enjoy guiding others to achieve their goals.
`
	// string data type using quotation ("") for single line text
	personality := "I'm Friendly, Curious, and Ambitious personality"

	// print variable
	fmt.Println("First Name:", firstname)
	fmt.Println("Last Name:", lastname)
	// String Operation to concate each different string
	fmt.Println("Full Name:", firstAndLastName)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)
	fmt.Println("Hobbies:", hobbies)
	fmt.Println("Interests:", interest)
	fmt.Println("Zodiac:", zodiac)
	fmt.Print("\n")

	//print integer type using printf in order to print the value
	fmt.Println("Signed Integers")
	fmt.Printf("SI (8): %v\n", si8)
	fmt.Printf("SI (16): %v\n", si16)
	fmt.Printf("SI (32): %v\n", si32)
	fmt.Printf("SI (64): %v\n", si64)
	fmt.Printf("SI (Platform): %v\n", sint)

	fmt.Print("\n")
	fmt.Println("Unsigned Integers")
	fmt.Printf("UI (8): %v\n", ui8)
	fmt.Printf("UI (16): %v\n", ui16)
	fmt.Printf("UI (32): %v\n", ui32)
	fmt.Printf("UI (64): %v\n", ui64)
	fmt.Printf("UI (Platform): %v\n", uint)

	// floating point data type
	fmt.Print("\n")
	fmt.Println("Floating Point")
	fmt.Printf("FP (32): %v\n", fp32)
	fmt.Printf("FP (64): %v\n", fp64)
	fmt.Printf("FP (Platform): %v\n", fp)
	fmt.Printf("Floating Point (32): %v\n", floatingpoint32)
	fmt.Printf("Floating Point (64): %v\n", floatingpoint64)
	fmt.Printf("Floating Point (Platform): %v\n", floatingpoint)

	fmt.Print("\n")
	fmt.Println("Boolean")
	fmt.Printf("Is Married: %v\n", isMarried)
	fmt.Printf("Has Children: %v\n", hasChildren)
	fmt.Printf("Is Alive: %v\n", isAlive)
	fmt.Printf("Is Student: %v\n", isStudent)

	fmt.Print("\n")
	fmt.Println("String")
	fmt.Printf("About Me: %v\n", aboutme)
	fmt.Printf("Personality: %v\n", personality)

	// string operation to find length
	fmt.Printf("About Me Length: %v\n", len(aboutme))
	fmt.Printf("Personality Length: %v\n", len(personality))
	fmt.Printf("First Character of First Name: %v\n", string(firstname[0]))
}
