package main

import "encoding/hex"
import "errors"
import "fmt"
import "os"

func fixedXOR(a, b []byte) ([]byte, error) {
	result := make([]byte, len(a))

	if len(a) != len(b) {
		return result, errors.New("Input buffers must match in length")
	}

	for x := 0; x < len(a); x++ {
		result[x] = a[x] ^ b[x]
	}

	return result, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Two hex string arguments required")
		os.Exit(1)
	}

	bytes1, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes2, err := hex.DecodeString(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes3, err := fixedXOR(bytes1, bytes2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(hex.EncodeToString(bytes3))
}
