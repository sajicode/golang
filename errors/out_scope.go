package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(msg)
	}()
	msg := "Hello world"
}

//* Even though defer is executed after the rest of the function, a closure doesn't have access to variables that are declared after the closure is declared.

//* Because msg isn't declared prior to the deferred function, when the code is evaluated, msg is undefined.
