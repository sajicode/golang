package main

import (
	"flag"
	"fmt"
)

//* A cli application using the flag package

//* One way to define a flag.
//* flag.String takes a flag name, default value & description as args
var name = flag.String("name", "World", "A name to say hello to.")

//* Second method using long & short flags
//* the variable type has to be same with the flag type
var spanish bool

//* we declare two variables, one for long & another for short
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {
	//* for the flag values to be in the variables, flag.Parse needs to be run
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}

//* Run without arg with go run flag_cli.go
//* with flag go run flag_cli.go -name "some name"
//* To use the spanish flag, go run flag_cli.go -s/--spanish -name "Naruto"
