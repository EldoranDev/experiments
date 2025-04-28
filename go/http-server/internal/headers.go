package internal

import "strings"

type Header map[string]string

func (h *Header) Set(header string, value string) {
	(*h)[strings.Title(header)] = value
}

func (h *Header) Get(header string) *string {
	if val, ok := (*h)[strings.Title(header)]; ok {
		return &val
	}

	return nil
}
