# Writing Your Own Functions

Type in the following program:

```go
package main

import "fmt"

func main() {
    var taxRate = 0.04

    // Multiple by 100 to turn a decimal into a percentage
    fmt.Printf("The current tax rate is %f%%.\n", taxRate * 100)

    var customer1Name = "Alice"
    var customer2Name = "Bob"
    var customer3Name = "Mallory"

    var customer1SubTotal = 24.0
    var customer2SubTotal = 11.25
    var customer3SubTotal = 110.0

    var customer1TaxOwed = customer1SubTotal * taxRate
    var customer1Total = customer1SubTotal + customer1TaxOwed
    fmt.Printf("%s owes $%f\n", customer1Name, customer1Total)

    var customer2TaxOwed = customer2SubTotal * taxRate
    var customer2Total = customer2SubTotal + customer2TaxOwed
    fmt.Printf("%s owes $%f\n", customer2Name, customer2Total)

    var customer3TaxOwed = customer3SubTotal * taxRate
    var customer3Total = customer3SubTotal + customer3TaxOwed
    fmt.Printf("%s owes $%f\n", customer3Name, customer3Total)
}
```

This was likely tedious to type. Not only is it boring, it's also very
error-prone, since so much of the code is repeated with only variable names
changing. Now consider if you had 300 customers. Or 3,000,000.

The last 9 lines of our `main` function basically repeats a formula for each
of our 3 customers:

1. Multiply the customer's sub-total by our tax rate
2. Add the tax owed to the customer's sub-total
3. Print a message with the customer's name and how much tottal they owe

When we do work repeatedly that follows the same formula and only the specific
values change, this is a good candidate for writing our own function. Type in
the following program:

```go
package main

import "fmt"

var taxRate = 0.04

func main() {
    // Multiple by 100 to turn a decimal into a percentage
    fmt.Printf("The current tax rate is %f%%.\n", taxRate * 100)

    calculateAndPrintTax("Alice", 24.0)
    calculateAndPrintTax("Bob", 11.25)
    calculateAndPrintTax("Mallory", 110.0)
}

func calculateAndPrintTax(name string, subTotal float64) {
    var taxOwed = subTotal * taxRate
    var total = subTotal + taxOwed
    fmt.Printf("%s owes $%f\n", name, total)
}
```

There are several interesting things to note about this program:

1. Unlike all of our `main()` function declarations thus far, our custom
function `calculateAndPrintTax` has something in between the parantheses.
These are called **parameters**, and will be covered in the next section.
1. Recall from chapter 1 that we can do two things with a function, define
it and call it. Notice that we call this function (3 times) before we define
it.
1. Our declaration of the `taxRate` variable has been moved outside of our
`main` function. When a variable is defined outside of any functions, it is
called a **global** variable. Global variables can be shared by any function
in the current package. This is both a strength and weakness of global
variables.

## Parameters

Let's take a closer look at our `calculateAndPrintTax` function:

```go
func calculateAndPrintTax(name string, subTotal float64) {
    var taxOwed = subTotal * taxRate
    var total = subTotal + taxOwed
    fmt.Printf("%s owes $%f\n", name, total)
}
```

This function declaration has two **parameters**:

1. `name string`, the parameter named `name` has a type `string`
1. `subTotal float64`, the parameter named `subTotal` has a type `float64`

Parameters are like variables, except they are never given values when they
are declared, but instead when the function is **called**. So the first time
that we call the function: `calculateAndPrintTax("Alice", 24.0)`, the parameter
`name` is given the value `"Alice"` and the parameter named `subTotal` is given
the value `24.0`.

Nerdy stuff: there are two concepts that are closely-related but often
confused by programmers: **parameters** and **arguments**. Feel free to treat
them as the same--most programmers do! Technically, when you define a function
you give it **parameters**, but when you call it you give it **arguments**.
That is, arguments are the actual values, while the parameters are just the
names and types that are part of the generic function.

## Return Values

So parameters allow a caller (the name for the part of the code that calls a
function) to pass data **into** a function. But functions can also pass data
back to its caller. This is called a **return value**.

You don't need to type this in, but compare this new version of our program:

```go
package main

import "fmt"

var taxRate = 0.04

func main() {
    // Multiple by 100 to turn a decimal into a percentage
    fmt.Printf("The current tax rate is %f%%.\n", taxRate * 100)

    fmt.Printf("%s owes $%f\n", "Alice", calculateTotal(24.0))
    fmt.Printf("%s owes $%f\n", "Bob", calculateTotal(11.25))
    fmt.Printf("%s owes $%f\n", "Mallory", calculateTotal(110.0))
}

func calculateTotal(subTotal float64) float64 {
    var taxOwed = subTotal * taxRate
    return subTotal + taxOwed
}
```

Note that we've changed the first line of our function definition:

```go
func calculateTotal(subTotal float64) float64 {
```

Because our function no longer prints our message, we no longer need a
`name string` parameter. And now, after our parantheses we have an extra
`float64`. This is the return type of the function. We have never had a return
type in any of our previous functions, because they never returned anything
useful. However, if you recall in Chapter 4, we used the `rand.Int()` function
from the `math/rand` package in the standard library. The official
**signature** ([link](https://pkg.go.dev/math/rand#Int)) for this
function is:

```go
func Int() int
```

You can read this as "the function named `Int` does not have any parameters
but it returns an `int`."

For our `calculateTotal` function, we could say "the function named
`calculateTotal` has a single parameter named `subTotal` of type `float64`
and returns a `float64`."

Note that you don't have to give your return value a name, just a type. What
gets returned is whatever the expression is after the `return` keyword.

## Project

Re-write the following program to be simpler by using a custom function:

```go
package main

import "fmt"

func main() {
    var firstName = "Alice"
    var secondName = "Bob"
    var thirdName = "Mallory"

    var firstAge = 17
    var secondAge = 32
    var thirdAge = 60

    var ageDescription string

    if firstAge % 2 == 0 {
        ageDescription = "even"
    } else {
        ageDescription = "odd"
    }
    fmt.Printf("The age of %s is %s.\n", firstName, ageDescription)

    if secondAge % 2 == 0 {
        ageDescription = "even"
    } else {
        ageDescription = "odd"
    }
    fmt.Printf("The age of %s is %s.\n", secondName, ageDescription)

    if thirdAge % 2 == 0 {
        ageDescription = "even"
    } else {
        ageDescription = "odd"
    }
    fmt.Printf("The age of %s is %s.\n", thirdName, ageDescription)
}
```
