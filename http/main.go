package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//* create a byte slice that would hold all the data passed by the Read function
	bs := make([]byte, 99999)
	//* the 2nd parameter (99999) specifies the capacity? of the byte slice
	//* the reason we set up our byte slice with make & 99999 is bcos the Read function is not set up to resize the size of the slice if full
	//* the response Body has access to the Read interface
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
