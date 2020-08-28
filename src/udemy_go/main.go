package main

import "fmt"

// variables can be initialized outside of a function but they cannot be assigned a value with the ':=' operator.
// Will not run.
var deckSize int

// Go is a statically typed language.
func main() {
	// only a variable of type string will be assigned to this variable.
	// long way of writing a variable in Go -> var card string = "Ace of Spades"
	card := "Ace of Spades"
	// the go compiler figures out what type of data you are storing in the variable.
	// only use := when we are initializing a NEW VARIABLE.
	card = "Five of Diamonds"
	deckSize = 5
	fmt.Println(card)
	fmt.Println(deckSize)
}
