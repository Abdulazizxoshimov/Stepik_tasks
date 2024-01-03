package main
import (
	"fmt"
	"errors"
)

func main() {
    // package main уже объявлен, нужные импорты уже есть
    var a, b int 
    fmt.Scan(&a, &b)
    ans, err := divide(a, b)
    if err != nil{
        fmt.Print("ошибка")
    }else {
        fmt.Print(ans)
    }
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Divide by zero")
	} else {
		return a / b, nil
	}
}
