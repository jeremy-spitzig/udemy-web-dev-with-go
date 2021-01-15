package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe("localhost:8080", r)
}

func index(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(resp, "Welcome!\n")
}
