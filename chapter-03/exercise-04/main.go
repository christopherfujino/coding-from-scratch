package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i int
	fmt.Printf("What is your first number? ")
	i = rand.Int()
	fmt.Printf("Your number was: %d\n", i)
	i = rand.Int()
	fmt.Printf("Your number was: %d\n", i)
}
