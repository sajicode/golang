package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing64/counters"
)

func main() {
	// create a variable of the unexported type using the exported New function
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}

//* code doesn't work but is meant to show how we can use
//* an unexported type via an exported func 