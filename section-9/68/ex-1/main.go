package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(resp http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(resp, "Failed to create session id", http.StatusInternalServerError)
		}
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(resp, cookie)
	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}
	resp.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(resp, "index.gohtml", u)
}

func bar(resp http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(resp, "bar.gohtml", u)
}
