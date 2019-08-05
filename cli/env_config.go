package main

import (
	"fmt"
	"net/http"
	"os"
)

//* retrieving port number from the environment

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

//* not yet familiar with how to use env variables here
