//* buffering a template response
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("simple.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Bad Blood",
		Content: "You! have always worn your flaws upon ur sleeves.",
	}
	//* creates a buffer to store the output of the executed template
	var b bytes.Buffer
	err := t.Execute(&b, p)
	//* handles any errors from template execution
	if err != nil {
		fmt.Fprint(w, "An error occurred.")
		return
	}
	//* copies the buffered output to the response writer
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
