package main

import "fmt"

// Constant declaration

// Private constant
const pi = 3.1415

// Public constant(Because of the first letter is uppercase)
const Token = "sjsajkfhdfhnkfnlasfndjfsknf"

func main() {
	var username string = "root"
	fmt.Println("Username: ", username)
	fmt.Printf("Type: %T\n", username)

	var isAlive bool = false
	fmt.Println("Is user alive: ", isAlive)
	fmt.Printf("Type: %T\n", isAlive)

	var smallInt uint8 = 8
	fmt.Println("A smallInt value: ", smallInt)
	fmt.Printf("Type: %T\n", smallInt)

	var smallFloat float32 = 255.446463154215463120
	fmt.Println("A smallFloat value: ", smallFloat)
	fmt.Printf("Type: %T\n", smallFloat)

	// Default values and aliases

	var defaultInt int

	fmt.Println("Default int value: ", defaultInt)
	fmt.Printf("Type: %T\n", defaultInt)

	// Implicit type

	var implicitInt = 10
	fmt.Println("Implicit int value: ", implicitInt)

	// No var style (only works inside functions)

	implicitString := "Hello, World!"
	fmt.Println("Implicit string value in no var style: ", implicitString)

	fmt.Println("Pi value: ", pi)
	fmt.Printf("Type: %T\n", pi)
	fmt.Println("Token value: ", Token)
	fmt.Printf("Type: %T\n", Token)
}
