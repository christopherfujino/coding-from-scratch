# Data Structures

## Exercise 1

```go
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
```

## Exercise 2

```go
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
```

### Questions

1. What would happen if you switched the order of the calls to `user.getUserData()` and `user.showUserData()`?

## Exercise 3

```go
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
```

### Questions

1. What kind of problems might you face as the number of users in your systems grows large?

## Exercise 4

```go
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
```

## Exercise 5

```go
package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const height = 480
const width = 640

func main() {
	rl.InitWindow(width, height, "Title")
	rl.SetTargetFPS(60)

	// Did the user hit escape or close the window?
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		const circleCenterX = 20
		const circleCenterY = 20
		const circleRadius = 10
		var circleColor = rl.Pink
		rl.DrawCircle(circleCenterX, circleCenterY, circleRadius, circleColor)

		const rectangleX = 500
		const rectangleY = 300
		const rectangleWidth = 50
		const rectangleHeight = 75
		var rectangleColor = rl.Blue
		rl.DrawRectangle(rectangleX, rectangleY, rectangleWidth, rectangleHeight, rectangleColor)

		rl.EndDrawing()
	}
}
```

### Questions

1. Can you make the shapes move?
2. Can you change the colors of the shapes?
