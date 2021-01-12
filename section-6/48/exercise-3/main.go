package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	f := http.FileServer(http.Dir("./public"))

	http.HandleFunc("/", index)
	http.Handle("/pics", f)
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	t.ExecuteTemplate(resp, "index.gohtml", nil)
}
