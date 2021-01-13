package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(resp http.ResponseWriter, req *http.Request) {
	http.SetCookie(resp, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(resp, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(resp, "In Chrome go to: dev tools / application / cookies")
}

func read(resp http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(resp, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(resp, "YOUR COOKIE:", c)
}
