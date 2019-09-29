package main

import (
	"html/template"
	"net/http"
)

//* A map to store templates in a map of named templates
var t map[string]*template.Template

func init() {
	//* set up the template map
	t := make(map[string]*template.Template)
	//* load templates along with base into the map
	temp := template.Must(template.ParseFiles("base.html", "user.html"))
	t["user.html"] = temp
	temp = template.Must(template.ParseFiles("base.html", "page.html"))
	t["page.html"] = temp
}

//* data objects to pass into templates
type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	//* populates dataset for the page
	p := &Page{
		Title:   "Dooms Day",
		Content: "Breaks my heart into a million pieces",
	}
	//* invokes the template for the page
	t["page.html"].ExecuteTemplate(w, "base", p)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "Bastille",
		Name:     "Dan Smith",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
