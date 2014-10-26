package conversion

import "testing"
import "bytes"

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	output, err := HexToBase64(input)

	if err != nil {
		t.Error(err)
	} else if output != expected {
		t.Errorf("Got %v, expected %v", output, expected)
	}
}

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
