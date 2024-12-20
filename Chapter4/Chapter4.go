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
	fmt.Println(randSlice)
}
