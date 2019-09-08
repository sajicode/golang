//* Recovering from a panic
package main

import (
	"errors"
	"fmt"
)

func main() {
	//* provides a deferred closure to handle panic recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}

//* emits a panic with an error for a body
func yikes() {
	panic(errors.New("something bad happened"))
}

//* The recover function in Go returns a value (interface{}) if a panic has been raised, but in all other cases it returns nil

//* %T gives us information about the type of error, which is the error type created by errors.New
