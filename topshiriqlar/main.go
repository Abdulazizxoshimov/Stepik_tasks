package main

import "fmt"

func main(){
    var a, sum int
    fmt.Scan(&a)
    for a != 0 {
        sum = sum + (a % 10)
        a = a / 10
    }
    fmt.Print(sum)
}




