package main

import "encoding/hex"
import "fmt"
import "os"
import "sort"

var engFrequencies = map[byte]float32{'E': 0.1270, 'T': 0.0906, 'A': 0.0817, 'O': 0.0751, 'I': 0.0697, 'N': 0.0675, 'S': 0.0633, 'H': 0.0609, 'R': 0.0599, 'D': 0.0425, 'L': 0.0403, 'C': 0.0278, 'U': 0.0276, 'M': 0.0241, 'W': 0.0236, 'F': 0.0223, 'G': 0.0202, 'Y': 0.0197, 'P': 0.0193, 'B': 0.0129, 'V': 0.0098, 'K': 0.0077, 'J': 0.0015, 'X': 0.0015, 'Q': 0.0010, 'Z': 0.0007}

// Decryption result
type decrypt struct {
	Key       byte
	Plaintext string
	Score     float32
}

// List of decryption results (sortable by weight)
type decryptList []*decrypt

func (a decryptList) Len() int           { return len(a) }
func (a decryptList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a decryptList) Less(i, j int) bool { return a[i].Score < a[j].Score }

func singleXOR(a []byte, b byte) (out []byte) {
	out = make([]byte, len(a))

	for x := 0; x < len(a); x++ {
		out[x] = a[x] ^ b
	}

	return
}

func letterFrequency(s []byte) map[byte]float32 {
	letters := map[byte]float32{'A': 0, 'B': 0, 'C': 0, 'D': 0, 'E': 0, 'F': 0, 'G': 0, 'H': 0, 'I': 0, 'J': 0, 'K': 0, 'L': 0, 'M': 0, 'N': 0, 'O': 0, 'P': 0, 'Q': 0, 'R': 0, 'S': 0, 'T': 0, 'U': 0, 'V': 0, 'W': 0, 'X': 0, 'Y': 0, 'Z': 0}

	sum := float32(0)

	for x := 0; x < len(s); x++ {
		c := s[x]

		// Uppercase all a-z
		if c >= 'a' && c <= 'z' {
			c -= ('a' - 'A')
		}

		// Record only a-z characters
		if c >= 'A' && c <= 'Z' {
			letters[c]++
			sum++
		}
	}

	if sum > 0 {
		for k, _ := range letters {
			letters[k] /= sum
		}
	}

	return letters
}

// Score a plaintext string by its similarity to English
// Lower score indicates a better match
func scoreString(s []byte) (score float32) {
	// Score by similarity to standard letter distribution
	frequencies := letterFrequency(s)
	for c, freq := range frequencies {
		d := freq - engFrequencies[c]
		if d < 0 {
			d *= -1
		}

		score += d
	}

	// Penalize control characters
	for x := 0; x < len(s); x++ {
		if s[x] < ' ' {
			score += .01
		}
	}

	return
}

func doDecrypt(c chan *decrypt, input []byte, key byte) {
	plaintext := singleXOR(input, key)
	weight := scoreString(plaintext)

	c <- &decrypt{key, string(plaintext), weight}
}

func bestDecryption(input []byte) *decrypt {
	// Attempt each possible key in a separate goroutine
	c := make(chan *decrypt, 256)
	for x := 0; x < 256; x++ {
		go doDecrypt(c, input, byte(x))
	}

	// Collect all results and sort
	results := make(decryptList, 256)
	for x := 0; x < 256; x++ {
		results[x] = <-c
	}
	sort.Sort(results)

	return results[0]
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Single hex string argument required")
		os.Exit(1)
	}

	crypt, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := bestDecryption(crypt)

	fmt.Println(result.Key, "\t", result.Plaintext)
}
