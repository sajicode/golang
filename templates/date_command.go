//* displaying formatted dates in a template

package main

import (
	"html/template"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8">
		<title>Date Example</title>
	</head>
	<body>
		<p>{{.Date | dateFormat "2 Jan, 2006"}}</p>
	</body>
</html>`

//* map go functions to template functions
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

//* function to convert a time to a converted string
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	//* creates a new template.Template instance
	t := template.New("date")
	//* passes additional function in map into template engine
	t.Funcs(funcMap)
	//* parses the template string into the template engine
	t.Parse(tpl)
	//* creates a dataset to pass into template to display
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	//* sends template with data to output response
	t.Execute(res, data)
}

//* serves a template and dataset using a web server
func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}

//? the date and time used in format strings - line 18 are specific i.e. they shouldn't change
