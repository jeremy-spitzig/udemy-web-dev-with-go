package main

import "net/http"

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

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}
