package main

import "fmt"

// func main() {
//     fmt.Println("начало")

//     for i := 0; i < 5; i++ {
//         defer fmt.Println(i)
//     }
//     fmt.Println("конец")
// }


func double(a int) {
    fmt.Println(a*2) 
} 

func main() {
    for i := 4; i > 0; i-- {
        defer double(i)
    }
}