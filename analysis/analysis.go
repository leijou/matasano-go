package analysis

import "errors"

type Keysize struct {
	Keysize int
	Samples int
	Score   float64
}

func NewKeysize(ciphertext []byte, keysize, samples int) (*Keysize, error) {
	if samples < 1 {
		return nil, errors.New("Not enough samples")
	}
	if keysize < 1 {
		return nil, errors.New("No keysize")
	}
	if keysize*(samples+1) > len(ciphertext) {
		return nil, errors.New("File too small")
	}

	totaldistance := float64(0)

	for i := int(0); i < samples; i++ {
		slice := ciphertext[(keysize * i) : (keysize*i)+(keysize*2)]

		distance, err := Distance(slice[:keysize], slice[keysize:])
		if err != nil {
			return nil, err
		}

		totaldistance += float64(distance)
	}

	return &Keysize{
		keysize,
		samples,
		(totaldistance / float64(samples)) / float64(keysize),
	}, nil
}

func BestKeysize(ciphertext []byte, samples int) (*Keysize, error) {
	var best *Keysize

	for keysize := 2; keysize <= 40; keysize++ {
		ka, err := NewKeysize(ciphertext, keysize, samples)
		if err != nil {
			return nil, nil
		}

		if best == nil || ka.Score < best.Score {
			best = ka
		}
	}

	return best, nil
}
