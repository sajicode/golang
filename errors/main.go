package main

import (
	"errors"
	"os"
	"strings"
)

func main() {
	//* get the array of cmd line values
	args := os.Args[1:]
	result, err := Concat(args...)

	if err != nil {
		println(err)
		return
	}
	println(result)
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
