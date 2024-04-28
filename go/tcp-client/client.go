package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error Connecting:", err)
	}

	defer conn.Close()

	data := []byte("Hello Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error Sending:", err)
		return
	}
}
