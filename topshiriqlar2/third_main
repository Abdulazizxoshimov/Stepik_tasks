package main
import "fmt"

func main(){
    var a string
    fmt.Scan(&a)
    b := []rune(a)
    max := b[0]
    for _, i := range b {
        if  max < i{
            max = i
        }
    }
    fmt.Print(string(max))
}