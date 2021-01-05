package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	go listen()
	read()
}

func read() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("This is the client")
	fmt.Println(string(bs))
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
	fmt.Println("This is the server")
	fmt.Fprintln(conn, "Sending some data")
}
