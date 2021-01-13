package main

import (
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
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
		pass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(resp, "Failed to encrypt password", http.StatusInternalServerError)
		}
		u := user{un, pass, f, l}
		dbUsers[un] = u

		http.Redirect(resp, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(resp, "signup.gohtml", nil)
}

func login(resp http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(resp, req, "/", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(resp, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(resp, "Username and/or password do not match", http.StatusForbidden)
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
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(resp, "login.gohtml", nil)
}
