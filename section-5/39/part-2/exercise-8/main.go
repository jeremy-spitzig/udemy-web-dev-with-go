package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	if scanner.Scan() {
		status := scanner.Text()
		parts := strings.Fields(status)
		fmt.Println("Method: " + parts[0])
		fmt.Println("Path: " + parts[1])
		fmt.Println("Version: " + parts[2])
	} else {
		return
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(conn, "HTTP/1.1 200 OK\n")
	io.WriteString(conn, "Content-Length: 36\n")
	io.WriteString(conn, "Content-Type: text/html\n")
	io.WriteString(conn, "\n")
	io.WriteString(conn, "<h1>HOLY COW THIS IS LOW LEVEL</h1>\n")
}
