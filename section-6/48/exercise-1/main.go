package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/assets/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)
}

func foo(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "foo ran")
}

func dog(resp http.ResponseWriter, req *http.Request) {
	t.ExecuteTemplate(resp, "dog.gohtml", nil)
}

func toby(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "./assets/toby.jpg")
}
