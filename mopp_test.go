package mopp

import (
	"testing"
)

// TestCharToCode calls morse.CharToCode with a character, checking
// for a valid return value.
func TestMOPPEncodePARIS(t *testing.T) {
	characters := ".--. .- .-. .. ..."
	protocol := 1
	serial := 27
	speed := 16
	want := []byte{0x5b, 0x41, 0xa4, 0x61, 0x91, 0x45, 0x70}

	code, err := Encode(protocol, serial, speed, characters)
	if len(want) != len(code) || err != nil {
		t.Fatalf(`MOPPEncode incorrect length returned %d, expected %d, or error %v`, len(code), len(want), err)
	}
	for i := range want {
		if want[i] != code[i] || err != nil {
			t.Fatalf(`MOPPEncode = %q, %v, expected %q, nil, different at index %d`, code, err, want, i)
		}
	}
}

// TestCharToCode calls morse.CharToCode with a character, checking
// for a valid return value.
func TestMOPPEncodePARISwithTrailingSpace(t *testing.T) {
	characters := ".--. .- .-. .. ... "
	protocol := 1
	serial := 27
	speed := 16
	want := []byte{0x5b, 0x41, 0xa4, 0x61, 0x91, 0x45, 0x70}

	code, err := Encode(protocol, serial, speed, characters)
	if len(want) != len(code) || err != nil {
		t.Fatalf(`MOPPEncode incorrect length returned %d, expected %d, or error %v`, len(code), len(want), err)
	}
	for i := range want {
		if want[i] != code[i] || err != nil {
			t.Fatalf(`MOPPEncode = %q, %v, expected %q, nil, different at index %d`, code, err, want, i)
		}
	}
}

func TestMOPPDecodePARIS(t *testing.T) {
	wantedCharacters := ".--. .- .-. .. ..."
	wantedProtocol := 1
	wantedSerial := 27
	wantedSpeed := 16
	inputData := []byte{0x5b, 0x41, 0xa4, 0x61, 0x91, 0x45, 0x70}

	protocol, serial, speed, characters, err := Decode(inputData)
	if protocol != wantedProtocol {
		t.Fatalf(`Decode expected protocol %d, actual %d`, wantedProtocol, protocol)
	}
	if speed != wantedSpeed {
		t.Fatalf(`Decode expected speed %d, actual %d`, wantedSpeed, speed)
	}
	if serial != wantedSerial {
		t.Fatalf(`Decode expected serial %d, actual %d`, wantedSerial, serial)
	}
	if wantedCharacters != characters || err != nil {
		t.Fatalf(`MOPPDecode incorrect characters returned "%s", expected "%s"., or error %v`, characters, wantedCharacters, err)
	}
}
