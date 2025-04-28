package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

var directoryFlag = flag.String("directory", ".", "directory that contains the files to serve")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	s := NewServer()

	s.Add("GET", "/", func(req *Request, res *Response) {
		res.status = StatusOk
	})

	s.Add("GET", "/echo/([a-zA-Z]+)", func(req *Request, res *Response) {
		res.status = StatusOk
		res.Write([]byte(req.params[1]))
		res.SetContentType("text/plain")
	})

	s.Add("GET", "/files/(.+)", func(req *Request, res *Response) {
		path := filepath.Join(*directoryFlag, req.params[1])

		if _, err := os.Stat(path); err != nil {
			res.status = StatusNotFound
			return
		}

		dat, err := os.ReadFile(path)
		if err != nil {
			res.status = StatusInternalServerError
			return
		}

		res.status = StatusOk
		res.SetContentType("application/octet-stream")
		res.Write(dat)
	})

	s.Add("POST", "/files/(.+)", func(req *Request, res *Response) {
		path := filepath.Join(*directoryFlag, req.params[1])

		err := os.WriteFile(path, req.body.Bytes(), 0644)

		if err != nil {
			res.status = StatusInternalServerError
			return
		}

		res.status = StatusCreated
	})

	s.Add("GET", "/user-agent", func(req *Request, res *Response) {
		ua := req.headers.Get("User-Agent")

		res.status = StatusOk

		if ua != nil {
			res.SetContentType("text/plain")
			res.Write([]byte(*ua))
			return
		}
	})

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			s.Handle(conn)
		}(conn)
	}
}
