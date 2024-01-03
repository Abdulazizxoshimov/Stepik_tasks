package main
import "fmt"
func f(text string){
    //print your code here
    fmt.Println(text)
}
func main(){
	var a string
	fmt.Scan(&a)
	f(a)
}