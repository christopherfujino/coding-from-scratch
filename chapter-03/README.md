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

> Write a program to verify your theory.

## Re-assignment

Note that the following program will lead to a compiler error:

```
// chapter-03/exercise-01/main.go

package main

import "fmt"

func main() {
    var x
    x = 1
}
```

Why?

## Constants

## Global Variables
