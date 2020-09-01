package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 0; i < 9; i++ {
		sum += i
	}
	fmt.Println(sum)
}
