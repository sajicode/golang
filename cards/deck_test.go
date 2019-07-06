package main

import (
	"os"
	"testing"
)

// we pass the go test runner as an argument to test funcs
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndeNewDeckFromFile(t *testing.T) {
	// remove file with test file name
	os.Remove("_decktesting")

	// create new deck and save to disc
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// load deck from disc
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
