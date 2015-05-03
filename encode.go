package amp

import (
    "encoding/binary"
    "bytes"
)

func Encode(messages []Message) ([]byte, error) {
    encoder := newEncoder()

    encoder.WriteHeader(VERSION, len(messages))
    encoder.WriteBody(messages)

    return encoder.Bytes(), nil
}

type Encoder struct {
    bytes.Buffer
}

func newEncoder() *Encoder {
    return &Encoder{}
}

func (enc *Encoder) WriteHeader(version, messages_count int) {
    // Write version and number of messages
    header := (1 << 4 | messages_count) & 0xff
    enc.Buffer.WriteByte(byte(header))
}

func (enc *Encoder) WriteBody(messages []Message) {
    count := len(messages)
    cursor := 5

    // Write messages
    for index := 0; index < count; index++ {
        // Write the length of the message
        length := len(messages[index])
        enc.Buffer.Write(packUint64(length))
        cursor += 4 + length

        // Write the message
        enc.Buffer.Write(messages[index])
    }
}

func (enc *Encoder) Bytes() []byte {
    return enc.Buffer.Bytes()
}

func packUint64(integer int) []byte {
    // pack 64 unsigned integer into 4 bytes
    pack := make([]byte, 4)
    binary.BigEndian.PutUint32(pack, uint32(integer))
    return pack
}
