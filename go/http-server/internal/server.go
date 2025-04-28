package internal

import (
	"fmt"
	"io"
	"net"
	"os"
	"regexp"

	"github.com/EldoranDev/experiments/tree/main/go/http-server/internal/encoding"
)

type Handler func(req *Request, res *Response)

type Listener interface {
	Add(method string, path string, handler Handler) error
	Listen()
	Close()
}

type listener struct {
	tcp     net.Listener
	handler map[string]Handler
}

func NewListener() Listener {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	return &listener{
		handler: make(map[string]Handler),
		tcp:     l,
	}
}

func (s *listener) Add(method string, path string, handler Handler) error {
	_, err := regexp.Compile(fmt.Sprintf("^%s-%s$", method, path))
	if err != nil {
		return err
	}

	s.handler[fmt.Sprintf("%s-%s", method, path)] = handler

	return nil
}

func (l *listener) Listen() {
	for {
		conn, err := l.tcp.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			l.handle(conn)
		}(conn)
	}
}

func (l *listener) Close() {
	l.tcp.Close()
}

func (s *listener) handle(conn net.Conn) {
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
			Version: "1.1",
			Header:  make(Header),
		}

		var handler Handler = func(req *Request, res *Response) {
			res.Status = StatusNotFound
		}

		for path, h := range s.handler {
			r := regexp.MustCompile(fmt.Sprintf("^%s$", path))

			matches := r.FindStringSubmatch(fmt.Sprintf("%s-%s", req.Method, req.Path))

			if len(matches) == 0 {
				continue
			}

			req.Params = matches

			handler = h
			break
		}

		handler(req, res)

		conHeader := req.Header.Get("Connection")

		if conHeader != nil {
			close = *conHeader == "close"
			res.Header.Set("Connection", "close")
		}

		var enc *string

		encHeader := req.Header.Get("Accept-Encoding")

		if encHeader != nil {
			enc = encoding.GetEncoding(*encHeader)
		}

		_, err = conn.Write(res.ToBytes(enc))

		if err != nil {
			fmt.Println("Error writing response:", err.Error())
		}
	}
}
