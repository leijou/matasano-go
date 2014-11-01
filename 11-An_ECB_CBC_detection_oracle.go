package main

import "bytes"
import "fmt"
import "github.com/leijou/matasano-go/analysis"

func main() {
	// Test repeated 16-byte block
	input := []byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

	encrypted, actualmode, err := analysis.RandomOracle(input)
	if err != nil {
		fmt.Println(err)
	} else {
		mode := "CBC"
		for a := 0; a < 32; a++ {
			b := a + 16
			c := b + 16
			if bytes.Compare(encrypted[a:b], encrypted[b:c]) == 0 {
				mode = "ECB"
				break
			}
		}

		fmt.Println("I think it's\t", mode)
		fmt.Println("It's really\t", actualmode)
	}
}
