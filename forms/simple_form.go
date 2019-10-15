package main

import (
	"fmt"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	//* parses a simple form containing only text-based fields
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	//* get the first value for the name field from the form
	name := r.FormValue("name")
}
