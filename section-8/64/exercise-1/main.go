package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/clear", clear)
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	var visits int64 = 0
	c, err := req.Cookie("visits")
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		visits, err = strconv.ParseInt(c.Value, 10, 64)
	}
	visits++
	http.SetCookie(resp, &http.Cookie{
		Name:  "visits",
		Value: fmt.Sprintf("%d", visits),
	})
	fmt.Fprintf(resp, "You have visited %d times", visits)
}

func clear(resp http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("visits")
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	c.MaxAge = -1
	http.SetCookie(resp, c)
	http.Redirect(resp, req, "/", http.StatusSeeOther)
}
