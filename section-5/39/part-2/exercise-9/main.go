package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var t *template.Template

func init() {
	t = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

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
	var method string
	var path string
	headers := make(map[string]string)
	if scanner.Scan() {
		status := scanner.Text()
		parts := strings.Fields(status)
		method = parts[0]
		path = parts[1]
		fmt.Println(method, path)
	} else {
		return
	}
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		} else {
			parts := strings.SplitN(ln, ":", 2)
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	// for key, value := range headers {
	// 	fmt.Println(key, ": ", value)
	// }
	if headers["Content-Length"] != "" {
		bytes, err := strconv.ParseInt(headers["Content-Length"], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println("Trying to read ", bytes, " bytes")
		// buf := make([]byte, 1)
		// conn.Read(buf)
		// io.ReadFull(conn, buf)
		// fmt.Println(string(buf))
		fmt.Printf("Bytes sent: %d\n", bytes)
	}

	switch {
	case method == "GET" && path == "/":
		getIndex(conn)
	case method == "GET" && path == "/apply":
		getApply(conn)
	case method == "POST" && path == "/apply":
		postApply(conn)
	}
}

func getIndex(conn net.Conn) {
	fmt.Println("getIndex")
	writeResponse(conn, "index.gohtml", nil)
}

func getApply(conn net.Conn) {
	fmt.Println("getApply")
	writeResponse(conn, "apply.gohtml", "GET")
}

func postApply(conn net.Conn) {
	fmt.Println("postApply")
	writeResponse(conn, "apply.gohtml", "POST")
}

func writeResponse(conn net.Conn, template string, data interface{}) {
	file := "templates/" + template
	stat, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(conn, "HTTP/1.1 200 OK\n")
	io.WriteString(conn, fmt.Sprintf("Content-Length: %d\n", stat.Size()))
	io.WriteString(conn, "Content-Type: text/html\n")
	io.WriteString(conn, "\n")

	t.ExecuteTemplate(conn, template, data)
}
