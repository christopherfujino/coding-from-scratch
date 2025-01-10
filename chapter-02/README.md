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

## Exercise 2 - Arithmetic
