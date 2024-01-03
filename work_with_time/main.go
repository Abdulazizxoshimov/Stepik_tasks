package main
import (
	"bufio"
	"fmt"
	"os"
	"time"
)


func main(){
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	t, _ := time.Parse(time.RFC3339, text.Text())
	a := string(t.Format(time.UnixDate))
	fmt.Println(a)
}