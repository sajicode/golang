package main

import (
	"fmt"
	"net/http"
	"os"
)

//* Anti pattern method of shutting down servers with a callback

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0) //* we tell the application to exit immdeiately
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

//* Advantages of stopping a server this Way
//* 1. The URL needs to be blocked in production or removed before going to production; having separate code for prod & dev could lead to bugs
//* 2. When the callback URL receives a request, the server shuts down immediately causing actions that might be in progress to be immediately stopped.
//* any data not saved to disk is lost as a consequence.
//* 3. Using a URL sidesteps typical operations tooling such as Ansible, Chef & Puppet or initialization toolchains.
