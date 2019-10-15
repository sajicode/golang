package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	//* get the array of cmd line values
	args := os.Args[1:]
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}

//* Concat returns a string and an error
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings attached")
	}

	return strings.Join(parts, " "), nil
}

//* simple error returner.
//* new error returns hexadecimal
