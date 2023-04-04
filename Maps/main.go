package main

import "fmt"

func main() {
	//map declaration
	colors := map[string]string{
		"red":   "danger",
		"green": "nature",
	}

	fmt.Println(colors)

	//Declaring using var keyword
	var colorss = map[string]string{}
	fmt.Println(colorss)

	//declaring using make keyword
	colorsss := make(map[string]string)
	fmt.Println(colorsss)

	//assigning value
	// colors["red"] = "danger"
	// delete(colors, "red")

	// fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, def := range c {
		println("Definition of " + color + " is " + def)
	}
}
