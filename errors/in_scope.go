package main

import "fmt"

func main() {
	//* defines the variable outside the closure
	var msg string
	defer func() {
		//* prints the variable in the deferred closure
		fmt.Println(msg)
	}()
	//* sets the value of the variable
	msg = "Hello world"
}

//* Because msg is defined before the closure, the closure may reference it and as expected, the value of the message will reflect whatever the state of msg is - Hello world - when the deferred function executes.
