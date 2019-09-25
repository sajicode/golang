package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	//* data object to pass to template containing properties to print
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	//* parses a template for later use
	t := template.Must(template.ParseFiles("simple.html"))
	//* writes to http output using template and dataset
	t.Execute(w, p)
}

func main() {
	//* serves the output via simple web server
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

//* the html/template package is used here instead of the text/template package because it's context aware and handles some operations for us
