package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	//	 if the top level type is just a name, you can omit it from the elems of the list.
	m := map[string]Vertex{
		"Bell Labs": {40.6843, -74.7891},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

//insert or update an element in a map named m.
// m[key] = elem

//Retrieve an element
// elem = m[key]

//delete an element
//delete (m,key)

//test that a key is present wit a two value assignment
//elem, ok = m[key] where elem is the value and ok is true or false if the element is there.
