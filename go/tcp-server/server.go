package main

import (
	"fmt"
	"io"
	"net"

	"github.com/EldoranDev/experiments/go/tcp-server/tcp"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error Starting Server:", err)
	}

	defer listener.Close()

	fmt.Println("Server ist listening on port 3000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Errror Accepting Message:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {

		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println("Error Reading:", err)
			return
		}

		var message tcp.TCPMessage

		err = message.UnmarshalBinary(buffer[:n])
		if err != nil {
			return
		}

		fmt.Printf("Received: %v - %s\n", message.Command, message.Data)
	}
}
