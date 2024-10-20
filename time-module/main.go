package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time program!")
	presentTime := time.Now()
	fmt.Println("Present time: ", presentTime)

	// Formatting time using a specific layout
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 01-02-2006 is the layout for MM-DD-YYYY and this is the reference date for Go time package

	// Creating a custom time
	customTime := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Custom time: ", customTime)
}
