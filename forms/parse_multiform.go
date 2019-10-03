//* parsing a form with multiple values for a field
package main

import (
	"fmt"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	//* the maximum memory to store file parts, where the rest is stored to the disk. below is 16mb. default is 32mb
	maxMemory := 16 << 20
	//* parses a multipart form
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		fmt.Println(err)
	}
	//* iterates over all the POST values of the 'names' form field
	for k, v := range r.PostForm["names"] {
		fmt.Println(v)
	}
}
