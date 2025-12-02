package main

import "fmt"

func main() {
	a := [5]int{2, 0, 5, 12, 19}

	var s []int = a[0:3]

	s[0] = 1000

	fmt.Println(a)

}