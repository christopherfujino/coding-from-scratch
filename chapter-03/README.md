# Lesson 3 - Variables

Recall from lesson 1 that there are 3 things that you can do with a variable:

1. Declare it - give it a name and a type
2. Assign it - give it a value
3. Reference it - retrieve the value from it

We've already done all 3, but in this lesson we will explore them in more
depth.

## Declarations

We have already declared many variables with a syntax that looks like:

```go
var x int = 0
```

However, there are other ways to declare a variable. The following lines are
all equivalent, and any one of them would be compiled into the same machine
code instructions:

```go
var x int = 42
var x = 2
x := 3
var x int
```

In the second example, `var x = 2`, we ommitted the type annotation `int`.
The Go compiler is smart enough to figure out that the type of `x` should be
`int` because the value on the right hand side is a whole number. This feature
of the compiler is called "type inference", where the programmer does not
specify a type but the compiler (or other tools) figure it out anyway from the
context. Because it involves less typing (and programmers are lazy), this form
is used more often than the longer one.

There are two situations where you cannot use type inference. The first, is if
you are assigning the variable a value at the same time you are declaring it
(for example, our fourth line above, `var x int`). If you just said `var x`,
the compiler would throw an error, because it cannot infer what type `x`
should be. The second situation, is if you know that you want your variable to
be of a type that is not the default for the value you are assigning it. For
the value `2`, the default type is `int`. If you wanted `x` to be of type
`float32` or `uint8`, you would need to put the type annotation after the
variable name.

In the third example, `x := 3`, we are introduced to a new syntax. This is
called a ["short variable declaration"](https://go.dev/ref/spec#Short_variable_declarations).
Short variable declarations have the advantage that they use slightly less key
strokes than *long* variable declarations. They can also be used in one
situation where long variable declarations cannot be used, but we will look
at that later. For now, you can think of `var x = 3` and `x := 3` as
interchangeable.

In our fourth example, `var x int`, note that we do not assign `x` a value.
Remember that variables are "buckets" that store values. So what do you think
is the value that is currently in the variable `x`?

**Write a program to verify your theory.**

**What about for some of the other data types you learned in the previous
lesson?**

You can declare multiple variables in a single statement, like:

```go
var x, y = 1, 2
var firstName, lastName string
```

These examples are not particularly useful, and its often easier to read your
code if each declaration was on its own line. However, later on, when you call
functions that return multiple values, you will need to do this, for example:

```go
// Read more about this function at https://pkg.go.dev/bufio@go1.23.5#Reader.ReadString
var line, err = reader.ReadString('\n')
```

The function `ReadString` *might* fail, so it returns two values: the first is
a `string`, and the second is an `error`. We will learn more about error
handling in a future lesson, but what is important here is that it is often
necessary to declare more than one variable in a single statement.

## Re-assignment

We've already seen that when we declare a variable, we have the option to
assign it a value or not. However, when you declare a variable isn't the only
time you can assign it a value. An important aspect of variables are that you
can re-assign them at (almost) any time you want.

```go
// chapter-03/exercise-01/main.go

package main

import "fmt"

func main() {
    var greetingCount = 0

    fmt.Println("Hello user!")
    greetingCount = greetingCount + 1

    fmt.Println("Hello user!")
    greetingCount = greetingCount + 1

    fmt.Println("Hello user!")
    greetingCount = greetingCount + 1

    fmt.Printf("You were greeted %d times!\n", greetingCount)
}
```

Notice that in the statement `greetingCount = greetingCount + 1` we are both
*referencing* the variable `greetingCount` and re-assigning it. From this
single line, the compiler will actually generate multiple steps for computer's
processor to do. These process instructions will look something roughly like:

1. Look up the value currently being stored in the variable `greetingCount`.
2. Add `1` to that value, and store in a temporary place (called a *register*).
3. Take that value you calculated and were storing temporarily in the previous
step, and store it back in the variable `greetingCount` (and overwriting the
previous value).

## Constants

## Global Variables
