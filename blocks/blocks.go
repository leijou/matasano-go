package blocks

import "crypto/aes"
import "crypto/cipher"

type Block struct {
	Key       []byte
	Cypher    cipher.Block
	BlockSize int
}

func NewAES(key []byte) (*Block, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &Block{key, c, aes.BlockSize}, nil
}
