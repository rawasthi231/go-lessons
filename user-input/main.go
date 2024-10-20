package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Hello, World!"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok || error ok syntax

	name, error := reader.ReadString('\n')
	fmt.Println("Hello,", name)
	fmt.Printf("Type: %T\n", name)
	fmt.Println("Error: ", error)

}
