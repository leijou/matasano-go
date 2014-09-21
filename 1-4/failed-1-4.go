//
// Was attempting to determine single-character XOR without having to decrypt each one
// Unfortunately, I assume because of uppsercase/lowercase, this didn't work
//
// Something to revisit in the future
//

package main

import "bufio"
import "encoding/hex"
import "fmt"
import "os"
import "sort"

var engFrequencies = []float64{0.1270, 0.0906, 0.0817, 0.0751, 0.0697, 0.0675, 0.0633, 0.0609, 0.0599, 0.0425, 0.0403, 0.0278, 0.0276, 0.0241, 0.0236, 0.0223, 0.0202, 0.0197, 0.0193, 0.0129, 0.0098, 0.0077, 0.0015, 0.0015, 0.0010, 0.0007}

// Analysis result
type analysis struct {
	Line  string
	Score float64
}

// List of decryption results (sortable by weight)
type analysisList []*analysis

func (a analysisList) Len() int           { return len(a) }
func (a analysisList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a analysisList) Less(i, j int) bool { return a[i].Score < a[j].Score }

type sortBytes []byte

func (a sortBytes) Len() int           { return len(a) }
func (a sortBytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBytes) Less(i, j int) bool { return a[i] < a[j] }

type sortFloat64 []float64

func (a sortFloat64) Len() int           { return len(a) }
func (a sortFloat64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortFloat64) Less(i, j int) bool { return a[i] < a[j] }

// Take a list of bytes and return, in descending order, the frequency of their
// occurance
func letterFrequency(input []byte) []float64 {
	sort.Sort(sortBytes(input))

	var totals []int

	for x := 0; x < len(input); {
		bytetotal := float64(0)
		lastvalue := input[x]

		for ; x < len(input); x++ {
			if lastvalue != input[x] {
				break
			}

			bytetotal++
		}

		totals = append(totals, bytetotal)
	}

	distribution := make([]float64, len(totals))

	for x := 0; x < len(totals); x++ {
		distribution[x] = float64(totals[x]) / float64(len(input))
	}

	sort.Reverse(sortFloat64(distribution))

	return distribution
}

// Score an encrypted string by its similarity to English letter frequency
// Lower score indicates a better match
func scoreString(s []byte) (score float64) {
	frequencies := letterFrequency(s)

	// Score by similarity to standard letter distribution
	for x := 0; x < len(engFrequencies); x++ {
		if x >= len(frequencies) {
			score += engFrequencies[x]
		} else {
			d := frequencies[x] - engFrequencies[x]
			if d < 0 {
				d *= -1
			}
			score += d
		}
	}

	return
}

func doAnalysis(c chan *analysis, input string) {
	crypt, _ := hex.DecodeString(input)

	score := scoreString(crypt)

	c <- &analysis{input, score}
}

func bestAnalysis(inputs []string) *analysis {
	// Attempt each input
	c := make(chan *analysis, len(inputs))
	for x := 0; x < len(inputs); x++ {
		go doAnalysis(c, inputs[x])
	}

	// Collect all results and sort
	results := make(analysisList, len(inputs))
	for x := 0; x < len(inputs); x++ {
		results[x] = <-c
	}
	sort.Sort(results)

	for i := 0; i < 10; i++ {
		fmt.Println(results[i].Score, "\t", results[i].Line)
	}

	return results[0]
}

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

	// Build list of inputs
	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	result := bestAnalysis(inputs)

	fmt.Println(string(result.Line))
}
