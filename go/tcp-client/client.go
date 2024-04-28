package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EldoranDev/experiments/go/tcp-client/tcp"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error Connecting:", err)
	}

	defer conn.Close()

	pkg := tcp.TCPMessage{
		Command: ECHO,

		Data: []byte("Hello Server!"),
	}

	data, err := pkg.MarshalBinary()
	if err != nil {
		log.Fatalln("Error marshalling data:", err.Error())
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error Sending:", err)
		return
	}
}
