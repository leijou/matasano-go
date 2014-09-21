package main

import "encoding/base64"
import "errors"
import "fmt"
import "os"
import "io/ioutil"
import "sort"

var engFrequencies = map[byte]float64{'E': 0.1270, 'T': 0.0906, 'A': 0.0817, 'O': 0.0751, 'I': 0.0697, 'N': 0.0675, 'S': 0.0633, 'H': 0.0609, 'R': 0.0599, 'D': 0.0425, 'L': 0.0403, 'C': 0.0278, 'U': 0.0276, 'M': 0.0241, 'W': 0.0236, 'F': 0.0223, 'G': 0.0202, 'Y': 0.0197, 'P': 0.0193, 'B': 0.0129, 'V': 0.0098, 'K': 0.0077, 'J': 0.0015, 'X': 0.0015, 'Q': 0.0010, 'Z': 0.0007}

// Decryption result
type decrypt struct {
	Key       byte
	Plaintext string
	Score     float64
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

func letterFrequency(s []byte) map[byte]float64 {
	letters := map[byte]float64{'A': 0, 'B': 0, 'C': 0, 'D': 0, 'E': 0, 'F': 0, 'G': 0, 'H': 0, 'I': 0, 'J': 0, 'K': 0, 'L': 0, 'M': 0, 'N': 0, 'O': 0, 'P': 0, 'Q': 0, 'R': 0, 'S': 0, 'T': 0, 'U': 0, 'V': 0, 'W': 0, 'X': 0, 'Y': 0, 'Z': 0}

	sum := float64(0)

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
func scoreString(s []byte) (score float64) {
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

func hammingWeight(b byte) uint8 {
	hammingTable := []uint8{
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
	}
	return hammingTable[b]
}

func hammingDistance(a, b []byte) (uint, error) {
	if len(a) != len(b) {
		return 0, errors.New("Hamming distance buffers must match in length")
	}

	distance := uint(0)

	for i := 0; i < len(a); i++ {
		distance += uint(hammingWeight(a[i] ^ b[i]))
	}

	return distance, nil
}

type cryptbuffer struct {
	input []byte
}

func (c *cryptbuffer) keysizeDistance(keysize int, samples int) (float64, error) {
	if samples < 1 {
		return 0, errors.New("Not enough samples")
	}
	if keysize < 1 {
		return 0, errors.New("No keysize")
	}
	if keysize*(samples+1) > len(c.input) {
		return 0, errors.New("File too small")
	}

	totaldistance := float64(0)

	for i := int(0); i < samples; i++ {
		slice := c.input[(keysize * i) : (keysize*i)+(keysize*2)]

		distance, err := hammingDistance(slice[:keysize], slice[keysize:])
		if err != nil {
			return 0, err
		}

		totaldistance += float64(distance)
	}

	return (totaldistance / float64(samples)) / float64(keysize), nil
}

func (c *cryptbuffer) transposeBlocks(keysize int) [][]byte {
	blocks := make([][]byte, keysize)
	for i := 0; i < keysize; i++ {
		blocks[i] = make([]byte, (len(c.input)/keysize)+1)
	}

	for i := 0; i < len(c.input); i++ {
		blocks[i%keysize][i/keysize] = c.input[i]
	}

	return blocks
}

func (c *cryptbuffer) decrypt(key []byte) []byte {
	plaintext := make([]byte, len(c.input))

	for i := 0; i < len(c.input); i++ {
		plaintext[i] = c.input[i] ^ key[i%len(key)]
	}

	return plaintext
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("File to scan required")
		os.Exit(1)
	}

	// Read input file
	inputfile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode base64
	base64.StdEncoding.Decode(inputfile, inputfile)

	inputbuffer := &cryptbuffer{inputfile}

	samples := 5

	for keysize := 2; keysize <= 40; keysize++ {
		distance, _ := inputbuffer.keysizeDistance(keysize, samples)

		fmt.Println("Key Size:", keysize, "Distance:", distance)
	}

	keysize := 29
	key := make([]byte, keysize)
	blocks := inputbuffer.transposeBlocks(keysize)
	for i := 0; i < keysize; i++ {
		decrypt := bestDecryption(blocks[i])
		key[i] = decrypt.Key
	}

	fmt.Println(string(key))
	fmt.Println(string(inputbuffer.decrypt(key)))

}
