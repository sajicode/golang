package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// declaring a struct
	// *1
	// alex := person{"Alex", "Anderson"}
	// *2
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// *3 The ininialized variable comes back as an empty struct
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Morgan"
	fmt.Println(alex)
	// * Below logs the fields and their values
	fmt.Printf("%+v", alex)
}
