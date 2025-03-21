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
