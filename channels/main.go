package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a slice of popular website urls
	links := []string{
		"http://google.com",
		"http://twitter.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://medium.com",
	}

	//* create a channel
	c := make(chan string)

	for _, link := range links {
		//* we only use the go keyword in front of function calls
		go checkLink(link, c)
	}

	//* setup infinite loop to repeat routines
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		//* send a message into a channel
		c <- link
		return
	}

	fmt.Println(link, " is up")
	c <- link
}
