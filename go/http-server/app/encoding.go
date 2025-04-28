package main

import (
	"bytes"
	"compress/gzip"
	"log"
	"strings"
)

var SupportedEncodings = map[string]bool{
	"gzip": true,
}

var Encoders = map[string]Encoder{
	"gzip": &GzipEncoder{},
}

type Encoder interface {
	Encode(data bytes.Buffer) bytes.Buffer
}

func GetEncoding(header string) *string {
	options := strings.SplitSeq(header, ",")

	for option := range options {
		o := strings.TrimSpace(option)
		if _, ok := SupportedEncodings[o]; ok {
			return &o
		}
	}

	return nil
}

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
