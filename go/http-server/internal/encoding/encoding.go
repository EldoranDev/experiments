package encoding

import (
	"strings"
)

var supportedEncodings = map[string]bool{
	"gzip": true,
}

func getSupportedEncoding(header *string) *string {
	options := strings.SplitSeq(*header, ",")

	for option := range options {
		o := strings.TrimSpace(option)
		if _, ok := supportedEncodings[o]; ok {
			return &o
		}
	}

	return nil
}
