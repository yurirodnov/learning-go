package main

import "fmt"

func main() {
	var list [3]string
	list2 := [3]int{1, 2, 3}

	list[0] = "Hello, "
	list[1] = "Yuri "
	list[2] = "Rodnov"

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	fmt.Println("-----------")

	for i := 0; i < len(list2); i++{
		fmt.Println(list2[i])
	}
}