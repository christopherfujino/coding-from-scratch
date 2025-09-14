package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
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
	const userCount = 3
	var users = [userCount]User{}

	for i := 0; i < userCount; i += 1 {
		users[i].getUserData()
	}

	var targetUsername string
	fmt.Printf("Which user would you like to know about? ")
	fmt.Scanf("%s\n", &targetUsername)

	var foundUser = false
	for _, user := range users {
		if user.name == targetUsername {
			foundUser = true
			user.showUserData()
		}
	}

	if !foundUser {
		fmt.Printf("Oops, I don't know about the user %s!\n", targetUsername)
	}
}
