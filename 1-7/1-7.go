package main

import "crypto/aes"
import "encoding/base64"
import "fmt"
import "os"
import "io/ioutil"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File to scan required")
		os.Exit(1)
	}

	// Read input file
	inputfile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode base64
	base64.StdEncoding.Decode(inputfile, inputfile)

	key := []byte("YELLOW SUBMARINE")
	plaintext := make([]byte, len(inputfile))

	cypher, _ := aes.NewCipher(key)

	for x := 0; x < len(inputfile)-aes.BlockSize; x += aes.BlockSize {
		a := x
		b := x + aes.BlockSize
		cypher.Decrypt(plaintext[a:b], inputfile[a:b])
	}

	fmt.Println(string(plaintext))

}
