package main

func main() {

	cards := newDeck()

	// because "cards" is of type "deck", it gets access to the "print" method
	// cards.print()
	hand, remainingCards := deal(cards, 5)

	// hand & remainingCards have access to the print function bcos they are of type "deck"
	hand.print()
	remainingCards.print()
}
