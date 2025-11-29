package main

import "fmt"




func main(){
	// var age int

	// fmt.Scanln(&age)

	// if age >= 18{
	// 	fmt.Println("Welcome!")
	// } else {
	// 	fmt.Println("You are not allowed!")
	// }

	// var height int
    // fmt.Scan(&height)
    // if height > 185 {
    //     fmt.Println("Высокий рост")
    // } else if height >= 170 && height <= 185 {
    //     fmt.Println("Средний рост")
    // } else {
    //     fmt.Println("Небольшой рост")
    // }

	x := 12
	res := 0
	switch {
    	case x > 12:
        	res += 1
    	case x > 5 && x < 15:
        	res += 2
        	fallthrough
    	case x > 12:
        	res += 3
    	case x < 20:
        	res += 4
    	default:
       		res = 13
	}
	fmt.Println(res)

}