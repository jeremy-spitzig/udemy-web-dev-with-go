package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {

	go listen()
	read()
	wg.Wait()
}

func read() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("This is the client.  Sending message.")
	fmt.Fprintln(conn, "My message is really interesting")
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
		wg.Add(1)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("This is the server")
		fmt.Printf("I heard you say: %s\n", ln)
	}
	wg.Done()
	fmt.Println("Connection was closed by client")
}
