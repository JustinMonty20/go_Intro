package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(45, 22))
	result, error := sqrt(-9)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}

// here is how you declare functions and parameters of those functions.
func sum(x int, y int) int {
	return x + y
}

/* returning multiple values is useful because there are no built in exceptions in Go.
we need to create our own errors and print them to the console
or return something based on if we are getting an undesired result. */

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
