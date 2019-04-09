package main

import "fmt"

// If you are a type in this program with a function called 'getGreeting'
// and you return a string then you are now an honorary member of type 'bot'
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
