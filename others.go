package main

// Import the required Libraries
import (
	"fmt"
	"strings"
)

// Create the validator that will be used to make the calculations
func validator(input1 int, typee string, input2 int) int {
	if typee == "+" {
		return input1 + input2
	} else if typee == "-" {
		return input1 - input2
	} else if typee == "/" {
		return input1 / input2
	} else if typee == "*" {
		return input1 * input2
	}

	// Go obrigatory return
	return 0
}

// Loop the project in a CMD window
func reset() int {
	a := 0
	b := 0
	method := ""

	fmt.Println("Waiting for an input")
	fmt.Scanln(&a)

	fmt.Println("Method")
	fmt.Scanln(&method)

	fmt.Println("Waiting for an input")
	fmt.Scanln(&b)

	fmt.Println("Result: ", validator(a, method, b))

	newinput := ""
	fmt.Println("Press Enter to exit write H to reset")
	fmt.Scanln(&newinput)
	if strings.ToLower(newinput) == "h" {
		reset()
	}

	// Go obrigatory return
	return 0
}
