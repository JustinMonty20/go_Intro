package main

import "fmt"

func main() {
	sum := 1
	// here is a loop with no initializer and no post statement. for is alos go's while loop
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
