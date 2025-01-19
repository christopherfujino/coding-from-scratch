# Lesson 4: Conditional Branching

```go
// chapter-04/exercise-01/main.go

package main

import "fmt"

func main() {
    var condition = true

    if condition {
        fmt.Println("A special message!")
    }

    fmt.Println("The end.")
}
```

## Comparison Operators

The previous example is a little silly. In particular, since `condition` is
`true`, we know that this program will *always* print `"A special message!"`.
The `if` check is not actually necessary. And if we were to change our
declaration of `condition` to be `var condition = false`, then our program
would *never* print `"A special message!"`.

Boolean values are usually created by comparing two different values. The
following table shows the different **comparison operators** that Go has:

|Operator|Description|Example|
|--|--|--|
|`==`|Are two values equal?|`"Yes" == "yes" // false`|
|`!=`|Are two values not equal (the opposite of `==`)?|`"Yes" != "yes" // true`|
|`>`|Is the value on the left greater than the value on the right?|`2.0 > 2.0 // false`|
|`>=`|Is the value on the left greater than or equal to the value on the right?|`2.0 >= 2.0 // true`|
|`<`|Is the value on the left less than the value on the right?|`2.0 < 2.0 // false`|
|`<=`|Is the value on the left less than or equal to the value on the right?|`2.0 <= 2.0 // true`|

## Logical Operators
