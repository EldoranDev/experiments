package messages

const (
	COMMANDS = iota
	ECHO
)

var (
	VERSION     byte = 1
	HEADER_SIZE byte = 4
)

type TCPMessage struct {
	Command byte
	Data    []byte
}
