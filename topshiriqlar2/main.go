package main
import (
    "fmt"
    "math"
)

func main(){
   var a, b int
    fmt.Scan(&a, &b)
    fmt.Print(int(math.Hypot(float64(a), float64(b))))
}




