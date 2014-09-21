package main

import "bufio"
import "bytes"
import "encoding/hex"
import "fmt"
import "os"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File to scan required")
		os.Exit(1)
	}

	// Open input file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		l, _ := hex.DecodeString(line)
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
			fmt.Println(m, "\t", line)
		}

	}
}
