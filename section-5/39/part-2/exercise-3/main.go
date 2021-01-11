package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(conn, "HTTP/1.1 200 OK\n")
	io.WriteString(conn, "Content-Length: 20\n")
	io.WriteString(conn, "Content-Type: text/plain\n")
	io.WriteString(conn, "\n")
	io.WriteString(conn, "I see you connected\n")
}
