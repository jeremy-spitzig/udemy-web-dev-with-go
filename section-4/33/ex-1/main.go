package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type handler int

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func (h handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	t.ExecuteTemplate(resp, "index.gohtml", data)
}

func main() {
	http.ListenAndServe(":8080", handler(0))
}
