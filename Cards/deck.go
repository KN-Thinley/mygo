package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function to separate dealt card and non-dealt card
func (d deck) deal(handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

// Helper function to convert the deck type slice string into string type
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// creating saveToFile() function to write a file
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// reading from a file
func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//convert byteSlice to deck
	str := strings.Split(string(byteSlice), ",")
	return deck(str)
}

func (d deck) shuffle() {
	for i := range d {
		//generate new random index
		newIndex := rand.Intn(len(d) - 1)

		//Swap the positions of the two indexes
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}




