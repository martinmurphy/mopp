package mopp

func MOPPEncode(protocol int, serial int, speed int, codechars []string) ([]byte, error) {
	retval := make([]byte, 0, 30)
	currbyte := (byte(protocol) << 6) | (byte(serial) & 0x3f)
	retval = append(retval, currbyte)

	currbyte = (byte(speed) << 6)

	for i := 0; i < 6; i++ {
		retval = append(retval, 0)
	}

	return retval, nil
}
