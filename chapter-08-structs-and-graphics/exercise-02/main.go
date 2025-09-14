package main

import (
	"fmt"
)

type User struct {
	name string
	age int
}

func (u *User) getUserData() {
	fmt.Printf("What is your name? ")
	fmt.Scanf("%s\n", &u.name)

	fmt.Printf("What is your age? ")
	fmt.Scanf("%d\n", &u.age)
}

func (u User) showUserData() {
	fmt.Printf("User %s is %d years old.\n", u.name, u.age)
}

func main() {
	var user User

	user.getUserData()

	user.showUserData()
}
