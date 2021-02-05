package mopp

import (
	"testing"
)

// TestCharToCode calls morse.CharToCode with a character, checking
// for a valid return value.
func TestMOPPEncodePARIS(t *testing.T) {
	words := []string{".--.", ".-", ".-.", "..", "..."}
	protocol := 1
	serial := 27
	speed := 16
	want := []byte{0x5b, 0x41, 0xa4, 0x61, 0x91, 0x45, 0x70}

	code, err := MOPPEncode(protocol, serial, speed, words)
	if len(want) != len(code) || err != nil {
		t.Fatalf(`MOPPEncode incorrect length returned %d, expected %d, or error %v`, len(code), len(want), err)
	}
	for i, _ := range want {
		if want[i] != code[i] || err != nil {
			t.Fatalf(`MOPPEncode = %q, %v, want match for %q, nil, different at index %d`, code, err, want, i)
		}
	}
}
