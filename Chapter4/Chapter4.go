package main

import (
	"fmt"
 	"math/rand"
)

func main(){

	// Exercise 1

	randSlice := make([]int, 0, 100)
	for i := 0; i < 100; i ++{
		randSlice = append(randSlice, rand.Intn(100))
	}

	// Exercise 2
	for _, v := range randSlice{
		switch{
			case v%6 == 0:
				fmt.Println("Six!")
			case v%2 == 0:
				fmt.Println("Two!")
			case v%3 == 0:
				fmt.Println("Three!")
			default:
				fmt.Println("Never mind")
		}
	}

	// Exercise 3
	var total int
	for i := 0; i < 10; i++{
		total := total + i
		fmt.Println(total)
	}
}
