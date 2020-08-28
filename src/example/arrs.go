package main

import "fmt"

func main() {
	// one way to create an arrray in go.
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	// another way to make an array in go.
	b := [5]int{4, 5, 6, 7, 8}
	fmt.Println(b)

	// another type in go is a slice. exactly an array except you do not have specify a length.
	c := []int{3, 4, 5, 6, 12}
	// := initalizes a new variable. if c already exists and you want to change it, just use an equal sign.
	c = append(c, 14)
	// in the back ground this creates a new array with another element 14 added on to the back of it. sets it to c.
	fmt.Println(c)
}
