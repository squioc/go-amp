package amp

import (
    "testing"
)

func TestEncode(t *testing.T) {

    messages := make([]Message, 4)
    messages[0] = Message{ 'H', 'e', 'l', 'l', 'o' }
    messages[1] = Message{ 0x20 }
    messages[2] = Message{ 'W', 'o', 'r', 'l', 'd' }
    messages[3] = Message{ 0x33 }

    expected := []byte{ 0x14, 0x00, 0x00, 0x00, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, 0x00, 0x1, 0x20, 0x00, 0x00, 0x00, 0x5, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x00, 0x00, 0x00, 0x1, 0x33 }

    actual, err := Encode(messages)
    if err != nil {
        t.Fatalf("An error occured during the encoding")
    }

    expected_length := len(expected)
    actual_length := len(actual)
    if actual_length != expected_length {
        t.Fatalf("Unexpected length of the frame. expected=%d actual=%d", expected_length, actual_length)
    }

    for i:=0; i<len(expected); i++ {
        if actual[i] != expected[i] {
            t.Fatalf("Mismatching byte postion %d", i);
        }
    }

}
