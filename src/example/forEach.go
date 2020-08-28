package main

import "fmt"

func main() {
	// running a for each style loop on an array.
	a := []int{4, 5, 6, 7}
	for index, value := range a {
		fmt.Println("index", index, "value", value)
	}

	//	running a for each style loop on a map.
	b := make(map[string]int)
	b["a"] = 5
	b["c"] = 6

	fmt.Println("----------")

	for key, value := range b {
		fmt.Println("key", key, "value", value)
	}

}
