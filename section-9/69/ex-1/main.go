package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
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
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	// No problem if there's no session.  User is just not logged in yet.
	u, _ := getUser(resp, req)
	resp.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(resp, "index.gohtml", u)
}

func bar(resp http.ResponseWriter, req *http.Request) {
	u, _ := getUser(resp, req)
	if u == nil {
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(resp, "bar.gohtml", u)
}

func signup(resp http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUsers[un]; ok {
			http.Error(resp, "Username already taken", http.StatusForbidden)
			return
		}

		sID, err := uuid.NewV4()
		if err != nil {
			http.Error(resp, "Failed to create session", http.StatusInternalServerError)
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(resp, cookie)
		dbSessions[cookie.Value] = un
		u := user{un, p, f, l}
		dbUsers[un] = u

		http.Redirect(resp, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(resp, "signup.gohtml", nil)
}
