package analysis

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := map[byte]uint8{
		0x00: 0, // 00000000
		0x10: 1, // 00010000
		0x31: 3, // 00110001
		0xB7: 6, // 10110111
		0xFF: 8, // 11111111
	}

	for b, w := range tests {
		r := ByteWeight(b)
		if r != w {
			t.Errorf("For %v got weight %v, expected %v", b, r, w)
		}
	}
}

func TestHammingDistace(t *testing.T) {
	tests := map[[2]string]uint{
		[2]string{"this is a test", "wokka wokka!!!"}: 37,
	}

	for s, w := range tests {
		a := []byte(s[0])
		b := []byte(s[1])
		r, err := Distance(a, b)

		if err != nil {
			t.Errorf("For %v,%v got error %v", s[0], s[1], err)
		}
		if r != w {
			t.Errorf("For %v,%v got weight %v, expected %v", s[0], s[1], r, w)
		}
	}
}
