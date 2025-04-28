package main

import "strings"

type Headers map[string]string

func (h *Headers) Set(header string, value string) {
	(*h)[strings.Title(header)] = value
}

func (h *Headers) Get(header string) *string {
	if val, ok := (*h)[strings.Title(header)]; ok {
		return &val
	}

	return nil
}
