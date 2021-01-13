package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
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
