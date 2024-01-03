package main

import "fmt"

func convert(n int64)uint16{
    return uint16(n)
}
func main(){
	var a int64
	fmt.Scan(&a)
	fmt.Print(convert(a))
}