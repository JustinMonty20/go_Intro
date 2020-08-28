package main

import "fmt"

type person struct {
	name string
	age  int
}

// structs seem more like a java script object than anything else.
func main() {
	firstPerson := person{name: "Justin", age: 23}
	fmt.Println(firstPerson.name)
	fmt.Println(firstPerson.age)
}
