package atbash

import (
	"github.com/cheveuxdelin/cipher/errors"
	"github.com/cheveuxdelin/cipher/utils"
)

type AtbashJSON struct {
	Text string `json:"text" binding:"required"`
}

func Encode(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.ErrEmptyString
	}

	var rtn []rune = []rune(s)

	for i := 0; i < len(rtn); i++ {
		if !utils.IsASCIIPrintable(s[i]) {
			return "", errors.ErrStringNotASCII
		}
		rtn[i] = ' ' + '~' - rtn[i]
	}

	return string(rtn), nil
}
