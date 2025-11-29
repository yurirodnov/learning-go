package main

import (
	"fmt"
	"strings"
)

func star(n int) {
	// var result string
	// var ch string = "*"

	// for i := 1; i <= n; i++ {
	// 	result += ch

	// }
	// fmt.Print(result)


	// or
	// fmt.Print(strings.Repeat("*", s))

	// or
	var sb strings.Builder
    for i := 0; i < n; i++ {
        sb.WriteString("*")
    }
    fmt.Print(sb.String())

}

func main() {
	star(6)
}