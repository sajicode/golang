package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func starWars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "May the Force Be With You!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//* get the name of the function
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/wars", log(starWars))
	server.ListenAndServe()
}

//* logging a function name everytime its called
