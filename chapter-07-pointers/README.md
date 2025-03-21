# Pointers

Before running this program, guess what it will do:

```go
package main

import "fmt"

func main() {
	var x = 1
	fmt.Println(x)
	fun(x)
	fmt.Println(x)
}

func fun(x int) {
	x++
}
```

Was your guess correct? Why does it print the same value twice, even though we
increment the value of x inside of the `fun()` function?

The answer is that the variable named `x` in function `main()` and the
argument named `x` in function `fun()` are two separate buckets that just
happen to share the same name (in Go code, anyway). The name `x` in
Go code is just for the benefit of humans reading the code. We could change
these variables to `bert` and `ernie`, and the compiler would (probably)
generate the exact same machine code instructions. All the computer knows about
when it is actually running your program (at least when it comes to finding
these buckets that we call variables or arguments) are **memory addresses**.

You can think of memory addresses like street addresses. Except, humans have
created systems of organizing street addresses, like zip codes, streets,
address numbers, and then sometimes unit numbers, in order to make it easier
to sort and deliver mail.

However, computers are really good at keeping track of large numbers, so
memory addresses are just sequential numbers, each number referencing a byte
(8 bits) of memory. Go has an operator `&`, which is pronounced "address-of
operator" which will yield the actual address that the computer is storing the
contents of a variable at. Type in and run the following program:

```go
package main

import "fmt"

func main() {
	var x = 1
	fmt.Printf("The value of x in main is %d\n", x)
	fmt.Printf("The address of x in main is %d\n", &x)
	fun(x)
}

func fun(x int) {
	x++
	fmt.Printf("The value of x in fun is %d\n", x)
	fmt.Printf("The address of x in fun is %d\n", &x)
}
```

As discussed before, you will almost certainly get different output for the
addresses (the values must be the same, however). Here is one possible output:

```
The value of x in main is 1
The address of x in main is 824634843200
The value of x in fun is 2
The address of x in fun is 824634843208
```

The addresses are very high, because the machine I ran this on has 16
gigabytes of memory. This means it has
[approximately](https://en.wikipedia.org/wiki/Byte#Multiple-byte_units) 16
billion bytes of memory.

Since we have two separate `x` buckets in our program, what happens when we
call the function `fun(x)`? The Go runtime *copies* the value in our `main()`
function's `x` bucket into the `fun()` function's `x` bucket. In nerdspeak,
this is called "pass by value" or
"[call by value](https://en.wikipedia.org/wiki/Evaluation_strategy#Call_by_value)."

What happens if you want two functions to share the same bucket of memory?
This is where a Go feature called **pointers** become useful. A pointer is a
special variable that holds the address of *another* variable. Let's look at an
updated version of our program:

```go
package main

import "fmt"

func main() {
	var x = 1
	fmt.Println(x)
	fun(&x)
	fmt.Println(x)
}

func fun(x *int) {
	(*x)++
}
```

First, note that type signature of our `fun()` function has slightly changed:

```go
func fun(x *int)
```

Here, `*int` means a pointer to an `int`. Next, note that when we call this
function, we use the previously covered `&` address-of operator:

```go
fun(&x)
```

Now, instead of copying the value in the `x` variable to the `fun()` function,
it instead looks up *where* the `x` variable is being stored, and passes that
address to the `fun()` function.

Finally, when we want to increment our variable in the `fun()` function, we
put an `*` before the pointer name. This means "access the value in the
storage pointed to, not the address":

```go
(*x)++
```

Here the parentheses are actually optional, but they were added for clarity.
**First** we want to access the value `x` is pointing to, **then** we want to
increment that value by 1. We could have also wrote this line like:

```go
*x = *x + 1
```

If we tried this without the `*` operators, like:

```go
x = x + 1
```

We would get the error:

```
./main.go:13:6: invalid operation: x + 1 (mismatched types *int and untyped int)
```

In Go, you cannot do arithmetic with memory addresses.

## Scanf

Let's look at a more real-world example of using pointers:

```go
package main

import "fmt"

func main() {
	var name string

	fmt.Printf("What is your name? ")
	fmt.Scan(&name)
	fmt.Printf("Hello %s!\n", name)
}
```

The function [`fmt.Scan()`](https://pkg.go.dev/fmt#Scan) will read text in from
standard in (also known as STDIN, this will usually be reading keyboard text
from your terminal). There are a few quirks to this function (and the related
functions `fmt.Scanf()` and `fmt.Scanln()`):

* it takes one or more variable pointers as its arguments
