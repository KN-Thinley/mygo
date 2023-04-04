package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test1
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %d", len(d))
	}

	//test2
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but found %v", d[0])
	}

	//test3
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but found %v", d[len(d)-1])
	}
}

func TestWriteandRead (t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile(("_decktesting"))

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 card in deck but got %v", len(loadedDeck))
	}
}
