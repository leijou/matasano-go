package conversion

func PadPKCS(src []byte, blocksize int) []byte {
	srclength, dstlength := len(src), len(src)
	dstlength += blocksize - (srclength % blocksize)

	char := byte(dstlength - srclength)

	result := make([]byte, dstlength)
	for i := 0; i < dstlength; i++ {
		if i < srclength {
			result[i] = src[i]
		} else {
			result[i] = char
		}
	}

	return result
}

func UnPadPKCS(src []byte) []byte {
	srclength := len(src)
	trim := int(src[srclength-1])

	return src[:srclength-trim]
}

func DetectPKCS(src []byte) bool {
	srclength := len(src)
	padding := src[srclength-1]

	if padding == 0 || int(padding) > srclength {
		return false
	}

	for i := srclength - int(padding); i < srclength; i++ {
		if src[i] != padding {
			return false
		}
	}

	return true
}
