// template inheritance or mixing templates.
package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("./templates/base.html", "./templates/user.html"))
	t["user.html"] = temp
	temp = template.Must(template.ParseFiles("./templates/base.html", "./templates/page.html"))
	t["page.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Mixing Templates Example.",
		Content: "Have fun stormin' da castle",
	}
	t["page.html"].ExecuteTemplate(w, "base", p)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "devil",
		Name: "Devi Ram",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.HandleFunc("/user", displayUser)
	http.ListenAndServe(":8080", nil)
}