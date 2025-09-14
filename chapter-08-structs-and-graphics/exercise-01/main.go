package main

import (
	"fmt"
)

type User struct {
	name string
	age int
}

func main() {
	var user User
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s\n", &user.name)

	fmt.Printf("What is your age? ")
	fmt.Scanf("%d\n", &user.age)

	fmt.Printf("User %s is %d years old.\n", user.name, user.age)
}
