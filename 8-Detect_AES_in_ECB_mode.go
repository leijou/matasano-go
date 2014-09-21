package main

import "github.com/leijou/matasano-go/conversion"
import "fmt"
import "os"
import "bufio"
import "bytes"

func main() {
	inputfile := "resources/8.txt"

	// Open input file
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		l, _ := conversion.HexToBytes(line)
		m := 0
		for x := 0; x < len(l); x += 16 {
			matches := 0
			for i := 0; i < len(l); i += 16 {
				if bytes.Equal(l[x:x+16], l[i:i+16]) {
					matches++
				}
			}
			if matches > 1 {
				m++
			}

		}

		if m > 0 {
			fmt.Println("Detected", m, "repetitions:", line)
		}
	}
}
