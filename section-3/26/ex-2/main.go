package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"os"
	"strings"
)

type httpMethod int

type response struct {
	Request       request
	Status        int
	StatusReason  string
	Body          io.Reader
	BodySize      int64
	EffectivePath string
}

func (r response) String() string {
	return fmt.Sprintf("%s %d %s", r.Request.String(), r.Status, r.StatusReason)
}

type request struct {
	Method      httpMethod
	Path        string
	HTTPVersion string
	Headers     map[string]string
}

func (r request) String() string {
	switch r.Method {
	case GET:
		return "GET " + r.Path
	case POST:
		return "POST " + r.Path
	case PUT:
		return "PUT " + r.Path
	case DELETE:
		return "DELETE " + r.Path
	}
	return ""
}

const (
	UNKNOWN httpMethod = iota
	GET
	POST
	PUT
	DELETE
)

func main() {
	listen()
}

func listen() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request := read(conn)
	response := respond(conn, request)
	log.Println(response)
}

func read(conn net.Conn) request {
	i := 0
	scanner := bufio.NewScanner(conn)
	var request request
	request.Headers = make(map[string]string)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			fields := strings.Fields(ln)
			request.Method = httpMethodFromString(fields[0])
			request.Path = fields[1]
			request.HTTPVersion = fields[2]
		} else if ln != "" {
			parts := strings.Split(ln, ":")
			if len(parts) >= 2 {
				request.Headers[parts[0]] = strings.TrimSpace(strings.Join(parts[1:], ":"))
			}
		} else if ln == "" {
			break
		}
		i++
	}
	return request
}

func respond(conn net.Conn, request request) response {
	var response response
	if request.Method != GET {
		// This server only supports GET requests
		response = errorResponse(request, 400, "Bad request")
	} else {
		// Super naïve, and should actually ensure that the path ends up locked to the public folder,
		// but this is just an example server, sooo...
		response = openFile(request, "./public"+request.Path)
	}

	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", response.Status, response.StatusReason)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", response.BodySize)
	fmt.Fprintf(conn, "Content-Type: %s\r\n", determineMimeType(response.EffectivePath))
	fmt.Fprint(conn, "\r\n")
	io.Copy(conn, response.Body)
	return response
}

func httpMethodFromString(httpMethodString string) httpMethod {
	switch httpMethodString {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	}
	return UNKNOWN
}

func openFile(request request, path string) response {
	finfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errorResponse(request, 404, "Not found")
		}
		return errorResponse(request, 500, "Internal error")
	}
	if finfo.IsDir() {
		var separator string = "/"
		if strings.HasSuffix(path, "/") {
			separator = ""
		}
		return openFile(request, path+separator+"index.html")
	}
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errorResponse(request, 404, "Not found")
		}
		return errorResponse(request, 500, "Internal error")
	}
	return response{request, 200, "OK", file, finfo.Size(), path}
}

func determineMimeType(path string) string {
	index := strings.LastIndex(path, ".")
	slashIndex := strings.LastIndex(path, "/")
	if index == -1 || index < slashIndex {
		return "application/octet-stream"
	}
	extension := path[index:]
	detected := mime.TypeByExtension(extension)
	if detected == "" {
		return "application/octet-stream"
	}
	return detected
}

func errorResponse(request request, status int, statusReader string) response {
	return response{request, status, statusReader, strings.NewReader(""), 0, request.Path}
}
