package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// declaring a struct
	// *1
	// alex := person{"Alex", "Anderson"}
	// *2
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// *3 The ininialized variable comes back as an empty struct
	// var alex person

	// every last line in a nested struct takes a comma

	alex := person{
		firstName: "Alex",
		lastName:  "Morgan",
		contactInfo: contactInfo{
			email:   "alexmorgan@usa.com",
			zipCode: 104121,
		},
	}

	// * when we change the value of a property in a struct, golang doesn't change the original struct
	// * but creates a copy of the struct and changes the appropriate property value

	//* get the memory address of the value this variable is pointing at
	//* i.e. the memory address where the original struct is located, alexPointer is an address
	alexPointer := &alex
	alexPointer.updateName("Aless")
	alex.print()

}

//* *person => give me the value this memory address is pointing at
//* i.e give me the exact struct this address is pointing at

func (pointerToPerson *person) updateName(newFirstName string) {

	//* *person is a type description which simply means we are working with a pointer to a person
	//* that is why we called updateName on alexPointer & not on Alex

	//* *pointerToPerson is an operator which means we want to manipulate the value the pointer is referencing

	//* the line below turns the receiver struct into the original struct sitting in memory
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {

	fmt.Printf("%+v", p)
}

//* Turn value into address with *pointer
//* Turn address into value with &value
