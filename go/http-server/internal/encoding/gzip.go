package encoding

import (
	"bytes"
	"compress/gzip"
	"log"
)

type GzipEncoder struct {
}

func (e *GzipEncoder) Name() string {
	return "gzip"
}

func (e *GzipEncoder) Encode(data []byte) bytes.Buffer {
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	if err = zw.Close(); err != nil {
		log.Fatal(err)
	}

	return buf
}
