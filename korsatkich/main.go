package main
import "fmt"

func main(){
     a := new(int)
	 b := new(int)
	 test(a,b)
}
func test(x1 *int, x2 *int) {
	// здесь пишите ваш код
    fmt.Print( *x1 * *x2)
    
}