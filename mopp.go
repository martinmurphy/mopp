package mopp

import (
	"errors"
	"strings"
)

// Encode - encodes a morse word into an MOPP data packet
//  protocol: 01b
//  serial number: 6 bits
//  speed: 6 bits - between 5 and 60wpm inclusive
//  codechars: {".--.", ".-", ".-.", "..", "..."}
func Encode(protocol int, serial int, speed int, codechars string) ([]byte, error) {
	codechars = strings.TrimSpace(codechars)
	retval := make([]byte, 0, 30)
	currbyte := (byte(protocol) << 6) | (byte(serial) & 0x3f)
	retval = append(retval, currbyte)

	currbyte = (byte(speed) << 2)
	byteposition := 0
	val := byte(0)
	for _, character := range codechars {
		switch character {
		case '.':
			val = 1
		case '-':
			val = 2
		case ' ':
			val = 0
		}
		currbyte = currbyte | (val << (byteposition * 2))
		byteposition--
		if byteposition < 0 {
			retval = append(retval, currbyte)
			currbyte = 0
			byteposition = 3
		}
	}
	val = 3
	currbyte = currbyte | (val << (byteposition * 2))
	retval = append(retval, currbyte)

	return retval, nil
}

func Decode(packet []byte) (protocol int, serial int, speed int, codechars string, err error) {
	if len(packet) < 2 {
		err = errors.New("invalid packet, too short")
		return
	}

	protocol = int(packet[0] >> 6)
	speed = int(packet[1] >> 2)
	serial = int(packet[0] & 0x3f)

	bytePosition := 0
	val := ""
	for byteIndex := 1; byteIndex < len(packet); byteIndex++ {
		for ; bytePosition >= 0; bytePosition -= 2 {
			switch (packet[byteIndex] >> bytePosition) & 3 {
			case byte(0):
				val = " "
			case byte(1):
				val = "."
			case byte(2):
				val = "-"
			case byte(3):
				val = " "
				bytePosition = -1 // ensure we terminate
			}
			codechars += val
		}
		bytePosition = 6
	}

	codechars = strings.TrimSpace(codechars)
	return //0, 0, 0, make([]string, 0), errors.New("not implemented")
}
