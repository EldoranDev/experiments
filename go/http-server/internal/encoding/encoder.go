package encoding

import "bytes"

var encoders = map[string]Encoder{
	"gzip": &GzipEncoder{},
}

type Encoder interface {
	Encode(data []byte) bytes.Buffer
	Name() string
}

func GetEncoder(header *string) Encoder {
	if header == nil {
		return nil
	}

	encoder := getSupportedEncoding(header)

	if encoder == nil {
		return nil
	}

	if enc, ok := encoders[*encoder]; ok {
		return enc
	}

	return nil
}
