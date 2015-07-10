package amp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
)

var (
	ErrEmptyInput     = errors.New("Empty frame")
	ErrInvalidVersion = errors.New("Incompatible version")
)

func Decode(frame []byte) ([]Message, error) {
	if len(frame) <= 0 {
		return nil, ErrEmptyInput
	}

	decoder := newDecoder(bytes.NewBuffer(frame))
	version, argc, error := decoder.ReadHeader()
	if error != nil {
		return nil, error
	}

	if version > VERSION {
		return nil, ErrInvalidVersion
	}

	return decoder.ReadBody(argc)
}

type Decoder struct {
	*bytes.Buffer
}

func newDecoder(buffer *bytes.Buffer) *Decoder {
	return &Decoder{Buffer: buffer}
}

func (dec *Decoder) ReadHeader() (int, uint64, error) {
	header, error := dec.Buffer.ReadByte()
	if error != nil {
		return 0, 0, error
	}

	// Extract version and number of messages
	return int(header >> 4 & 0xff), uint64(header & 0xff), nil
}

func (dec *Decoder) ReadBody(argc uint64) ([]Message, error) {
	// compute number of messages
	size := math.Min(float64(argc), float64(dec.Buffer.Len())*1/5)
	messages := make([]Message, uint64(size))

	// extract messages
	for index := 0; index < int(size); index++ {
		// Read message length
		pack := dec.Buffer.Next(4)
		message_len := binary.BigEndian.Uint64(pack)
		// Extract message content
		content := make([]byte, message_len)
		_, error := dec.Buffer.Read(content)
		if error != nil {
			return nil, error
		}

		messages[index] = Message(content)
	}

	return messages, nil
}
