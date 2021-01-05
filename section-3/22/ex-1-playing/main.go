package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if strings.TrimSpace(ln) == "" {
			break
		}
	}
	fmt.Println("Left loop")
	fmt.Fprintln(conn, `HTTP/1.1 200 OK
Server: FunnyServer
Content-Length: 70
Content-Type: text/html

<html>
	<head><title>Hello</title></head>
	<body>Hello</body>
</html>
`)
	fmt.Println("Closing connection")
}
