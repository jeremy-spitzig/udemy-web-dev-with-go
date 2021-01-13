package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const sessionLength int = 30

var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

func getUser(resp http.ResponseWriter, req *http.Request) (*user, error) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return nil, err
	}
	if s, ok := dbSessions[cookie.Value]; ok {
		u := dbUsers[s.un]
		return &u, nil
	}
	return nil, nil
}

func logIn(resp http.ResponseWriter, req *http.Request, u user) {
	sID, err := uuid.NewV4()
	if err != nil {
		http.Error(resp, "Failed to create session", http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  sID.String(),
		MaxAge: sessionLength,
	}
	http.SetCookie(resp, cookie)
	dbSessions[cookie.Value] = session{u.UserName, time.Now()}
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
	s := dbSessions[cookie.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func hasRole(resp http.ResponseWriter, req *http.Request, role string) bool {
	u, _ := getUser(resp, req)
	if u == nil {
		return false
	}
	return u.Role == role
}

func updateSession(req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return
	}
	s, ok := dbSessions[cookie.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}
}

func cleanSessions() {
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * time.Duration(sessionLength)) {
		fmt.Println("BEFORE CLEAN")
		showSessions()
		for k, v := range dbSessions {
			if time.Now().Sub(v.lastActivity) > (time.Second * time.Duration(sessionLength)) {
				delete(dbSessions, k)
			}
		}
		fmt.Println("AFTER CLEAN")
		showSessions()
	}
}

func showSessions() {
	fmt.Println("**********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
