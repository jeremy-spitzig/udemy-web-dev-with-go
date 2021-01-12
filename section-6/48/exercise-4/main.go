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
	http.HandleFunc("/", index)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	t.ExecuteTemplate(resp, "index.gohtml", nil)
}
