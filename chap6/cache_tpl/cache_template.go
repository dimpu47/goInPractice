package main

import (
	"html/template"
	"net/http"
)

var	t = template.Must(template.ParseFiles("../templates/sample.html"))
 

type Page struct {
	Title, Content string
}


func displayPage(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title:   "Simple Template Example",
		Content: "Simple Content.",
	}

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
