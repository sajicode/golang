package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//* the first arg to the Copy func is a value that implements the Writer func & the 2nd arg is a value that implements the Reader func

	//* The Copy func takes a value from a source to an external interface

	io.Copy(os.Stdout, resp.Body)

}
