package main

import "bufio"
import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"
import "fmt"
import "os"

func main() {
	inputfile := "resources/4.txt"

	// Open input file
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Loop over file lines
	var best *xor.Decryption
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := conversion.HexToBytes(scanner.Text())

		decryption, _ := xor.BestByteDecryption(line)

		if best == nil || decryption.Score < best.Score {
			best = decryption
		}
	}

	fmt.Println("Line:", conversion.BytesToHex(best.Input))
	fmt.Println("Key:", string(best.Key))
	fmt.Println("Output:", string(best.Output))
}
