package conversion

import "testing"
import "bytes"

var pkcstests = map[string][]byte{
	"YELLOW SUBMARINE123":  []byte("YELLOW SUBMARINE123\x01"),
	"YELLOW SUBMARINE":     []byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
	"YELLOW SUBMARI":       []byte("YELLOW SUBMARI\x06\x06\x06\x06\x06\x06"),
	"YELLOW SUBMARINE1234": []byte("YELLOW SUBMARINE1234\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14\x14"),
}

func TestPadPKCS(t *testing.T) {
	blocksize := 20

	for input, expected := range pkcstests {
		output := PadPKCS([]byte(input), blocksize)

		if bytes.Compare(output, expected) != 0 {
			t.Errorf("Got %v, expected %v", output, expected)
		}
	}
}

func TestUnPadPKCS(t *testing.T) {
	for expected, input := range pkcstests {
		output := UnPadPKCS(input)

		if bytes.Compare(output, []byte(expected)) != 0 {
			t.Errorf("Got %v, expected %v", output, expected)
		}
	}
}

func TestDetectPKCS(t *testing.T) {
	for unpadded, padded := range pkcstests {
		if DetectPKCS([]byte(unpadded)) {
			t.Errorf("Got true, expected false on %v", unpadded)
		}

		if !DetectPKCS([]byte(padded)) {
			t.Errorf("Got false, expected true on %v", padded)
		}
	}
}
