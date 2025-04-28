package internal

import (
	"bytes"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Params  []string
	Header  Header
	Body    bytes.Buffer
}

func NewRequest(data []byte) *Request {
	req := &Request{
		Header: make(Header),
	}

	lines := strings.Split(string(data), "\r\n")

	reqLine := strings.Split(lines[0], " ")

	req.Method = reqLine[0]
	req.Path = reqLine[1]
	req.Version = reqLine[2]

	line := 1

	for lines[line] != "" {
		split := strings.Split(lines[line], ":")
		req.Header.Set(split[0], strings.TrimSpace(split[1]))

		line++
	}

	if line < len(lines) {
		req.Body.Write([]byte(lines[line+1]))
	}

	return req
}
