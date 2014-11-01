package blocks

import "errors"
import "github.com/leijou/matasano-go/conversion"
import "github.com/leijou/matasano-go/xor"

var BlockSizeError error = errors.New("input length does not match block length")
var IVSizeError error = errors.New("iv length does not match block length")

func (e *Block) EncryptCBCBlock(input, iv []byte) ([]byte, error) {
	if len(input) != e.BlockSize {
		return nil, BlockSizeError
	}

	encryptedtext := make([]byte, e.BlockSize)

	input, err := xor.ApplyFixed(input, iv)
	if err != nil {
		return nil, err
	}

	e.Cypher.Encrypt(encryptedtext[0:e.BlockSize], input[0:e.BlockSize])

	return encryptedtext, nil
}

func (e *Block) DecryptCBCBlock(input, iv []byte) ([]byte, error) {
	if len(input) != e.BlockSize {
		return nil, BlockSizeError
	}

	plaintext := make([]byte, e.BlockSize)

	e.Cypher.Decrypt(plaintext[0:e.BlockSize], input[0:e.BlockSize])

	return xor.ApplyFixed(plaintext, iv)
}

func (e *Block) EncryptCBC(input, iv []byte) ([]byte, error) {
	input = conversion.PadPKCS(input, e.BlockSize)

	length := len(input)
	ciphertext := make([]byte, length)

	if len(iv) != e.BlockSize {
		return ciphertext, IVSizeError
	}

	// Loop over each block
	for a, b := 0, e.BlockSize; b <= length; a, b = a+e.BlockSize, b+e.BlockSize {
		plainslice := input[a:b]

		cipherslice := make([]byte, e.BlockSize)
		cipherslice, err := e.EncryptCBCBlock(plainslice, iv)
		if err != nil {
			return ciphertext, err
		}

		copy(ciphertext[a:b], cipherslice)

		iv = cipherslice
	}

	return ciphertext, nil
}

func (e *Block) DecryptCBC(input, iv []byte) ([]byte, error) {
	length := len(input)
	plaintext := make([]byte, length)

	if len(iv) != e.BlockSize {
		return plaintext, IVSizeError
	}

	// Loop over each block
	for a, b := 0, e.BlockSize; b <= length; a, b = a+e.BlockSize, b+e.BlockSize {
		plainslice := make([]byte, e.BlockSize)
		plainslice, err := e.DecryptCBCBlock(input[a:b], iv)
		if err != nil {
			return plaintext, err
		}

		copy(plaintext[a:b], plainslice)

		iv = input[a:b]
	}

	return plaintext, nil
}
