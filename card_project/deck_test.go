package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Wrong deck length, expected 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("The first card is expected to be of Ace of Spades, but got %v", d[0])
	}

	lastCard := d[len(d)-1]
	if lastCard != "Four of Clubs" {
		t.Errorf("The last card is expected to be of Four of Clubs, but got %v", lastCard)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing")

	deck := newDeck()
	deck.saveToFile("_deck_testing")

	loadedDeck := newDeckFromFile("_deck_testing")
	if len(loadedDeck) != 16 {
		t.Errorf("The length of the loaded deck should be equals to 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deck_testing")
}
