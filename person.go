package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// *person Type description - it means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstname string) {
	// *pointer Give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
