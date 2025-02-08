# Arrays and Loops

## Arrays

Thus far, we have learned about the so-called "primitive data types" that Go
supports, such as ints, floats, bools, strings, and runes. These are the basic
building blocks of data used in all Go applications.

However, often programs use large amounts of data. Consider a spreadsheet
program that might have a sheet with thousands of cells. Or an image editing
program, where you could easily have a million pixels loaded into memory. And
then video games take this to the extreme, needing to store 3D character
models, 2D image textures, many audio files, map data, simulation data, etc.
It would not be very feasible for a programmer to assign a unique name to
every single one of these pieces of data.

Programmers need to be able to organize their data. The different organization
strategies that they use are called **data structures**. The simplest and most
important data structure is called an **array**. An array is a collection of
multiple pieces of data of the same type.

Type in the following program. Before you run it, can you guess what it will
do?

```go
package main

import "fmt"

func main() {
    var myFavoriteColors [3]string = [3]string{"blue", "green", "yellow"}
    var aColor string = myFavoriteColors[1]
    fmt.Printf("One of my favorite colors is %s\n", aColor)
}
```

If you guessed that this program would print "One of my favorite colors is
blue", then you have discovered one of the most mistakes programmers make.
But before discuss that, let's look a little closer at the new syntax that
we've introduced:

```go
var myFavoriteColors = [3]string{"blue", "green", "yellow"}
```

Here `[3]string` is an array type. We can read this as "an array of 3
strings." Directly after this we have a curly braces, and three string values
with commas in between them.

Now a look at the next line of code, where we access one of **elements** of
our array:

```go
var aColor string = myFavoriteColors[1]
```

If a variable is like a bucket that holds a particular piece of data, you
can think of an array like a row of buckets all next to each other and
numbered. In programmer-speak, we call each individual bucket in the row
an **element**, and we call the number that identifies each element its
**index**. With this line of code, we are assigning the value contained at the
1-th element into the new variable `aColor` (note that the type of
`myFavoriteColors` is `[3]string` while the type of `aColor` is `string`--we
could say that each element in an array of strings is of type string).

If we want to visualize what our array looks like in memory, it would look
like:

|Index|Value|
|--|--|
|0|"blue"|
|1|"green"|
|2|"yellow"|

