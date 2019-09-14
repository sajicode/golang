//* Using the safely package that handles panics in goroutines
package main

import (
	"errors"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

//* defines a callback that matches the GoDoer type - Go in Practice pp110
func message() {
	println("Inside goroutine")
	panic(errors.New("oops"))
}

func main() {
	//* Instead of <go message>, we use..
	safely.Go(message)
	println("Outside goroutine")
	//* make sure the goroutine has a chance to execute before the program exits
	time.Sleep(1000)
}

// not sure why imported package not detected
