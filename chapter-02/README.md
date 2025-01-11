# Lesson 2 - Primitive Data

## Exercise 1 - Primitive Data Types

In the following exercise, you will be introduced to the basic *primitive*
types in the Go programming language:

```go
// chapter_02/exercise_01/main.go

package main

import "fmt"

func main() {
	var s string = "Hello, world!"
	var r rune = 'c'
	var b bool = true
	var i int = 3
	var f float64 = 3

	fmt.Printf("s = %s\n", s)
	fmt.Printf("r = %c\n", r)
	fmt.Printf("b = %t\n", b)
	fmt.Printf("One third of i = %d\n", i / 2)
	fmt.Printf("One third of f = %f\n", f / 2)
}
```

## Strings and Runes

A `string` is a portion of text. A string is made up of zero or more `rune`s.
In older languages like C, strings are made up of characters, and each
character would be exactly 8-bits in size (capable of representing 256--or 2^8
--possible values). However, Go supports the more robust Unicode system of
representing text, so we call the units of a Go string runes. For now, all you
need to know is that `string`s are made up of `rune`s, and because Go supports
Unicode, you can put text from almost any language into Go strings.

```go
// An example, you don't need to type this one

package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

Go code itself is also unicode, and you can even name your variables (and other
things we will introduce later) with Unicode text:

```go
// An example, you don't need to type this one

package main

import "fmt"

func main() {
    var 世界 string = "world"
    fmt.Printf("Hello, %s\n", 世界)
}
```

## Booleans

A `bool` (short for "boolean") is a value that is either `true` or `false`.
One of the most important things that a computer can do is *conditional
branching*, which means the computer will do something *if* a boolean
condition is true.

We will officially introduce conditional branching in a later chapter, but
this is a preview of what it looks like:

```go
// An example, you don't need to type this

func printMessage(isMorning bool) {
    if isMorning {
        fmt.Println("Good morning!")
    } else {
        fmt.Println("Good day!")
    }
}
```

When this function is called, it will print *either* "Good morning!" or "Good
day!" depending on what the value of `isMorning` is. That is, if `isMorning`
has the value `true` it will print "Good morning!". If it has the value
`false`, it will print "Good day!".

## Numbers

Go supports many numeric data types. At a glance it can seem overwhelming:

- `int`
- `int8`
- `int16`
- `int32`
- `int64`
- `uint`
- `uint8`
- `uint16`
- `uint32`
- `uint64`
- `float32`
- `float64`
- `uintptr`
- `complex64`
- `complex128`

However, in practice, you will only be using a few of these, and the rest
are for special cases.

### Integers

Integers ("int" for short) are positive or negative whole numbers (without a
fractional part or any digits after a decimal point). An unsigned integer ("uint" for short) means the set of integers that are *not* negative--that is,
either positive or zero.

The numbers that come after some of the numeric types represents how many
*bits* are used to store them in memory. A single bit is *binary* digit--that
is, in memory, it is either a 1 or a 0.

---
| |

## Exercise 2 - Arithmetic
