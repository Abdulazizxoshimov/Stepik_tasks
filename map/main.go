package main

import (
	"fmt"
	"time"
)

func main() {
	d := 0

	a := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&d)
		if value, inMap := a[d]; inMap {
			fmt.Print(value, " ")
		} else {
			a[d] = work(d)
			fmt.Print(a[d], " ")
		}
	}
}

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}
