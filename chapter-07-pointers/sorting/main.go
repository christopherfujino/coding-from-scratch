package main

import (
	"fmt"
	"math/rand"
	"time"
)

const CAP int = 100000

func main() {
	var a = makeArray()
	fmt.Println(a)
	var start = time.Now()
	a = sort(a)
	var elapsedTime = time.Since(start)
	fmt.Println(a)
	fmt.Printf("Sorting took %v\n", elapsedTime)
}

func makeArray() [CAP]int {
	var a [CAP]int
	for i := 0; i < CAP; i++ {
		a[i] = rand.Int() % 10
	}
	return a
}

func sort(a [CAP]int) [CAP]int {
	for i := 0; i < (CAP - 1); i++ {
		for j := i + 1; j < CAP; j++ {
			if a[i] > a[j] {
				var temp = a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
	return a
}

