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
