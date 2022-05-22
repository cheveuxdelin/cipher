package caesar

const N_LETTERS byte = 26
const N_ASCII_PRINTABLE = 224

//Encodes single letter giving n offset
func encodeLetter(b byte, n byte) byte {
	// converting to lowercase if needed
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}

	return byte((b+n-'a')%N_LETTERS + 'a')
}

func encodeChar(b byte, n byte) byte {
	return byte((b+n-' ')%N_ASCII_PRINTABLE + ' ')
}

// Encodes incoming string by n places. Can specify to only encode letters
func Encode(s string, n byte, onlyLetters bool) (string, error) {
	var encodedMessage []byte = []byte(s)

	if onlyLetters {
		for i := range encodedMessage {
			encodedMessage[i] = encodeLetter(s[i], n)
		}
	} else {
		for i := range encodedMessage {
			encodedMessage[i] = encodeChar(s[i], n)
		}
	}

	return string(encodedMessage), nil
}

// Decodes incoming string by n places. Can specify to only decode letters
func Decode(s string, n byte, onlyLetters bool) (string, error) {
	var encodedMessage []byte = []byte(s)

	if onlyLetters {
		for i := range encodedMessage {
			encodedMessage[i] = encodeLetter(s[i], -n)
		}
	} else {
		for i := range encodedMessage {
			encodedMessage[i] = encodeChar(s[i], -n)
		}
	}

	return string(encodedMessage), nil
}
