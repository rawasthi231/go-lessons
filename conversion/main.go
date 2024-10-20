package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(("Welcome to the conversion program!"))
	fmt.Println("Enter a number: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)

	// Conversion
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Converted number: ", number)
	}

}
