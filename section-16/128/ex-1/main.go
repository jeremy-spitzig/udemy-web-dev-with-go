package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "hello from a docker container")
}
