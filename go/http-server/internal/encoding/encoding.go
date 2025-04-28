package encoding

import (
	"strings"
)

var SupportedEncodings = map[string]bool{
	"gzip": true,
}

var Encoders = map[string]Encoder{
	"gzip": &GzipEncoder{},
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
