package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		str := ""

		for a != 0 {
			if (a%10)%2 == 0 && a%10 != 0 {
				str += strconv.FormatUint(uint64(a%10), 10)
			}
			a /= 10
		}
		str2 := ""
		for i := len(str) - 1; i >= 0; i-- {
			str2 += string(str[i])
		}
		ans, _ := strconv.ParseUint(str2, 10, 16)
		if ans == 0 {
			return 100
		}
		return uint(ans)

	}
	var a uint
	fmt.Scan(&a)
	fmt.Print(fn(a))
}
