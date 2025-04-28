package main

import (
	"bytes"
	"strings"
)

type Request struct {
	method  string
	path    string
	version string
	params  []string
	headers Headers
	body    bytes.Buffer
}

func NewRequest(data []byte) *Request {
	req := &Request{
		headers: make(Headers),
	}

	lines := strings.Split(string(data), "\r\n")

	reqLine := strings.Split(lines[0], " ")

	req.method = reqLine[0]
	req.path = reqLine[1]
	req.version = reqLine[2]

	line := 1

	for lines[line] != "" {
		split := strings.Split(lines[line], ":")
		req.headers.Set(split[0], strings.TrimSpace(split[1]))

		line++
	}

	if line < len(lines) {
		req.body.Write([]byte(lines[line+1]))
	}

	return req
}
