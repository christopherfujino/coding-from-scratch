package main

import "fmt"

func main() {
	var myFavoriteNumbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%T -> %v\n", myFavoriteNumbers, myFavoriteNumbers)
}
