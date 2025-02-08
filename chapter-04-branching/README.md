# Lesson 4: Conditional Branching

Branching is what allows computer programs to make decisions.

```go
// chapter-04/exercise-01/main.go

package main

import "fmt"

func main() {
    var condition = true

    if condition {
        fmt.Println("A special message!")
    }

    fmt.Println("The end.")
}
```

The lines of code in between curly braces (`{` and `}`) are called a **block**.
Blocks are ways to group together lines of code. We've already seen blocks,
because every Go program you've written has at least one block for the `main`
function. So blocks are used to group together the lines of code that make up
a function, or a conditional branch, or several other Go constructs that we
will see later.

It's important to note that if the condition is `false`, then the entire block
of code will be skipped. It's not uncommon in larger programs to have
conditional blocks that are hundreds of lines of code long. This can make it
difficult to read the code, and to tell the difference between what the
computer will do if the condition is `true` or `false`.

Sometimes you want to execute certain code ONLY if the condition was `false`.
In this case, you can use the `else` keyword after the block of code that
follows your `if`:

```go
// An example, you don't need to type in

var isEven bool

if isEven {
    fmt.Println("Your number is even!")
} else {
    fmt.Println("Your number is odd!")
}
```

Note that you can *only* use the `else` keyword after an `if` comparison (and
block of code).

Sometimes there are multiple possible things you might want to do at a certain
point in a program (this is why we use the term "branch", just as there are
many branches coming from a single tree, there could be multiple possible code
paths that your program could take any time that it runs). You can chain
together multiple `if` checks with `else if`. No matter how long your chain of
`else if`s are, once one of them evaluates to true, the rest are skipped.

Before you run the following program, can you predict what it will print?

```go
// chapter-04/exercise-02/main.go

package main

import "fmt"

func main() {
    conditionA := false
    conditionB := true
    conditionC := false

    if conditionA {
        fmt.Println("A")
    } else if conditionB {
        fmt.Println("B")
    } else if conditionC {
        fmt.Println("C")
    } else {
        fmt.Println("else")
    }
}
```

## Comparison Operators

The previous example is a little silly. In particular, since `condition` is
`true`, we know that this program will *always* print `"A special message!"`.
The `if` check is not actually necessary. And if we were to change our
declaration of `condition` to be `var condition = false`, then our program
would *never* print `"A special message!"`.

Boolean values are usually created by comparing two different values. The
following table shows the different **comparison operators** that Go has:

|Operator|Description|Example|
|--|--|--|
|`==`|Are two values equal?|`"Yes" == "yes" // false`|
|`!=`|Are two values not equal (the opposite of `==`)?|`"Yes" != "yes" // true`|
|`>`|Is the value on the left greater than the value on the right?|`2.0 > 2.0 // false`|
|`>=`|Is the value on the left greater than or equal to the value on the right?|`2.0 >= 2.0 // true`|
|`<`|Is the value on the left less than the value on the right?|`2.0 < 2.0 // false`|
|`<=`|Is the value on the left less than or equal to the value on the right?|`2.0 <= 2.0 // true`|

```go
package main

import "fmt"

func main() {
    condition := "Yes" == "yes"

    if condition {
        fmt.Println("Strings in Go are NOT case sensitive")
    } else {
        fmt.Println("Strings in Go ARE case sensitive")
    }
}
```

## Logical Operators

Sometimes the condition that you care about is the combination of multiple
other boolean conditions. For example:

- Is the person at least 18 years old and a US citizen?
- Is the balance on a bank account more than $100,000 or less $100?

To combine two boolean values, Go provides the `&&` and `||` operators, called
**logical AND** and **logical OR** respectively.

```
var isEligible = isEighteen && isCitizen
var exceptionalAccount = moreThan100k || lessThan100
```

## Project

Use the stdlib function `rand.Int()` from `math/rand` to generate a random
integer, and then print a message explaining if the number is even or odd.

<details>
<summary>Hint</summary>
Think back to chapter-02/exercise-02 on arithmetic.
</details>

<details>
<summary>Solution</summary>
```go


```
</details>
