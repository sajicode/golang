package main

import (
	"fmt"
	"net/http"
	"time"
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
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alt loop => loop over the channels values and pass each value to l

	//* never attempt to access the same variable inside different go routines
	//* that is why we pass the 'l' value as a parameter into the function literal below
	//* this gives the child routine access to a reference of the original value
	//* we only share information between routines (parent to child) by passing the info as a function argument

	for l := range c {
		//* place sleep statement inside function literal (IIFE)
		go func(link string) {
			//* pause between requests for 5 seconds
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//* do not put sleep statement here
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
