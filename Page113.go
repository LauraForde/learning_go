package main

import (
	"fmt"
)

func main() {

	// Exercise 1
	// Setting up slice
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// Subslice containing first two values
	var x []string = greetings[0:2]

	// Subslice containing second, third & fourth values
	var y []string = greetings[1:4]

	// Subslice containing fourth & fifth values
	var z []string = greetings[3:5]

	// Printing all slices
	fmt.Println(greetings)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
