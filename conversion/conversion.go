package conversion

import "encoding/base64"
import "encoding/hex"

func HexToBase64(src string) (string, error) {
	plain, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(plain), nil
}

func HexToBytes(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func BytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}

func Base64BytesToBytes(src []byte) []byte {
	result := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	base64.StdEncoding.Decode(result, src)
	return result
}

func TransposeBlocks(src []byte, blocksize int) [][]byte {
	blocks := make([][]byte, blocksize)
	for i := 0; i < blocksize; i++ {
		e := 0
		if len(src)%blocksize > i {
			e = 1
		}
		blocks[i] = make([]byte, (len(src)/blocksize)+e)
	}

	for i := 0; i < len(src); i++ {
		blocks[i%blocksize][i/blocksize] = src[i]
	}

	return blocks
}
