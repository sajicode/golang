package main

import (
	"fmt"
	"net/http"
)

type StarWarsHandler struct{}

func (s StarWarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "May the Force Be With You!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protected!")
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	wars := StarWarsHandler{}
	http.Handle("/wars", protect(log(wars)))
	server.ListenAndServe()
}

//* chaining handlers.
