package main

import "fmt"

func main() {
    var x uint8 = 0
    var y uint8 = 10
    fmt.Printf("value of x = %d\n", x)
    fmt.Printf("address of x = %v\n", &x)
    fmt.Printf("value of y = %d\n", y)
    fmt.Printf("address of y = %v\n", &y)
}

