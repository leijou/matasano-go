package main

import "github.com/leijou/matasano-go/conversion"
import "fmt"

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result, err := conversion.HexToBase64(input)

	if err == nil {
		fmt.Println("Input:", input)
		fmt.Println("Output:", result)
	} else {
		fmt.Println(err)
	}
}
