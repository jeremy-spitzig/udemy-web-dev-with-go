package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(resp http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(resp, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	`+v)
}
