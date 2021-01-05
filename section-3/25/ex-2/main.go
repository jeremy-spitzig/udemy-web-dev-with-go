package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var db map[string]string

func main() {
	db = make(map[string]string)
	listen()
}

func listen() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Fprint(conn, `
In memory database
Use: 
SET key value    - Save a value to the database
GET key          - Retrieve a value from the database
DEL key          - Remove a key from the database
LST              - Show all key value pairs saved to the database
BYE              - End session
`)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.TrimSpace(scanner.Text())
		parts := strings.Split(ln, " ")
		if len(parts) == 0 {
			continue
		}
		switch {
		case parts[0] == "LST" && len(parts) == 1:
			for key, val := range db {
				fmt.Fprintf(conn, "%s = %s\n", key, val)
			}
		case parts[0] == "GET" && len(parts) == 2:
			fmt.Fprintln(conn, db[parts[1]])
		case parts[0] == "SET" && len(parts) >= 3:
			db[parts[1]] = strings.Join(parts[2:], " ")
		case parts[0] == "DEL" && len(parts) == 2:
			delete(db, parts[1])
		case parts[0] == "BYE" && len(parts) == 1:
			return
		}
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
