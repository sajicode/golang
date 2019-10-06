package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

//* a function whose response is true or false if a network timeout caused the error
func hasTimedOut(err error) bool {
	//* use a type switch to detect the type of underlying error
	switch err := err.(type) {
	//* a url error may be caused by an underlying net error that can be checked for a timeout
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	//* look for timeouts detected by the net package
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	//* some errors, w/out a custom type of varaiable to check against, can indicate a timeout
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

//* utilizing the above function
func main() {
	res, err := http.Get("http://example.com/test.zip")
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
