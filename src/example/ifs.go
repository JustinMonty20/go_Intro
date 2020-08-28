package main

import (
	"fmt"
)

// simple if statements within go.

func main() {
	x := 3
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 6 {
		fmt.Println("Less than 6")
	}

}
