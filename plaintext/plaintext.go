package plaintext

var engFrequencies = map[byte]float64{'E': 0.1270, 'T': 0.0906, 'A': 0.0817, 'O': 0.0751, 'I': 0.0697, 'N': 0.0675, 'S': 0.0633, 'H': 0.0609, 'R': 0.0599, 'D': 0.0425, 'L': 0.0403, 'C': 0.0278, 'U': 0.0276, 'M': 0.0241, 'W': 0.0236, 'F': 0.0223, 'G': 0.0202, 'Y': 0.0197, 'P': 0.0193, 'B': 0.0129, 'V': 0.0098, 'K': 0.0077, 'J': 0.0015, 'X': 0.0015, 'Q': 0.0010, 'Z': 0.0007}

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
func Score(s []byte) float64 {
	score := float64(0)

	// Score by similarity to standard letter distribution
	frequencies := letterFrequency(s)
	for c, freq := range frequencies {
		d := freq - engFrequencies[c]
		if d < 0 {
			d *= -1
		}

		score += d
	}

	for x := 0; x < len(s); x++ {
		// Penalize control characters
		if s[x] < ' ' {
			score += .01
		}

		// Penalize out-of-ascii characters
		if s[x] >= 127 {
			score += 1
		}
	}

	return score
}
