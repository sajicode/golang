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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		return
	}

	fmt.Println(link, " is up")
}
