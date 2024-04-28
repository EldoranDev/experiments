package tcp

import "encoding/binary"

const VERSION byte = 1

const HEADER_LENGTH = 1 + /* Version */ 1 /* Command */ + 2 /* Version */

type TCPMessage struct {
	Command byte

	Data []byte
}

func (t *TCPMessage) MarshalBinary() (data []byte, err error) {
	length := uint16(len(t.Data))

	lengthData := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthData, length)

	b := make([]byte, 0, HEADER_LENGTH+length)

	b = append(b, VERSION)
	b = append(b, t.Command)
	b = append(b, lengthData...)

	return append(b, t.Data...), nil
}
