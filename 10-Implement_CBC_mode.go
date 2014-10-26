package main

import "github.com/leijou/matasano-go/conversion"
import "fmt"
import "io/ioutil"
import "os"
import "github.com/leijou/matasano-go/blocks"

func main() {
	inputfile := "resources/10.txt"

	// Read input file
	input, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := []byte("YELLOW SUBMARINE")

	// Decode base64
	input = conversion.Base64BytesToBytes(input)

	// Decrypt ECB CBC
	e, err := blocks.NewAES(key)
	if err != nil {
		fmt.Println(err)
	} else {
		iv := make([]byte, e.BlockSize)
		plaintext, err := e.DecryptCBC(input, iv)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(plaintext))
		}
	}
}
