package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	// we use := the first time we initialize a variable
	card = "Five of Diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
