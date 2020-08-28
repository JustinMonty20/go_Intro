package main

import (
	"fmt"
)

func main() {
	// this is how you allow go to infer the type of the variable you are passing in.
	greeting := "Hello,World"
	fmt.Println(greeting)
}
