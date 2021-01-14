package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_SCHEMA")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbPort, dbSchema))
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(resp http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(resp, "Successfuly completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
