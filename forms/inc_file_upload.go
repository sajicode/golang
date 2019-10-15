//* Incrementally save uploaded files
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func incFileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_inc.html")
		t.Execute(w, nil)
	} else {
		//* retrieves the multipart reader giving access to the uploaded files and handles any errors
		mr, err := r.MultipartReader()
		if err != nil {
			panic("Failed to read multipart message")
		}
		//* A map to store form field values not relating to files
		values := make(map[string][]string)
		//* 10mb counter for nonfile field size
		maxValueBytes := int64(10 << 20)
		//* continue looping until all of the multipart message has been read
		for {
			//* attempt to read the next part, breaking the loop if the end of the request is reached
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			//* Retrieve the name of the form field, continuing the loop if there's no name
			name := part.FormName()
			if name == "" {
				continue
			}
			//* Retrieve the name of the file if one exists
			filename := part.FileName()
			//* A buffer to read the value of the text field into
			var b bytes.Buffer
			//* if there is no filename, treat as a text field
			if filename == "" {
				//* copy the contents of the part into a buffer
				n, err := io.CopyN(&b, part, maxValueBytes)
				//* if there is an error reading the contents of the part, handle the error
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				//* Using a byte counter, make sure the total size of text fields isn't too large
				maxValueBytes -= n
				if maxValueBytes == 0 {
					msg := "multipart message too large"
					fmt.Fprint(w, msg)
					return
				}
				//* put the contents of the form field into a map for later access
				values[name] = append(values[name], b.String())
				continue
			}
			//* create a location on the filesystem to store the content of a file
			dst, err := os.Create("tmp/" + filename)
			//* close the file when exiting the http handler
			defer dst.Close()
			if err != nil {
				return
			}
			//* as the file content of a part is uploaded, write it to the file
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", incFileForm)
	http.ListenAndServe(":8080", nil)
}
