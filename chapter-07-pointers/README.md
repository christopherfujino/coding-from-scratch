# Pointers

In the last section, we discussed how arrays are the simplest and most
important data structure. One of the reasons why they are so important is
because the way your computer organizes its Random Access Memory (RAM) is in
an array data structure.

The same way that a Go program has **variables** that are like buckets for
storing **values** for your program, a computer has **addresses**

Consider the following program:

```go
package main

func main() {
    var x uint8 = 0
    var y uint8 = 10
}
```

Where does your computer store the data for the variables `x` and `y`? One
possible layout for the computer's memory while running this program could be:

|Address|Value|Note|
|--|--|--|
|9|42|???|
|10|0|`var x`|
|11|10|`var y`|
|12|0|???|

In computer memory, each address references a byte (or 8 bits) of memory. Why
doesn't each address point to each bit? We could build a computer system this
way. However, we rarely want to access only 1 bit of data. It would be less
efficient to find the data we want if we stored each bit of data at its own
address. This means that in cases where we only want to store 1 bit of data,
we will be wasting memory space because the smallest chunk of memory you can
reserve is 1 byte. Why not address more than 1 byte at a time? Because there
are many use cases in which having exactly 1 byte of data is useful, and
if--for example--memory was addressed in 16-bit chunks, all of these uses of
1 byte of data would take twice as much storage.

Let's look again at the above diagram of our memory.
The Go runtime stores the value for variable `x` (currently `0`) at
address 10, and the value for `y` at address 11. Note that the addresses 9 and
12 are not used by our Go code. They could be used by the Go runtime for its
own housekeeping, or it could be used by another program running on the same
computer, or it could be currently free.

It's actually an important aspect of computer security that programs should
only know about the memory addresses that have been given to them by the
operating system. It would be a very dangerous security vulnerability if--for
example--a computer game could read the memory that your browser is currently
using to store the password you are typing in to your banking website.

Also note that if you run this program again a second time, it's very likely
the variable `x` will be stored at a different address. With modern operating
systems, applications request a certain amount of data from the operating
system (in our program, we ask for two 8-bit chunks of memory), and the
operating system decides which free chunks of memory to give it access to. Or,
if we ask for more memory than the system has available, an error will be
thrown and your program will probably crash.

With Go, we can write a program that will actually print out the addresses
that the operating system gave us for each of our variables. The expression
`&x` is read "the address of variable x." Type and run the following program:

```go
package main

import "fmt"

func main() {
    var x uint8 = 0
    var y uint8 = 10
    fmt.Printf("value of x = %d\n", x)
    fmt.Printf("address of x = %d\n", &x)
    fmt.Printf("value of y = %d\n", y)
    fmt.Printf("address of y = %d\n", &y)
}
```

As discussed before, you will almost certainly get different output for the
addresses (the values must be the same, however). Here is one possible output:

```
value of x = 0
address of x = 1374390599690
value of y = 10
address of y = 1374390599691
```

The addresses are very high, because the machine I ran this on has 16
gigabytes of memory. This means it has
[approximately](https://en.wikipedia.org/wiki/Byte#Multiple-byte_units) 16
billion bytes of memory.
