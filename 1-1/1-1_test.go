package main

import "testing"

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	output, err := convertHexToBase64(input)

	if err != nil {
		t.Error(err)
	} else if output != expected {
		t.Errorf("Got %v, expected %v", output, expected)
	}
}
