package amp

import (
	"strings"
	"testing"
)

func testDecode(t *testing.T) {

	frame := []byte{0x13, 0x00, 0X00, 0x00, 0x03, 0x61, 0x6D, 0x70, 0x00, 0x00, 0x00, 0x01, 0x20, 0x00, 0x00, 0x00, 0x05, 0x72, 0x6f, 0x63, 0x6b, 0x73}

	expected := make([]Message, 3)
	expected[0] = Message{'a', 'm', 'p'}
	expected[1] = Message{' '}
	expected[2] = Message{'r', 'o', 'c', 'k', 's'}

	actual, err := Decode(frame)
	if err != nil {
		t.Fatalf("An error occured during the decoding")
	}

	expected_length := len(expected)
	actual_length := len(actual)
	if actual_length != expected_length {
		t.Fatalf("Unexpected length of the frame. expected=%d actual=%d", expected_length, actual_length)
	}

	for i := 0; i < len(expected); i++ {
		if len(actual[i]) != len(expected[i]) {
			t.Fatalf("Mismatching message length %d", i)
		}
		if strings.EqualFold(string(actual[i]), string(expected[i])) {
			t.Fatalf("Mismatching message %d", i)
		}
	}

}
