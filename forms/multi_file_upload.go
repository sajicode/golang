package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func multiFileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("multi_form.html")
		t.Execute(w, nil)
	} else {
		//* parse the form in the request & handle any errors
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		//* retrieve a slice, keyed by the input name, containing the files from the Multipart form
		data := r.MultipartForm
		files := data.File["files"]
		//* iterates over the files uploaded to the files field
		for _, fh := range files {
			//* opens a file handler for one of the uploaded files
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			//* create a local file to store the contents of the uploaded file
			out, err := os.Create("tmp/" + fh.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			//* copy the uploaded file to the location on the filesystem
			_, err = io.Copy(out, f)

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", multiFileForm)
	http.ListenAndServe(":8080", nil)
}
