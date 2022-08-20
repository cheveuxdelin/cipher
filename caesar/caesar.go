package caesar

import (
	"github.com/cheveuxdelin/cipher/errors"
	"github.com/cheveuxdelin/cipher/utils"
)

const N_LETTERS byte = 26
const N_ASCII_PRINTABLE = 224

type CaesarJSON struct {
	Text        string `json:"text" binding:"required"`
	N           byte   `json:"n" binding:"required"`
	OnlyLetters bool   `json:"only_letters"`
}

// Encodes single letter giving n offset
func encodeLetter(b byte, n byte) byte {
	// converting to lowercase if needed
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		b = utils.ToUpper(b)
		return byte((b+n-'a')%N_LETTERS + 'a')
	} else {
		return b
	}
}

func encodeChar(b byte, n byte) byte {
	return byte((b+n-' ')%N_ASCII_PRINTABLE + ' ')
}

// Encodes incoming string by n places. Can specify to only encode letters
func Encode(s string, n byte, onlyLetters bool) (string, error) {
	var encodedMessage []byte = []byte(s)

	if onlyLetters {
		for i := range encodedMessage {
			if !utils.IsASCIIPrintable(s[i]) {
				return "", errors.ErrStringNotASCII
			}

			encodedMessage[i] = encodeLetter(s[i], n)
		}
	} else {
		for i := range encodedMessage {
			if !utils.IsASCIIPrintable(s[i]) {
				return "", errors.ErrStringNotASCII
			}

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
			if !utils.IsASCIIPrintable(s[i]) {
				return "", errors.ErrStringNotASCII
			}

			encodedMessage[i] = encodeLetter(s[i], -n)
		}
	} else {
		for i := range encodedMessage {
			if !utils.IsASCIIPrintable(s[i]) {
				return "", errors.ErrStringNotASCII
			}

			encodedMessage[i] = encodeChar(s[i], -n)
		}
	}

	return string(encodedMessage), nil
}
