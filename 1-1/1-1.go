package main

import "encoding/base64"
import "encoding/hex"
import "fmt"
import "os"

func convertHexToBase64(hexstr string) (string, error) {
	bytes, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Single hex string argument required")
		os.Exit(1)
	}

	encoded, err := convertHexToBase64(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(encoded)
}
