package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go has strict rules for curly braces!")
}

/**

package main
import (
    "fmt"
)
func main()
{
    fmt.Println("Go has strict rules for curly braces!")
}


$ go run curly.go
# command-line-arguments
.\curly.go:5:6: missing function body
.\curly.go:6:1: syntax error: unexpected semicolon or newline before {

**/
