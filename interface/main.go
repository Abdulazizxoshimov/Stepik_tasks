package main

import (
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
)// все полученные значения имеют тип пустого интерфейса



func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
    if _, ok := value1.(float64); !ok {
        fmt.Printf("value=%v: %T\n", value1, value1)
        return 
    }
    if _, ok := value2.(float64); !ok {
        fmt.Printf("value=%v: %T\n", value2, value2)
        return 
    }
    if v, ok := operation.(string); ok {
         
    
        switch v {
            case "*":
            fmt.Printf("%.4f\n", value1.(float64) * value2.(float64))
            case "/":
            fmt.Printf("%.4f\n", value1.(float64) / value2.(float64))
            case "+":
            fmt.Printf("%.4f\n", value1.(float64) + value2.(float64))
            case "-":
            fmt.Printf("%.4f\n", value1.(float64) - value2.(float64))
            default:
                fmt.Print("неизвестная операция")
        }
    }else {
        fmt.Printf("неизвестная операция")
        return 
    }
        
}
   
func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5.6, "/" 
	
	}