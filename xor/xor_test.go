package xor

import "bytes"
import "encoding/hex"
import "testing"

func TestFixedXOR(t *testing.T) {
	inputa, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	inputb, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	output, err := ApplyFixed(inputa, inputb)

	if err != nil {
		t.Error(err)
	} else if !bytes.Equal(output, expected) {
		t.Errorf("Got %v, expected %v", output, expected)
	}
}
