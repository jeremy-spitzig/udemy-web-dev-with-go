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

	resp.Header().Set("Content-Type", "text/plain;charset=utf-8")
	resp.WriteHeader(400)
	t.ExecuteTemplate(resp, "index.gohtml", data)
}

func main() {
	http.ListenAndServe(":8080", handler(0))
}
