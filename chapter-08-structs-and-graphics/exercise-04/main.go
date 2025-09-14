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
	var users = map[string]User{}

	for i := 0; i < userCount; i += 1 {
		var user = User{}
		user.getUserData()
		users[user.name] = user
	}

	var targetUsername string
	fmt.Printf("Which user would you like to know about? ")
	fmt.Scanf("%s\n", &targetUsername)

	var user, ok = users[targetUsername]
	if !ok {
		fmt.Printf("Oops, I don't know about the user %s!\n", targetUsername)
	} else {
		user.showUserData()
	}
}
