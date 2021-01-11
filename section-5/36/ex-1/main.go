package main

import (
	"io"
	"net/http"
)

type handler int

func (h handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(resp, "doggy doggy doggy")
	case "/cat":
		io.WriteString(resp, "kitty kitty kitty")
	}
}

func main() {
	http.ListenAndServe(":8080", handler(0))
}
