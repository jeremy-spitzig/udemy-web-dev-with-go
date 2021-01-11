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
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)
	http.ListenAndServe(":8080", nil)
}
