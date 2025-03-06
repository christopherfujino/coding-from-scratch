package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(101)
	fmt.Println("Guess the number from 0 to 100!")
	for {
		var guess int
		fmt.Scanf("%d", &guess)
		if guess == n {
			fmt.Printf("You correctly guessed %d!\n", guess)
			break
		} else if guess > n {
			fmt.Printf("Your guess of %d is too large.\n", guess)
		} else {
			fmt.Printf("Your guess of %d is too small.\n", guess)
		}
	}
}
