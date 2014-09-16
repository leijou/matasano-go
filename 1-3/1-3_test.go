package main

import "encoding/hex"
import "testing"

func TestBestDecryption(t *testing.T) {
	input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expectedkey := byte(88)
	expectedtext := "Cooking MC's like a pound of bacon"

	result := bestDecryption(input)

	if result.Key != expectedkey {
		t.Errorf("Got key %v, expected %v", result.Key, expectedkey)
	}
	if result.Plaintext != expectedtext {
		t.Errorf("Got plaintext %v, expected %v", result.Plaintext, expectedtext)
	}
}
