package main

import (
	"bufio"
	"fmt"
	"unicode"
	"os"
	
)

func main(){
	text, _:= bufio.NewReader(os.Stdin).ReadString('\n')
	ab := []rune(text)
	a := len(ab)
	

	
	

	if unicode.IsUpper(ab[0]) && (ab[a-1]) == '.'{
		fmt.Print("Right")

	}else {
		fmt.Print("Wrong")
	}
	


}




