package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())

	fmt.Println(rand.Intn(10))
	fmt.Println(add(4, 5))
}

func add(x, y int) (sum int) {
	sum = x + y
	return sum
}

/*
 using the rand package of math gives you a pseudorandom num. To really get a random number you need to seed with an ever changing value.
Like how we are seeding with current time it is always changing.
*/
