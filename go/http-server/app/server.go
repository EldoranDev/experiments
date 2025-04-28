package main

import (
	"fmt"
	"io"
	"net"
	"regexp"
)

type Handler func(req *Request, res *Response)

type Server struct {
	handler map[string]Handler
}

func NewServer() *Server {
	return &Server{
		handler: make(map[string]Handler),
	}
}

func (s *Server) Add(method string, path string, handler Handler) error {
	_, err := regexp.Compile(fmt.Sprintf("^%s-%s$", method, path))
	if err != nil {
		return err
	}

	s.handler[fmt.Sprintf("%s-%s", method, path)] = handler

	return nil
}

func (s *Server) Handle(conn net.Conn) {
	buffer := make([]byte, 8192)
	close := false

	for !close {
		size, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading request", err.Error())
			break
		}

		if size == 0 {
			break
		}

		req := NewRequest(buffer[0:size])
		res := &Response{
			version: "1.1",
			headers: make(Headers),
		}

		var handler Handler = func(req *Request, res *Response) {
			res.status = StatusNotFound
		}

		for path, h := range s.handler {
			r := regexp.MustCompile(fmt.Sprintf("^%s$", path))

			matches := r.FindStringSubmatch(fmt.Sprintf("%s-%s", req.method, req.path))

			if len(matches) == 0 {
				continue
			}

			req.params = matches

			handler = h
			break
		}

		handler(req, res)

		conHeader := req.headers.Get("Connection")

		if conHeader != nil {
			close = *conHeader == "close"
			res.headers.Set("Connection", "close")
		}

		var encoding *string

		encHeader := req.headers.Get("Accept-Encoding")

		if encHeader != nil {
			encoding = GetEncoding(*encHeader)
		}

		_, err = conn.Write(res.ToBytes(encoding))

		if err != nil {
			fmt.Println("Error writing response:", err.Error())
		}
	}
}
