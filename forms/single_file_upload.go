package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

//* http handler to display & process the form in form.html
func fileForm(w http.ResponseWriter, r *http.Request) {
	//* when the path is accessed with a GET request, display the HTML page & form
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w, nil)
	} else {
		//* get the file handler, header info & error for the form field keyed by its name
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		//* close the form fields before leaving the function
		defer f.Close()
		//* create a local location to save the file, including the file's name
		filename := "tmp/" + h.Filename
		//* create a local file to store the uploaded file
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		//* close the local file before leaving the function
		defer out.Close()

		//* copy the uploaded file to the loal location
		io.Copy(out, f)
		fmt.Fprint(w, "Upload Complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
