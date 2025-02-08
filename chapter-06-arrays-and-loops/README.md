# Arrays and Loops

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
    var myFavoriteColors = [3]string{"blue", "green", "yellow"}
    fmt.Printf("One of my favorite colors is %s\n", myFavoriteColors[1])
}
```

If you guessed that this program would print "One of my favorite colors is
blue", then you made a very good, intuitive and yet wrong guess.

|--|--|
|Index|0|1|2|
|Value|"blue"|"green"|"yellow"|
