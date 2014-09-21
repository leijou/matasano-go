package main

import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"
import "fmt"

func main() {
	inputa := "1c0111001f010100061a024b53535009181c"
	inputb := "686974207468652062756c6c277320657965"

	if a, err := conversion.HexToBytes(inputa); err != nil {
		fmt.Println(err)
	} else if b, err := conversion.HexToBytes(inputb); err != nil {
		fmt.Println(err)
	} else if result, err := xor.ApplyFixed(a, b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input 1:", inputa)
		fmt.Println("Input 2:", inputb)
		fmt.Println("Output:", conversion.BytesToHex(result))
	}
}
