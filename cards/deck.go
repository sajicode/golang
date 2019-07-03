package main

import "fmt"

// create a new type of 'deck' which is a slice of strings
// below, we are creating a type deck which is like a slice of strings
type deck []string

// the function below returns a value of type deck (slice of strings)
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// to populate the deck with all possible combinations, we do a nested loop to combine cardSuits & cardValues

	// replace unused values with "_"
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function below is a receiver function
func (d deck) print() {
	// any variable of type "deck" gets access to the "print" method
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (d deck) => d is a variable which represents the actual copy of the deck we're working with.
// deck is a reference to the type we want to attach the print function to.
// In main.go, cards === d
