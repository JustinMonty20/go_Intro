package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// method. takes in a receiver argument. Abs()
// every instance of Point needs to have the Abs() method which is why there is not pointer.
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
// uses a specific instance of point so we need the *
func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

func main() {
	p := Point{10, 12}
	v := Point{10, 12}
	fmt.Println(v.Abs())
	p.Scale(3)
	fmt.Println(p.Abs())
}
