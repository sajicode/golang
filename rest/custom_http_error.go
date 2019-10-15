package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//* A type to hold the information about an error, including metadata about its JSON structure
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

//* JSONError function is similar to http.Error, but the response body is JSON
func JSONError(w http.ResponseWriter, e Error) {
	//* Wrap the Error struct in anonymous struct with error property
	data := struct {
		Err Error `json:"error"`
	}{e}
	//* convert error to JSON & handle an error if one exists
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	//* Set the response MIME to application/json
	w.Header().Set("Content-Type", "application/json")
	//* make sure the http status code is properly set for the error
	w.WriteHeader(e.HTTPCode)
	//* write the JSON body as output
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	//* create an instance of error to use for the error response
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     123,
		Message:  "An Error occurred",
	}
	//* return the error message as JSON when the HTTP handler is called
	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
