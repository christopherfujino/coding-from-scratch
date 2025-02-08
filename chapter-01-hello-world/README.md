# Lesson 1 - Hello, World!

## Get and configure your editor

Download Visual Studio Code (commonly called "VS Code"). Note that this is a
different program from the similarly named "Visual Studio".

VS Code is a text editor. This means its primary job is to let you write and
edit the code you write. However, VS Code *can* do other things too,
especially if you install extensions. Extensions give VS Code the ability to
do things like run your code, detect and point out mistakes that you make,
and help you understand the code you write.

Open VS Code. In the menu, navigate to "View" -> "Extensions". There should
now be a pane on the left-hand side of your window titled "EXTENSIONS". In the
search bar, type in "Go". We will be installing the extension for the Go
programming language.

Click on the extension entitled "Go". Ensure that it says "Go Team at Google"
with a checkmark, meaning that the VS Code team has verified the identity of
the publishers of this extension. It is important that you only install VS
Code extensions from authors who you trust.

Once you've verified the author of the extension, you can press the "Install"
button.

Create a new directory (or folder) to hold all of your Go source code, called
something like `go_source`.

For each lesson, create a new directory under it, such as `lesson_01`.

Each Go program needs to have its own directory. For each exercise in this
lesson, create a new directory, such as `exercise_01`.

In the directory `go_source/lesson_01/exercise_01`, Create a new text file and
type the following program:

```go
// lesson_01/exercise_01

package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Save it as `main.go`. The name "main" could be anything, but its important
that the extension of all the source code you write is ".go".

From the left side-bar, select the debug section, and press "Run and Debug".

`func main()` defines a new function named `main`. You can do two things with
a function:

1. Define it
2. Call it

The `main` function is special. Every program must have exactly one `main`
program. When the program starts, the `main` function is called, and when the
`main` function is finished, the program finishes.

## Variables

```go
// lesson_01/exercise_02

package main

import "fmt"

func main() {
	var message string = "Hello, world!"

	fmt.Println(message)
}
```

You can do three things with a variable:

1. Declare it - give it a name and a type
2. Assign it - give it a value
3. Reference it - retrieve the value from it

## Printf

```go
// lesson_01/exercise_03

package main

import "fmt"

func main() {
    var name string = "Alice"
    fmt.Printf("Hello %s\n", name)
}
```

```go
// lesson_01/exercise_04

package main

import "fmt"

func main() {
    var answer int = 29 + 13
    fmt.Printf("The answer is %d\n", answer)
}
```
