package main

import "fmt"

func main() {
	fmt.Printf("What is your name? ")
	var name string
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hello %v!\n", name)
}
