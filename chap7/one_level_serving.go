package main

import (
	"fmt"
	"net/http"
	"path"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)
	http.ListenAndServe(":8080", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Gaurav Choudhary"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}