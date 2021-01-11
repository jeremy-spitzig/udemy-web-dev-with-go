package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (h hotdog) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "doggy doggy doggy")
}

func (h hotcat) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "kitty kitty kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)
	http.ListenAndServe(":8080", mux)
}
