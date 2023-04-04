package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email  string
	mobile int
}

func main() {
	//creating an instance: Method 1
	// person1 := person{firstName: "Kuenga", lastName: "Tshering", age: 19}
	// fmt.Println(person1)

	//Method II
	// person2 := person{"Kinley", "Norbu",}
	// fmt.Println(person2)

	var kuenga person

	kuenga.firstName = "Kuenga"
	kuenga.lastName = "Tshering"

	kinley := person{
		firstName: "Kinley",
		lastName:  "Thinley",
		contactInfo: contactInfo{
			email:  "ryoutamikasa@gmail.com",
			mobile: 17831049,
		},
	}
	kinley.print()
	kinley.updateName("Sonam")
	kinley.print()

	kinleyPointer := &kinley
	kinley.print()
	kinleyPointer.updateName("Deki")
	kinley.print()

	aSliceToBeChanged()

}

func (p person) print() {
	fmt.Println(p)
}

//using pointer

// even without the pointer variable, it will work because we are expecting pointer person inside the receiver type of the function
func (pPointer *person) updateName(newFirstName string) {
	(*pPointer).firstName = newFirstName
	fmt.Printf("The memory address of pointer variable is %p", pPointer)

}

// incase of slice
func aSliceToBeChanged() {
	slice1 := []string{"Hello", "How", "Are", "You?"}
	fmt.Println(slice1)
	slice1[0] = "Hi"
	fmt.Println(slice1)
}
