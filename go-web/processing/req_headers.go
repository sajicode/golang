package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	//* get all headers
	h := r.Header
	//* get a particular header
	t := r.Header["Cookie"]
	//* get a particular header 2
	u := r.Header.Get("Accept-Language")
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, t)
	fmt.Fprintln(w, u)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
