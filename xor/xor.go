package xor

import "errors"

func Apply(src, key []byte) ([]byte, error) {
	result := make([]byte, len(src))

	for i := 0; i < len(src); i++ {
		result[i] = src[i] ^ key[i%len(key)]
	}

	return result, nil
}

func ApplyByte(src []byte, key byte) ([]byte, error) {
	return Apply(src, []byte{key})
}

func ApplyFixed(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("Input buffers must match in length")
	}

	return Apply(a, b)
}
