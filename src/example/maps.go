package main

import "fmt"

func main() {
	//makes a map which is a dictionary or an object. Go map's are like python dictionaries.
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["hexagon"] = 6

	delete(vertices, "hexagon")

	fmt.Println(vertices)
}
