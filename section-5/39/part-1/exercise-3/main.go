package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handlerFunc(template string, data interface{}) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		t.ExecuteTemplate(resp, template, data)
	}
}
func main() {
	http.Handle("/", http.HandlerFunc(handlerFunc("index.gohtml", "Index")))
	http.Handle("/dog/", http.HandlerFunc(handlerFunc("dog.gohtml", "Dog")))
	http.Handle("/me/", http.HandlerFunc(handlerFunc("me.gohtml", "Jeremy")))
	http.ListenAndServe(":8080", nil)
}
