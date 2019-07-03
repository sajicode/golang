package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of spades")

	// because "cards" is of type "deck", it gets access to the "print" method
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
