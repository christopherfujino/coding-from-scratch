# Lesson 1 - Hello, World!

## Get and configure your editor

Download Visual Studio Code (commonly called "VS Code"). Note that this is a
different program from the similarly named "Visual Studio".

VS Code is a text editor. This means its primary job is to let you write and
edit the code you write. However, VS Code *can* do other things too,
especially if you install extensions. Extensions give VS Code the ability to
do things like run your code, detect and point out mistakes that you make,
and help you understand the code you write.

Open VS Code. In the menu, navigate to "View" -> "Extensions". There should
now be a pane on the left-hand side of your window titled "EXTENSIONS". In the
search bar, type in "Go". We will be installing the extension for the Go
programming language.

Click on the extension entitled "Go". Ensure that it says "Go Team at Google"
with a checkmark, meaning that the VS Code team has verified the identity of
the publishers of this extension. It is important that you only install VS
Code extensions from authors who you trust.

Once you've verified the author of the extension, you can press the "Install"
button.

Each Go program needs to have its own directory. Create a new directory named `01`.

Type the following program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Save it as `main.go`. The name "main" could be anything, but its important
that the extension of all the source code you write is ".go".
