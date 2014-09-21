package xor

import "github.com/leijou/matasano-go/plaintext"

type Decryption struct {
	Input  []byte
	Output []byte
	Key    []byte
	Score  float64
}

func NewDecription(ciphertext []byte, key []byte) (*Decryption, error) {
	text, err := Apply(ciphertext, key)
	if err != nil {
		return nil, nil
	}

	return &Decryption{
		ciphertext,
		text,
		key,
		plaintext.Score(text),
	}, nil
}

func BestByteDecryption(ciphertext []byte) (*Decryption, error) {
	var best *Decryption

	for key := 0; key < 256; key++ {
		decryption, err := NewDecription(ciphertext, []byte{byte(key)})
		if err != nil {
			return nil, nil
		}

		if best == nil || decryption.Score < best.Score {
			best = decryption
		}
	}

	return best, nil
}
