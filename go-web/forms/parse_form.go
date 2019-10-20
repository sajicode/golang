package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.FormValue("last_name"))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
