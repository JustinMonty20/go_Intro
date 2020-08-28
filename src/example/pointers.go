package main

import "fmt"

// look up what popinters are because i really have no idea wtf it is.
func main() {
	i := 7
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
