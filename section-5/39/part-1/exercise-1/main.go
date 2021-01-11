package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "This is the index page")
	})
	http.HandleFunc("/dog/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Woof!")
	})
	http.HandleFunc("/me/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Jeremy")
	})
	http.ListenAndServe(":8080", nil)
}
