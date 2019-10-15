package main

//* Using multiple function handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Julian Assange"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

//* checks if there is a path after /goodbye and picks it as name

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Maximus Meridius"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

//* Advantages
//* 1. Well tested and documented as an http package
//* 2. Paths and their mappings to functions are easy to read & follow

//* Disadvantages
//* 1. You can't use different functions for different http methods
//* 2. Wildcard or named sections to a path aren't available
//* 3. Every handler function needs to check for paths outside their bounds and handle returning a Page Not found message.
