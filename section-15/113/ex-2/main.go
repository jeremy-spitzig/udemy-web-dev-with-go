package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeremy-spitzig/udemy-web-dev-with-go/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
`
	resp.Header().Set("Content-Type", "text/html;charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(s))
}

func getUser(resp http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(resp, "%s\n", uj)
}
