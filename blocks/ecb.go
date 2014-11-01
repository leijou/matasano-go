package blocks

import "github.com/leijou/matasano-go/conversion"

func (e *Block) EncryptECBBlock(input []byte) ([]byte, error) {
	return e.EncryptCBCBlock(input, make([]byte, e.BlockSize))
}

func (e *Block) DecryptECBBlock(input []byte) ([]byte, error) {
	return e.DecryptCBCBlock(input, make([]byte, e.BlockSize))
}

func (e *Block) EncryptECB(input []byte) ([]byte, error) {
	input = conversion.PadPKCS(input, e.BlockSize)
	length := len(input)
	ciphertext := make([]byte, length)

	// Loop over each block
	for a, b := 0, e.BlockSize; b <= length; a, b = a+e.BlockSize, b+e.BlockSize {
		cipherslice := ciphertext[a:b]
		cipherslice, err := e.EncryptECBBlock(input[a:b])
		if err != nil {
			return ciphertext, err
		}

		copy(ciphertext[a:b], cipherslice)
	}

	return ciphertext, nil
}

func (e *Block) DecryptECB(input []byte) ([]byte, error) {
	length := len(input)
	plaintext := make([]byte, length)

	// Loop over each block
	for a, b := 0, e.BlockSize; b <= length; a, b = a+e.BlockSize, b+e.BlockSize {
		plainslice := plaintext[a:b]
		plainslice, err := e.DecryptECBBlock(input[a:b])
		if err != nil {
			return plaintext, err
		}

		copy(plaintext[a:b], plainslice)
	}

	return plaintext, nil
}
