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

func main() {
	// declaring a struct
	// *1
	// alex := person{"Alex", "Anderson"}
	// *2
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// *3 The ininialized variable comes back as an empty struct
	// var alex person

	// every last line in a nested struct takes a comma

	jim := person{
		firstName: "Alex",
		lastName:  "Morgan",
		contact: contactInfo{
			email:   "alexmorgan@usa.com",
			zipCode: 104121,
		},
	}

	fmt.Printf("%+v", jim)
}
