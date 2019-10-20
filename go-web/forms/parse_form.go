package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Form["first_name"])
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
