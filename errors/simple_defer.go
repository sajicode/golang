//* recovering from a panic

package main

import "fmt"

func main() {
	//* defers execution of goodbye
	defer goodbye()

	//* prints a line. this happens before goodbye
	fmt.Println("Hello world")
}

func goodbye() {
	fmt.Println("Goodbye")
}
