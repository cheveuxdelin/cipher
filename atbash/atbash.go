package atbash

import (
	"errors"

	"github.com/cheveuxdelin/cipher/utils"
)

func Encode(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("empty string")
	}

	var rtn []rune = []rune(s)

	for i := 0; i < len(rtn); i++ {
		if !utils.IsASCIIPrintable(s[i]) {
			return "", errors.New("not ASCII")
		}
		rtn[i] = ' ' + '~' - rtn[i]
	}

	return string(rtn), nil
}

func Decode(s string) (string, error) {
	return Encode(s)
}