Note that the first element in an array is always at index 0. Thus, the second
element of the array is at index 1, and the third at index 2, etc. If you
expected the expression `myFavoriteColors[1]` to give you the first element of
the array, you have made an [off-by-one error]
(https://en.wikipedia.org/wiki/Off-by-one_error), one of the most common in
all of programming.

"If they are so common, why start counting from 0?", you may ask. Today, the
answer is largely historical (almost every other programming language indexes
arrays this way, so doing it differently would confuse experienced
programmers). As for why programming languages first did this, the answer is
that it is the closest to what the machine code looks like when the CPU is
accessing data in an array (see [Zero-based number]
(https://en.wikipedia.org/wiki/Zero-based_numbering#Origin) for more details).

## Uninitialized Arrays

You don't have to assign a value to an array when you declare it. Can you
guess what this program will do before you run it?

```go
package main

import "fmt"

func main() {
    var myFavoriteNumbers [10]int
    fmt.Printf("One of my favorite numbers is %d\n", myFavoriteNumbers[0])
}
```

If you recall from Chapter 3 on variables, every type in Go has a
**zero-value**, or the default value if you do not explicitly initialize a
variable. The zero-value for an array is each element in the array set to the
zero-value for the element type (in this example, the element type is `int`,
so each element is set to `0`).

## Array Length

Go has a special, built-in function called `len()` that will return the
"length" of the item that you pass it. It works like this:

```go
package main

import "fmt"

func main() {
    var myFavoriteNumbers [10]int
    var arrayLength int = len(myFavoriteNumbers)
    fmt.Printf("The array %v has the length %d\n", myFavoriteNumbers, arrayLength)
}
```

## Loops

Consider the following program:

```go
package main

import "fmt"

func main() {
    var myFavoriteNumbers = [6]int{2, 3, 5, 7, 11, 13}

    fmt.Println(myFavoriteNumbers[0])
    fmt.Println(myFavoriteNumbers[1])
    fmt.Println(myFavoriteNumbers[2])
    fmt.Println(myFavoriteNumbers[3])
    fmt.Println(myFavoriteNumbers[4])
    fmt.Println(myFavoriteNumbers[5])
}
```

A program like this is tedious to write. A programmer writing a program like
this will probably copy, paste, and then edit the index number. This is an
error-prone process. We learned in Chapter 5 on functions that functions can
make repetitive code shorter and easier to read. However, if we used a
function here, it actually becomes a longer program:

```go
package main

import "fmt"

func main() {
	var myFavoriteNumbers = [6]int{2, 3, 5, 7, 11, 13}

	printElement(myFavoriteNumbers, 0)
	printElement(myFavoriteNumbers, 1)
	printElement(myFavoriteNumbers, 2)
	printElement(myFavoriteNumbers, 3)
	printElement(myFavoriteNumbers, 4)
	printElement(myFavoriteNumbers, 5)
}

func printElement(array [6]int, index int) {
	fmt.Println(array[index])
}
```

What we would like to do is repeat a line of code several times almost exactly
the same way, except for the index of the array that we access. It turns out
that this pattern is probably the core behavior that programs do the most
often. We could call this pattern of behaviors an **algorithm**, and this
specific algorithm we call "iterating the elements of an array".

Let's describe in human terms what we want this algorithm to do:

1. We want to keep track of where in the array we currently are. Let's declare
a variable `i` (short for index) for this purpose.
2. Let's start our algorithm at the beginning of the array.
3. Do something interesting with the i-th element (in our previous example we
would simply print it, but this is a general algorithm that could do all kinds
of interesting things with the elements).
4. Increment (add `1` to) our tracking variable `i`, so that we can index the
next element.
5. Check if `i` now points outside of the array, which means we are finished.
If not, then go back to step 3.

Because iterating the elements of an array is so common, programming langauges
have built-in syntax for doing it. In go, this would look like:

```go
package main

import "fmt"

func main() {
    var myFavoriteNumbers = [6]int{2, 3, 5, 7, 11, 13}

    for i := 0; i < len(myFavoriteNumbers); i++ {
        fmt.Println(myFavoriteNumbers[i])
    }
}
```

Let's break down this new syntax:

`for` is a keyword that defines a loop in Go. It generally looks like this:

```
for INITIALIZER; COMPARISON; INCREMENT {
  STATEMENTS
}
```

Where:

1. The `INITIALIZER` expression covers steps 1 and 2 above, declaring a
variable to track where in the loop we are, and then initializing it (usually
to `0`). Note that we cannot use the `var` keyword here (it would be a syntax
error), so this is one place where we have to use the short variable
declaration syntax, with the `:=` operator.
2. The `COMPARISON` covers step 5, it evaluates to a boolean `true` or `false`
value. If it is `true`, the computer will execute the loop (at least) one more
time. If it is false, the loop is finished and the computer will jump to after
the code block.
3. The `INCREMENT` does not actually need to increment a variable, but it
usually does. In our example, `i++` is a new syntax that is short for `i = i +
1`. Note that, if we didn't have this expression, our program would loop
forever, always printing the first element.

So let's look again at our algorithm, this time with the actual code:

1. `for i := 0` - let's declare the variable `i` to track our current index in
the `myFavoriteNumbers` array, and initialize it to 0, the beginning.
2. `fmt.Println(myFavoriteNumbers[i])` - do something interesting with the i-th
element of our array.
3. `; i++` - increment our tracking variable by 1 so that we can index the next
element.
4. `; i < len(myFavoriteNumbers)` - as long as our current index is less than
the overall length of our array, continue the loop; otherwise, we're finished.

## Project

Given the array (you'll have to import the `math` package to reference these
constants):

`var interestingNumbers = [6]float32{math.E, math.Pi, math.Phi, math.Sqrt2, math.Ln10, -0.0}`

write a program that uses a for loop to print the *last 5 numbers* of the
array.
