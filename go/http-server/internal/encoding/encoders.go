package encoding

import "bytes"

type Encoder interface {
	Encode(data bytes.Buffer) bytes.Buffer
}
