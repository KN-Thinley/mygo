package main

import "fmt"

func main() {
	//variable declaration, var variable_name variable_type
	//var card string = "Ace of Spades"
	//var card1 = "Ace of Spades"
	//var1 := "Ace of Spades"

	//since the function returns two data types, provide two different variables to store
	//var answer1 = newCard()

	//fmt.Println(card, card1, var1, answer1)
	slicesAndArrays()

}

// creating a function and declaring return_type
func newCard() string {
	return "Ace of Spades"
}

// Arrays and Slices
func slicesAndArrays() {
	cards := deck{"Ace of Diamonds", "Five of Spades", newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	//for loops (arrays and slices)
	cards.print()

	// Creating a slice using built in function 'make'
	cards = make([]string, 4)
	fmt.Println(cards)
	//Check length and capacity of the slice
	fmt.Println(len(cards), cap(cards))

	//Creating a slice using make and data type byte
	s := make([]byte, 5, 10)
	fmt.Println(s)

	//slicing
	// b := []string{"g", "o", "l", "a", "n", "g"}
	// fmt.Println(b[1:4]) // []byte{"o", "l", "a"}

	//for loops for iterating strings(returns ascii code for every letter in the string)
	var strings string = "This is a string"

	//use '_,' if you don't want the index
	//always use the range keyword before the name of the variable storing iterables
	for _, answer := range strings {
		fmt.Println(answer)
	}
	newDeckofCards := newDeck()
	newDeckofCards.print()

	cardsInHand, remainingCards := newDeckofCards.deal(8)
	cardsInHand.print()
	fmt.Println("Cards in the deck:")
	remainingCards.print()

	newDeckofCards.saveToFile("deck1")

	newDeckofCards = newDeckFromFile("deck1")
	newDeckofCards.print()

	//Shuffle the cards
	cardsToShuffle := newDeck()
	cardsToShuffle.print()

	//shuffle the current deck of cards and then print out the cards
	fmt.Println("The Shuffled Cards start from here:")
	cardsToShuffle.shuffle()
	cardsToShuffle.print()
}
