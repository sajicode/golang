package main

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	//* return an HTTP status 403 with a message
	http.Error(w, "An error occurred", http.StatusForbidden)
}

func main() {
	//* set all paths to serve the Http handler displayError
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
