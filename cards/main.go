package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	// cards.saveToFile("my_cards")
	cards.print()
}
