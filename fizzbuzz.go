package main

import "fmt"

func main() {
    i := 1
    for i <= 15 {
        if (i % 3 == 0 && i % 5 == 0) {
            fmt.Println("Fizzbuzz")
        } else if (i % 3 == 0) {
            fmt.Println("Fizz")
        } else if (i % 5 == 0) {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
        i = i + 1
    }
}


// func fizzBuzz(n int32) {
//     // Write your code here
//     var ade int32 =1
//     for ade <= n {
        
//         if (ade % 3 == 0 && ade % 5 == 0){
//             fmt.Println("FizzBuzz")
//         }else if (ade % 3 == 0){
//             fmt.Println("Fizz")
//         }else if (ade % 5 == 0){
//             fmt.Println("Buzz")
//         } else {
//             fmt.Println(ade)
         
//         }
//            ade++
//     }
 