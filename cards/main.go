package main

import (
	"fmt"
)

func main() {
	card := newCard() // === var var string = "Ace of Spades"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
