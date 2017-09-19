package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// notes about pointer:
// TURN value INTO address WITH <---- &value
// TURN address INTO value WITH <---- *address

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "email@email.com",
			zipCode: 94039,
		},
	}
	// version 1
	// jimpointer := &jim // get access to memory address
	// jimpointer.updateName("jimmy")
	// jimpointer.print()

	// shortcut version
	jim.updateName("jimmy")
	jim.print()
}

// *person -> the star is not an "operator" -> it means we are working with a pointer to a person
func (personPointer *person) updateName(newFirstName string) {
	// *personPointer -> give me the value of the pointer
	(*personPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
