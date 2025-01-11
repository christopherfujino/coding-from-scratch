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
	fmt.Printf("i = %d\n", i)
	fmt.Printf("f = %f\n", f)
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
|Go type|Size in bits|Value range|Number of possible values|
|--|--|--|--|
|`int8`|8|-128 to 127|256 (2 ^ 8)|
|`int16`|16|-32,768 to 32,767|65,536 (2 ^ 16)|
|`int32`|32|-2,147,483,648 to 2,147,483,647|4,294,967,296 (2 ^ 32)|
|`int64`|64|-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807|18,446,744,073,709,551,616 (2 ^ 64)|
---

The types that do not have a number at the end, `int` and `uint` have a size
that is most efficient for your computer system.

The following program will print the size of an int on your system:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	var size uintptr = unsafe.Sizeof(x)

	fmt.Printf("The size of an int on this system is %d bytes.\n", size)
	// There are 8 bits in a byte
	fmt.Printf("That is equal to %d bits.\n", size*8)
}
```

As you can tell from the concerning name, package `unsafe` should be used with
care. Go as a whole is a safe language, with pretty good strong guarantees
that a program that will compile will do what the user expects.

However, package `unsafe` has a collection of powerful utilities that could
lead to unexpected results or programs that break when you update your Go
compiler. This program is just a demonstration--in practice, if you need to
to know how big your integer is, choose one of the explicitly-sized integer
types.

### Floats

Floating point numbers--or "floats" for short--are numbers that are not
necessarily whole numbers, such as `1.5` or `0.3333333`.

You might think that floats can do everything an integer can do and more,
so there is no purpose in using integers. Indeed, JavaScript only has a
single numeric type, the equivalent of Go's `float64`. However, there are
two advantages that integers have over floats:

- integer arithmetic is faster for your computer's processor
- when comparing two types that are of the same size, such as `int32` and
  `float32`, while the float can represent larger numbers, at a certain size
  it *loses precision*, which means that is only storing an approximation of
  the correct value, and actual arithmetic might be (slightly) incorrect.
  That is to say, integers are more accurate with large (or very negative)
  numbers.

### Complex Numbers

Complex numbers, or imaginary numbers, are an esoteric mathematical concept
and not generally used outside of scientific computing. This course will not
cover them.

### Misc.

The final official numeric type built into the language is `uintptr`, which
stands for "unsigned integer pointer". It is a special type that is sized
big enough to hold pointers in your system. Pointers will be discussed in a
later chapter. You will generally not need to create or handle your own
`uintptr` variables.

There is also a `byte` type, which is another name for `uint8`, or an unsigned
8-bit integer (there are 8-bits in a byte). Although strictly not needed, the
`byte` type is frequently used to demonstrate that the underlying data is raw
data that came from some other system. This raw data usually has to be
converted into another format before it becomes meaningful to the rest of the
program.

## Exercise 2 - Arithmetic

```go
package main

import (
	"fmt"
)

func main() {
	var x int = 8
	var y int = 3
	fmt.Printf("Addition:\tx + y = %d\n", x+y)
	fmt.Printf("Subtraction:\tx - y = %d\n", x-y)

	fmt.Printf("Multiplication:\tx * y = %d\n", x*y)
	fmt.Printf("Division:\tx / y = %d\n", x/y)

	fmt.Printf("Modulo:\t\tx %% y = %d\n", x%y)
	fmt.Printf("Modulo:\t\t9 %% 3 = %d\n", 9%3)
	fmt.Printf("Modulo:\t\t10 %% 3 = %d\n", 10%3)
	fmt.Printf("Modulo:\t\t11 %% 3 = %d\n", 11%3)
	fmt.Printf("Modulo:\t\t12 %% 3 = %d\n", 12%3)
}
```

### Questions

- What does modulo do?
- What does `\t` do?
- Is there a problem with our arithmetic? Can you fix it?
