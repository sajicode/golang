package main

import (
	"fmt"
	"net/http"
)

type StarWarsHandler struct{}

func (s *StarWarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "May the Force Be With You!")
}

type StartTrekHandler struct{}

func (s *StartTrekHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Live Long & Prosper!")
}

func main() {
	starWars := StarWarsHandler{}
	starTrek := StartTrekHandler{}

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/wars", &starWars)
	http.Handle("/trek", &starTrek)

	server.ListenAndServe()
}
