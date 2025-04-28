package internal

import (
	"bytes"
	"errors"

	"fmt"
	"strconv"

	"github.com/EldoranDev/experiments/tree/main/go/http-server/internal/encoding"
)

type Response struct {
	Version string
	Status  Status
	Header  Header
	Body    bytes.Buffer
	Encoder encoding.Encoder
}

var ErrorEncoderAlreadySet = errors.New("encoder can only be set once")

func (r *Response) ToBytes() []byte {
	res := bytes.Buffer{}

	res.WriteString("HTTP/")
	res.WriteString(r.Version)
	res.WriteString(" ")
	res.WriteString(strconv.FormatInt(int64(r.Status), 10))
	res.WriteString(" ")
	res.WriteString(r.Status.ToString())

	res.WriteString("\r\n")

	data := r.Encoder.Encode(r.Body.Bytes())

	r.Header.Set("Content-Length", strconv.FormatInt(int64(data.Len()), 10))

	if r.Encoder != nil {
		r.Header.Set("Content-Encoding", r.Encoder.Name())
	}

	for header, value := range r.Header {
		res.WriteString(fmt.Sprintf("%s: %s\r\n", header, value))
	}

	res.WriteString("\r\n")

	res.Write(data.Bytes())

	return res.Bytes()
}

func (r *Response) SetContentType(contentType string) {
	r.Header.Set("Content-Type", contentType)
}
