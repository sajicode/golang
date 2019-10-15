//* caching a parsed template for re-use
package main

import (
	"html/template"
	"net/http"
)

//* parses the template when the package is initialized
var t = template.Must(template.ParseFiles("simple.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Think about the bad decisions.",
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

//* this is a subtle,simple way to speed up application responses since we are re-using an already parsed template
