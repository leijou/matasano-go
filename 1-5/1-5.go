package main

import "fmt"
import "encoding/hex"

func repeatingXOR(plaintext, key []byte) []byte {
	encrypted := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		encrypted[i] = plaintext[i] ^ key[i%len(key)]
	}

	return encrypted
}

func main() {
	plaintext := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	encrypted := repeatingXOR([]byte(plaintext), []byte(key))

	fmt.Println(hex.EncodeToString(encrypted))
}
