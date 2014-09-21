package main

import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"
import "fmt"

func main() {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	if result, err := xor.Apply([]byte(input), []byte(key)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input:", input)
		fmt.Println("Key:", key)
		fmt.Println("Output:", conversion.BytesToHex(result))
	}
}
