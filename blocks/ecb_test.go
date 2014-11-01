package blocks

import "bytes"
import "crypto/rand"
import "testing"

func TestECB(t *testing.T) {
	key := make([]byte, 16)
	_, _ = rand.Read(key)

	e, err := NewAES(key)
	if err != nil {
		t.Errorf("For NewAES got error %v", err)
	}

	tests := [][]byte{
		[]byte("Test block"),
		[]byte("Test plaintext for multi block ECB encryption test"),
	}

	for i := 0; i < len(tests); i++ {
		plaintext := tests[i]

		encrypted, err := e.EncryptECB(plaintext)
		if err != nil {
			t.Errorf("For e.EncryptECB got error %v", err)
		}

		decrypted, err := e.DecryptECB(encrypted)
		if err != nil {
			t.Errorf("For e.DecryptECB got error %v", err)
		}

		if bytes.Compare(plaintext, decrypted) != 0 {
			t.Errorf("Failed \"%s\"\nGot %v, expected %v", plaintext, decrypted, plaintext)
		}
	}

}
