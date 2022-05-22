package caesar

const N_LETTERS byte = 26
const N_ASCII_PRINTABLE = 224
const STARTING_ASCII_CODE_LOWERCASE byte = 97
const ENDING_ASCII_CODE_LOWERCASE byte = 122
const STARTING_ASCII_CODE_UPPERCASE byte = 65
const ENDING_ASCII_CODE_UPPERCASE byte = 90
const STARTING_ASCII_PRINTABLE byte = 32

//Encodes single letter giving n offset
func encodeLetter(b byte, n byte) byte {
	// converting to lowercase if needed
	if b >= STARTING_ASCII_CODE_UPPERCASE && b <= ENDING_ASCII_CODE_UPPERCASE {
		b += STARTING_ASCII_CODE_LOWERCASE - STARTING_ASCII_CODE_UPPERCASE
	}

	return byte((b+n-STARTING_ASCII_CODE_LOWERCASE)%N_LETTERS + STARTING_ASCII_CODE_LOWERCASE)
}

func encodeChar(b byte, n byte) byte {
	return byte((b+n-STARTING_ASCII_PRINTABLE)%N_ASCII_PRINTABLE + STARTING_ASCII_PRINTABLE)
}

// Encodes incoming string by n places. Can specify to only encode letters
func Encode(s string, n byte, onlyLetters bool) string {
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

	return string(encodedMessage)
}

// Decodes incoming string by n places. Can specify to only decode letters
func Decode(s string, n byte, onlyLetters bool) string {
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

	return string(encodedMessage)
}
