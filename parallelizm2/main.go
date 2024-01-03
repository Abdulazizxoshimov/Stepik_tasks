package main
import (
   "log"
   "time"
)
var d string
func main() {

   // *** тут Ваш код ***
   go work()
   time.Sleep(1 * time.Second)

   if d != "nil" {
      log.Fatal("Error")
   }
   log.Print("All right")
}

func work() {
   time.Sleep(3 * time.Second)
   d = "nil"
}

