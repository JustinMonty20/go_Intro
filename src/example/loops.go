package main

import "fmt"

// only one kind of loop in go. A for loop. It doubles as a while loop.

func main() {
	//a for loop acting like a while loop in go.
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------------")
	//	a for loop acting like itself in go.
	for i := 5; i <= 10; i++ {
		fmt.Println(i)
	}
}
