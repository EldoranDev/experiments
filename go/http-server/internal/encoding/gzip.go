package encoding

import (
	"bytes"
	"compress/gzip"
	"log"
)

type GzipEncoder struct {
}

func (e *GzipEncoder) Encode(data bytes.Buffer) bytes.Buffer {
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	if err = zw.Close(); err != nil {
		log.Fatal(err)
	}

	return buf
}
