package main

import "fmt"


func main(){
	// var num1, num2 int
    // fmt.Scan(&num1, &num2)
    // for ; num2 >= num1; num2-- {
    //    	fmt.Println(num2)
    // }

	// var x int
    // fmt.Scan(&x)
    // for i := 10; i >= x; i-- {
    //     fmt.Println(i)
    // }

    // sum := 0
    // for i := 0; i <= 30; i += 5 {
    //     sum += i
    // }    
    // fmt.Println(sum)

    var num, base int
    fmt.Scan(&num)
    base = 1
    for i := 2; i <= num; i++ {
        fmt.Println("Before")
        fmt.Println(base)
        fmt.Println("After")
        base *= i
    }
    fmt.Println(base)
}