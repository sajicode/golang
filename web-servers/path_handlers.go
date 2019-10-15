package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

//* Router that uses path matching to map URL paths and HTTP methods
//* to a handler function

func main() {
	//* get an instance of a path based router
	pr := newPathResolver()
	//* map functions to paths
	pr.Add("GET /hello", hello)
	//* asterik is used to specify a wildcard in the path
	pr.Add("* /goodbye/*", goodbye)
	//* set the HTTP server to use our router
	http.ListenAndServe(":8080", pr)
}

func newPathResolver() *pathResolver {
	//* create new initialized path resolver
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

//* Add paths to internal lookup
type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//* construct our method + path to check
	check := req.Method + " " + req.URL.Path
	//* iterate over registered paths
	for pattern, handlerFunc := range p.handlers {
		//* checks whether current path matches a registered one
		if ok, err := path.Match(pattern, check); ok && err == nil {
			//* executes the handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	//* if no path matches, the page wasn't found
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Julian Assange"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

//* checks if there is a path after /goodbye and picks it as name

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Maximus Meridius"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

//* Pros
//* Easy to get started with simple path matching
//* Path package is well tested.

//* Cons
//* The wildcard abilities of the path package are limited
//* E.g. path of foo/* will match foo/bar but not foo/bar/baz
//* Using * for a wildcard stops at the next *
//* Too match foo/bar/baz, you'll need a path like foo/*/*