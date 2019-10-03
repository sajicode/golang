import (
	"flag"
	"net/http"
	"text/template"
)

//* passing a url location to a template
var t *template.Template

//* gets the loccation of the static files from the app arguments
var l = flag.String("location", "http://localhost:8080", "a location")

var tpl = `<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8">
		<title>A Demo</title>
		<link rel="stylesheet" href="{{.Location}}/styles.css">
	</head>
	<body>
		<p>A Simple Demo</p>
	</body>
</html>`

//* The path to the css is relative to the location

//* An http handler passing the location into the template
func servePage(res http.ResponseWriter, req *http.Request) {
	data := struct{ Location *string }{
		Location: l,
	}
	t.Execute(res, data)
} 