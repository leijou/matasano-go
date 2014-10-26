package main

import "github.com/leijou/matasano-go/analysis"
import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"
import "fmt"
import "io/ioutil"
import "os"

func main() {
	inputfile := "resources/6.txt"

	// Read input file
	input, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input = conversion.Base64BytesToBytes(input)

	samples := 15
	ka, err := analysis.BestKeysize(input, samples)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := make([]byte, ka.Keysize)
	blocks := conversion.TransposeBlocks(input, ka.Keysize)
	for i := 0; i < ka.Keysize; i++ {
		decrypt, _ := xor.BestByteDecryption(blocks[i])
		key[i] = decrypt.Key[0]
	}

	if result, err := xor.Apply(input, key); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key:", string(key))
		fmt.Println("Output:", string(result))
	}
}
