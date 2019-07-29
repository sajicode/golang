//* Sample program to show how different functions from the
//* standard library use the io.Writer interface

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//* Create a buffer value and write a string to the buffer
	//* using the write method that implements io.Writer
	//* Below a variable of type Buffer is created & init to zero value
	var b bytes.Buffer
	//* a byte slice is created & init w/ string "Hello"
	//* the byte slice is passed into the write method & becomes
	//* the initial content for the buffer
	b.Write([]byte("Hello "))

	//* Use Fprintf to concatenate a string to the buffer
	//* passing the address of a bytes.Buffer value for io.Writer
	fmt.Fprintf(&b, "World!")

	//* Write the content of the buffer to the stdout device
	//* passing the address of a os.File value for io.Writer
	b.WriteTo(os.Stdout)
}

//* the first value of FprintF must implement the io.Writer interface
//* The Buffer type implements the Write method
