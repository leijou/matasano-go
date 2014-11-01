package analysis

import cryptrand "crypto/rand"
import "math/rand"
import "github.com/leijou/matasano-go/blocks"
import "time"

func randomAESKey() []byte {
	res := make([]byte, 16)

	_, _ = cryptrand.Read(res)

	return res
}

func RandomOracle(input []byte) ([]byte, string, error) {
	rand.Seed(time.Now().UnixNano())

	pre := make([]byte, rand.Intn(6)+5) // 5 - 10
	_, _ = cryptrand.Read(pre)

	post := make([]byte, rand.Intn(6)+5) // 5 - 10
	_, _ = cryptrand.Read(post)

	plaintext := make([]byte, len(pre)+len(post)+len(input))
	copy(plaintext, pre)
	copy(plaintext[len(pre):], input)
	copy(plaintext[len(pre)+len(input):], post)

	e, _ := blocks.NewAES(randomAESKey())

	if rand.Intn(2) > 0 {
		encrypted, err := e.EncryptCBC(plaintext, randomAESKey())
		return encrypted, "CBC", err
	} else {
		encrypted, err := e.EncryptECB(plaintext)
		return encrypted, "ECB", err
	}
}
