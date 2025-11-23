package main

import "fmt"


func main(){
	fmt.Println(5 <= 7 && 5 != 7) // true	

	fmt.Println(7 >= 4 || 5 == 9)	// true

	fmt.Println(!(7 >= 4) && 5 == 5) // false	

	fmt.Println(!(11 > 3 || 7 != 4))	// false

	fmt.Println(10 > 7 && 5 < 6 || 5 == 6) // true

}