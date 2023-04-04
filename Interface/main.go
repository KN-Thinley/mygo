package main

import (
	"fmt"
)

type englishBot struct{}
type dzongkhaBot struct{}

type bot interface {
	greet() string
}

func main() {
	eb := englishBot{}
	db := dzongkhaBot{}

	printGreeting(eb)
	printGreeting(db)
}

func (englishBot) greet() string {
	return "Hello"
}

func (dzongkhaBot) greet() string {
	return "Kuzuzangpo"
}

func printGreeting(b bot) {
	fmt.Println(b.greet())
}
