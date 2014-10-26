package main

import "github.com/leijou/matasano-go/conversion"
import "fmt"

func main() {
	input := "YELLOW SUBMARINE"
	result := conversion.PadPKCS([]byte(input), 20)

	fmt.Println("Input:", input)
	fmt.Println("Output:", string(result))
}
