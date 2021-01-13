package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(resp http.ResponseWriter, req *http.Request) (*user, error) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return nil, err
	}
	if un, ok := dbSessions[cookie.Value]; ok {
		u := dbUsers[un]
		return &u, nil
	}
	return nil, nil
}

func logIn(resp http.ResponseWriter, req *http.Request, u user) {
	sID, err := uuid.NewV4()
	if err != nil {
		http.Error(resp, "Failed to create session", http.StatusInternalServerError)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(resp, cookie)
	dbSessions[cookie.Value] = u.UserName
}

func logOut(resp http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return
	}
	delete(dbSessions, cookie.Value)
	cookie.MaxAge = -1
	cookie.Value = ""
	http.SetCookie(resp, cookie)
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}

func hasRole(resp http.ResponseWriter, req *http.Request, role string) bool {
	u, _ := getUser(resp, req)
	if u == nil {
		return false
	}
	return u.Role == role
}
