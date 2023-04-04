package main

import (
	"fmt"
)

// type person struct {
// 	firstname string
// 	lastname  string
// }

func weird() (string, bool) {
	return "Singye", true
}

func main() {
	// //implicit
	// var name = "Kinchap"
	// // explicit
	// var num int = 10
	// //shorthand
	// num1 := 12
	// fmt.Println(name,num, num1)
	// fmt.Println("hello")
	// hello.Hello()
	// numbers := map[string]string{
	// 	"一": "あなた",
	// 	"二": "君",
	// 	"三": "そして僕",
	// }
	// for num, val := range numbers {
	// 	fmt.Println(num, val)
	// }

	// integer := 100
	// var intPointer = &integer
	// fmt.Println("Printing for the integer")
	// integer := 2
	// var t *int = &integer
	// fmt.Println(t)
	// fmt.Printf("%p", t)

	// fmt.Println("Printing for the string")
	// strings := "Hello"
	// var pointerString *string = &strings
	// fmt.Println(pointerString)
	// fmt.Printf("%p", pointerString)

	// fmt.Println("Printing for the struct")
	// struct1 := person{
	// 	firstname: "Kinley",
	// 	lastname:  "Norbu",
	// }
	// var pointerStruct *person = &struct1
	// fmt.Println(pointerStruct)
	// fmt.Printf("%p", pointerStruct)

	//if, if statements
	str, ok := weird()
	if ok {
		fmt.Println("A")
	} else {
		fmt.Println(str)
	}

	if str == "Singye" {
		fmt.Println("B")
	}

	//if else if else statments
	if ok {
		fmt.Println("C")
	} else if str == "Singye" {
		fmt.Println("D")
	}

	var a int
	for a, b := 1, 1; a < 10; a++ {
		_ = a
		_ = b

		fmt.Println(a)
	}
	fmt.Println(a)
	
}
