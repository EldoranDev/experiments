package tcp

import (
	"encoding/binary"
	"errors"
	"log"
)

var (
	VERSION     byte = 1
	HEADER_SIZE byte = 4
)

type TCPMessage struct {
	Command byte

	Data []byte
}

func (t *TCPMessage) UnmarshalBinary(bytes []byte) error {
	if bytes[0] != VERSION {
		log.Printf("Data: %x\n", bytes)
		log.Printf("Version Missmatch in Binary Protocol")

		return errors.New("Version Missmatch")
	}

	length := int(binary.BigEndian.Uint16(bytes[2:]))
	data := bytes[HEADER_SIZE : int(HEADER_SIZE)+length]

	t.Command = bytes[1]
	t.Data = data

	return nil
}
