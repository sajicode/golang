package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//* the first arg to the Copy func is a value that implements the Writer func & the 2nd arg is a value that implements the Reader func

	//* The Copy func takes a value from a source to an external interface

	lw := logWriter{}

	//* since lw is a copy of logWriter which implements the Write func,
	//* <look at below comments> we can pass it as the first arg to copy
	io.Copy(lw, resp.Body)

}

//* By passing logWriter as a receiver for func Write,
//* logWriter automatically implements the Writer Interface,
//* which in turn has a Write method that returns int & error

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	//* the Write func is meant to return the length of the byte passed into it

	fmt.Printf("Just wrote %v bytes", len(bs))

	return len(bs), nil
}
