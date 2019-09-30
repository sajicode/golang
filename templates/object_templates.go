package main

import (
	"bytes"
	"html/template"
	"net/http"
)

//* variables to hold persistent data shared btw requests
var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("indices.html", "quote.html"))
}

//* types to store data for templates with differing & specific properties
type Page struct {
	Title string
	Content template.HTML
}

type Quote struct {
	Quote, Name string
}

func main() {
	//* populates a data set to supply to template
	q := &Quote{
		Quote: "Illegitimi non carborundum",
		Name: "Socrates",
	}
	//* writes template & data
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	//* stores quote as HTML in global variable
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	//* creates page dataset with quote HTML
	p := &Page{
		Title: "A User",
		Content: qc,
	}
	//* write quote & page to web server output
	t.ExecuteTemplate(w, "indices.html", p)
}