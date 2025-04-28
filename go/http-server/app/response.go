package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	version  string
	status   Status
	headers  Headers
	encoding *string

	body bytes.Buffer
}

func (r *Response) ToBytes(encoding *string) []byte {
	headers := strings.Builder{}

	var body bytes.Buffer

	if encoding != nil {
		r.headers.Set("Content-Encoding", *encoding)
		body = Encoders[*encoding].Encode(r.body)
	} else {
		body = r.body
	}

	r.headers.Set("Content-Length", strconv.FormatInt(int64(body.Len()), 10))

	for header, value := range r.headers {
		headers.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}

	res := fmt.Sprintf(
		"HTTP/%s %d %s\r\n%s\r\n%s",
		r.version,
		r.status,
		r.status.ToString(),
		headers.String(),
		body.Bytes(),
	)

	return []byte(res)
}

func (r *Response) SetContentType(contentType string) {
	r.headers.Set("Content-Type", contentType)
}

func (r *Response) Write(data []byte) error {
	_, err := r.body.Write(data)

	return err
}
