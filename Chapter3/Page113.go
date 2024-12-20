package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	// Exercise 1
	// Setting up slice
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	// Subslice containing first two values
	var x []string = greetings[:2]
	// Subslice containing second, third & fourth values
	var y []string = greetings[1:4]
	// Subslice containing fourth & fifth values
	var z []string = greetings[3:]

	// Printing all slices
	fmt.Println(greetings)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)


	// Exercise 2
	// Runes
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println("Rune:", string(runes[4]))

	// Exercise 3
	// Struct
	per1 := Employee{"John", "Doe", 12}
	per2 := Employee{
		firstName: "Jane",
		lastName:  "Doe",
		id:        34,
	}
	var per3 Employee
	per3.firstName = "JJ"
	per3.lastName = "Doe"
	per3.id = 56

	fmt.Println("Person 1:", per1)
	fmt.Println("Person 2:", per2)
	fmt.Println("Person 3:", per3)

}
