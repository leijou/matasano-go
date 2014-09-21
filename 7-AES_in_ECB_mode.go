package main

import "github.com/leijou/matasano-go/conversion"
import "fmt"
import "io/ioutil"
import "crypto/aes"
import "os"

func main() {
	inputfile := "resources/7.txt"

	// Read input file
	input, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode base64
	input = conversion.Base64BytesToBytes(input)

	key := []byte("YELLOW SUBMARINE")
	plaintext := make([]byte, len(input))

	cypher, _ := aes.NewCipher(key)

	for x := 0; x < len(input)-aes.BlockSize; x += aes.BlockSize {
		a := x
		b := x + aes.BlockSize
		cypher.Decrypt(plaintext[a:b], input[a:b])
	}

	fmt.Println(string(plaintext))
}
