package main
import "fmt"

func main(){
	workArray := [10]uint8{}
	var a uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}
	for i := 0; i < 3; i++ {
		var b, c int
		fmt.Scan(&b, &c)
		workArray[b], workArray[c] = workArray[c], workArray[b]
	}

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}


}