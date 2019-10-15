//* recovering from a panic
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	//* runs OpenCSV & handles any errors. This implementation always returns an error.
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	//* uses a deferred function to ensure that a file gets closed
	defer file.Close()

	//* Normally, we would do some extra work here
}

//* OpenCSV opens and preprocesses your file.
func OpenCSV(filename string) (file *os.File, err error) {
	//* the main deferred error handling happens here
	defer func() {
		if r := recover(); r != nil {
			file.Close()

			//* get the error from the panice and pass it back
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	//* opens the data file and handles any errors (such as file not found)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	//* runs our intentionally broken RemoveEmptyLines function
	RemoveEmptyLines(file)

	return file, err
}

func RemoveEmptyLines(f *os.File) {
	//* instead of stripping empty lines, we always fall here
	panic(errors.New("Failed parse"))
}
