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
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbPort, dbSchema))
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos/", amigos)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/insert/", insert)
	http.HandleFunc("/read/", read)
	http.HandleFunc("/update/", update)
	http.HandleFunc("/delete/", del)
	http.HandleFunc("/drop/", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println("This is the db!!!!!!!!!!!!!!!", db)
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("This is the db!!!!!!!!!!!!!!!", db)
	_, err := io.WriteString(resp, "at index.")
	check(err)
}

func amigos(resp http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT aName FROM amigos")
	check(err)

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(resp, s)
}

func create(resp http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("CREATE TABLE customer (name VARCHAR(20));")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(resp, "CREATED TABLE customer", n)
}

func insert(resp http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("INSERT INTO customer VALUES (?)")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec("James")
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(resp, "INSERTED RECORD ", n)
}

func read(resp http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM customer")
	check(err)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Println(name)

		fmt.Fprintln(resp, "RETREIVED RECORD: ", name)
	}
}

func update(resp http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("UPDATE customer SET name=?")
	check(err)

	r, err := stmt.Exec("Jimmy")
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(resp, "UPDATED RECORD ", n)
}

func del(resp http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("DELETE FROM customer WHERE name=? OR name=?")
	check(err)

	r, err := stmt.Exec("James", "Jimmy")
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(resp, "DELETED RECORD ", n)
}

func drop(resp http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("DROP TABLE customer")
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(resp, "DROPPED TABLE customer")
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
