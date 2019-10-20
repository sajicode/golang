package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
		<head><title>Go Web</title></head>
		<body>
		<h1>
			Hello Golang
		</h1></body>
	</html>
	`

	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Service is Down! I repeat, Service is Down!")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	//* the location header must be added before writing the status code
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post := &Post{
		User:    "Heung Son Min",
		Threads: []string{"Fake", "London", "Spurs"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

//* curl -i 127.0.0.1:8080/write|writeheader|redirect

//* The ResponseWriter interface has three methods
// - Write
// - WriteHeader
// - Header

//* The Write method takes in an array of bytes, and this gets writen into the body of the http response.
//* The WriteHeader method is useful for returning status codes
//* The Header method returns a map of headers that can be modified. The modified headers will be in the HTTP response that's sent to the client.
