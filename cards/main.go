package main

func main() {

	cards := newDeck()

	// because "cards" is of type "deck", it gets access to the "print" method
	cards.print()
}
