package main

import (
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond", "007"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	updateSession(req)
	u, _ := getUser(resp, req)
	resp.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(resp, "index.gohtml", u)
}

func bar(resp http.ResponseWriter, req *http.Request) {
	if !hasRole(resp, req, "007") {
		http.Error(resp, "007 agents only", http.StatusForbidden)
		return
	}
	updateSession(req)
	u, _ := getUser(resp, req)
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
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(resp, "Username already taken", http.StatusForbidden)
			return
		}

		pass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(resp, "Failed to encrypt password", http.StatusInternalServerError)
		}
		u := user{un, pass, f, l, r}
		dbUsers[un] = u
		logIn(resp, req, u)
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
		logIn(resp, req, u)
		http.Redirect(resp, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(resp, "login.gohtml", nil)
}

func logout(resp http.ResponseWriter, req *http.Request) {
	logOut(resp, req)
	http.Redirect(resp, req, "/", http.StatusSeeOther)
	go cleanSessions()
}
