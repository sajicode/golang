//* buffering a template response
package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Bad Blood",
		Content: "You! have always worn your flaws upon ur sleeves.",
	}
	//* invokes the template with the page data
	t.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
