package main

import (
	"fmt"
)



func main(){
	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name)

	var one, two, result int
    fmt.Scan(&one, &two)
    result = one + two
    fmt.Println(result)

}