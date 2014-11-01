package blocks

import "bytes"
import "crypto/rand"
import "testing"

func TestCBC(t *testing.T) {
	key := make([]byte, 16)
	_, _ = rand.Read(key)
	iv := make([]byte, 16)
	_, _ = rand.Read(iv)

	e, err := NewAES(key)
	if err != nil {
		t.Errorf("For NewAES got error %v", err)
	}

	tests := [][]byte{
		[]byte("Test block"),
		[]byte("Test plaintext for multi block CBC encryption test"),
	}

	for i := 0; i < len(tests); i++ {
		plaintext := tests[i]

		encrypted, err := e.EncryptCBC(plaintext, iv)
		if err != nil {
			t.Errorf("For e.EncryptCBC got error %v", err)
		}

		decrypted, err := e.DecryptCBC(encrypted, iv)
		if err != nil {
			t.Errorf("For e.DecryptCBC got error %v", err)
		}

		if bytes.Compare(plaintext, decrypted) != 0 {
			t.Errorf("Failed \"%s\"\nGot %v, expected %v", plaintext, decrypted, plaintext)
		}
	}

}
