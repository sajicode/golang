package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	handler := newHandler() //* get an instance of NewHandler

	ch := make(chan os.Signal)
	//* setup monitoring of operating system signals
	//* the signal package provides a means to get signals from the OS
	//* we set up a channel that receives interrupt & kill signals from the OS so the code can react to them
	signal.Notify(ch, os.Interrupt, os.Kill)

	//* ListenAndServe for both manners & http  blocks execution
	//* To monitor signals, a goroutine needs to run concurrently
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler) //* start the web server
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

//* handler responding to web requests

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

//* ListenForShutdown waits until it receives a signal on the channel
//* once a signal comes in, it sends a message to shutdown on the server
//* this tells the server to stop accepting new connections & shutdown after all current requests are done

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

//* Advantages of this method
//* 1. Allows the current HTTP requests to complete rather than stopping them mid-request
//* 2. Stops listening on the TCP port while completing its requests.
//* This opens up the port for another application to use

//* Disadvantages
//* 1. The manners package works for HTTP connections rather than all TCP connections
//* The manners application will not work on a non-web server
//* 2. If we have long-running socket connections between a server & client apps, the manners package
//* will attempt to wait or interrupt the connections rather than hand them off.

//* test with <http://localhost:8080/?name=hector%20bellerin>
