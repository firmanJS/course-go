package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Excepted deck length of 20, but got %v",len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Excepted first card of Ace of Spades, but got %v",d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Excepted first card of Four of Clubs, but got %v",d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveTofile("_decktesting")
	
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Excepted 16 card in deck, got %v",len(loadedDeck))
	}

	os .Remove("_decktesting")
}