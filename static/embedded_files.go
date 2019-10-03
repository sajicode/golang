package main

import (
	"net/http"
)

func main() {
	//* create a box to represent a location on the filesystem
	box := rice.MustFindBox("../files/")
	// an HTTPBox provides files using the http.FileSystem interface
	httpbox := box.HTTPBox()
	//* serve files from the box
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}

//* go in practice pp178
//* didn't finish this
