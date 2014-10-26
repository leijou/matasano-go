package main

import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/blocks"
import "fmt"
import "io/ioutil"
import "os"

func main() {
	inputfile := "resources/7.txt"

	// Read input file
	input, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := []byte("YELLOW SUBMARINE")

	// Decode base64
	input = conversion.Base64BytesToBytes(input)

	// Decrypt ECB
	e, err := blocks.NewAES(key)
	if err != nil {
		fmt.Println(err)
	} else {
		plaintext, err := e.DecryptECB(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(plaintext))
		}
	}
}
