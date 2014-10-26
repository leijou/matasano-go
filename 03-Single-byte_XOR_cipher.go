package main

import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"
import "fmt"

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	if a, err := conversion.HexToBytes(input); err != nil {
		fmt.Println(err)
	} else if result, err := xor.BestByteDecryption(a); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input:", input)
		fmt.Println("Key:", result.Key)
		fmt.Println("Output:", string(result.Output))
	}
}
