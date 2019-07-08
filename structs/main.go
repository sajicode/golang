package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
