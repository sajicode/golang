//* Resolve URLs using regular expressions

package main

import (
	"fmt",
	"net/http"
	"regexp"
	"strings"
)

func main() {
	rr := newPathResolver()
	//* register paths to functions
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	http.ListenAndServe(":8080", rr)
}

func newPathResolver() *regexResolver {
	return &regexResolver {
		handlers: make(map[string]http.HandlerFunc),
		cache: make(map[string]*regexp.Regexp),
	}
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache map[string]*regexp.Regexp
	//* cache stores compiled regular expressions fo re-use
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

//* ServeHTTP iterates over the regular expressions looking for a match.
//* When the first match is found, it will execute the handler function registerd to that regular expression
//* If more than one regular expression matches an incoming path, the first one added would be the first one checked & used.
func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//* looks up & executes the handler function
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
	//* if no path matches, return a Page not found error
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "King Julien"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "King Julien"
	}
	fmt.Fprint(res, "Goodbye ", name)
}