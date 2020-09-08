package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	// this is how you allow go to infer the type of the variable you are passing in.
	greeting := "Hello,World"
	fmt.Println(greeting)
	fmt.Println(cmp.Diff(greeting,"Hello,World"))
}
