package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	//* determine how much you want to read
	len := r.ContentLength
	//* create a byte array of the content length
	body := make([]byte, len)
	//* call the Read method to read into the byte array
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

//* you need to make a POST request to test this
//* curl -id "first_name=hiro&last_name=nakamura" 127.0.0.1:8080/body
