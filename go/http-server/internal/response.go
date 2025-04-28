package internal

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/EldoranDev/experiments/tree/main/go/http-server/internal/encoding"
)

type Response struct {
	Version  string
	Status   Status
	Header   Header
	Encoding *string

	Body bytes.Buffer
}

func (r *Response) ToBytes(enc *string) []byte {
	headers := strings.Builder{}

	var body bytes.Buffer

	if enc != nil {
		r.Header.Set("Content-Encoding", *enc)
		body = encoding.Encoders[*enc].Encode(r.Body)
	} else {
		body = r.Body
	}

	r.Header.Set("Content-Length", strconv.FormatInt(int64(body.Len()), 10))

	for header, value := range r.Header {
		headers.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}

	res := fmt.Sprintf(
		"HTTP/%s %d %s\r\n%s\r\n%s",
		r.Version,
		r.Status,
		r.Status.ToString(),
		headers.String(),
		body.Bytes(),
	)

	return []byte(res)
}

func (r *Response) SetContentType(contentType string) {
	r.Header.Set("Content-Type", contentType)
}

func (r *Response) Write(data []byte) error {
	_, err := r.Body.Write(data)

	return err
}
