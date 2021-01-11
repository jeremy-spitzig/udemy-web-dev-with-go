package main

import (
	"io"
	"net/http"
)

func d(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "doggy doggy doggy")
}

func c(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "kitty kitty kitty")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat/", http.HandlerFunc(c))
	http.ListenAndServe(":8080", nil)
}
