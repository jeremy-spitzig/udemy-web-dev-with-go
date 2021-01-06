package main

import (
	"html/template"
	"net/http"
)

type handler int

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func main() {
	http.ListenAndServe(":8080", handler(0))
}
