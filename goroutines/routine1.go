//* Program that echoes what you type in, but only for 30 seconds.
//* After that, it exists on its own.

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//* call the func echo as a goroutine
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	//* print out message saying we are done sleeping
	fmt.Println("Timed out.")
	//* Exit the program thereby stopping the routine.
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	//* io.Copy copies data to an os.Writer from a os.Reader
	io.Copy(out, in)
}
