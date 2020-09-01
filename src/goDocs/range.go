package main

import "fmt"

// whatever you type in before you declare a struct
// is what you should call it in your code from now on.
type values struct {
	index int
	value int
}

func main() {
	newSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var anotherSlice []values
	for i, v := range newSlice {
		anotherSlice = append(anotherSlice, values{i, v})
	}
	fmt.Println(anotherSlice)
}
