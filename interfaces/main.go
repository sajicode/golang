package main

import "fmt"

//* by declaring the getGreeting() string func inside the Bot struct,
//* any other type that has access to the getGreeting func which also returns a string
//* automatically becomes a member of the Bot interface
//* and so the type has access to every func attached to type bot

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

//* Concrete Types => Values can be created directly off them eg. map, struct englishBot
//* Interface Types => Values cannot be created directly of them e.g bot interface

//* Interfaces are not generic types
//* Interfaces are implicit ie nothing like type implements interface
//* Interfaces are a contract to help us manage types
//* interfaces are tough - difficult to read
